package main

// this example running with sequential call

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var ErrInterrupt1 = errors.New("received interrupt")
var ErrTimeout1 = context.DeadlineExceeded

type RunnerParrel struct {
	tasks []func(ctx context.Context, id int)
}

func NewRunnerParrel() *RunnerParrel {
	return &RunnerParrel{}
}

func (r *RunnerParrel) Add(tasks ...func(ctx context.Context, id int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *RunnerParrel) Start(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	g, gtx := errgroup.WithContext(ctx)

	for id, task := range r.tasks {
		task := task
		id := id

		g.Go(func() error {
			task(gtx, id)
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		if errors.Is(ctx.Err(), context.Canceled) {
			return ErrInterrupt1
		}
		return err
	}

	// Context may still have timed out
	if ctx.Err() != nil {
		return ctx.Err()
	}

	return nil
}

func AddTask3(sec time.Duration) func(ctx context.Context, id int) {
	return func(ctx context.Context, id int) {
		fmt.Printf("Processing - task %d\n", id)

		select {
		case <-time.After(sec):
			fmt.Printf("Finished - task %d\n", id)
		case <-ctx.Done():
			fmt.Printf("Cancelled - task %d\n", id)
			return
		}

	}
}

func main() {
	runner := NewRunnerParrel()
	runner.Add(AddTask3(2*time.Second), AddTask3(4*time.Second), AddTask3(5*time.Second), AddTask3(6*time.Second))

	err := runner.Start(10 * time.Second)
	if err != nil {
		fmt.Println(err)
	}
}
