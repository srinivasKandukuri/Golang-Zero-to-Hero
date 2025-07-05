package main

// this example running with sequential call

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var ErrInterrupt = errors.New("received interrupt")
var ErrTimeout = context.DeadlineExceeded

type Runner struct {
	tasks []func(ctx context.Context, id int)
}

func NewRunner() *Runner {
	return &Runner{}
}

func (r *Runner) Add(tasks ...func(ctx context.Context, id int)) {
	r.tasks = append(r.tasks, tasks...)
}

func (r *Runner) Start(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer stop()

	for id, task := range r.tasks {
		select {
		case <-ctx.Done():
			if errors.Is(ctx.Err(), context.Canceled) {
				return ErrInterrupt
			}
			fmt.Println("exited")
			return ErrTimeout
		default:
			task(ctx, id)
		}
	}
	return nil
}

func AddTask() func(ctx context.Context, id int) {
	return func(ctx context.Context, id int) {
		fmt.Printf("Processing - task %d\n", id)
		time.Sleep(1 * time.Second)
	}
}

func AddTask2() func(ctx context.Context, id int) {
	return func(ctx context.Context, id int) {
		fmt.Printf("Processing - task %d\n", id)

		select {
		case <-time.After(3 * time.Second):
			fmt.Printf("Finished - task %d\n", id)
		case <-ctx.Done():
			fmt.Printf("Cancelled - task %d\n", id)
			fmt.Println(ctx.Err())
			return
		}

	}
}

func main() {
	runner := NewRunner()
	runner.Add(AddTask(), AddTask2())

	err := runner.Start(1 * time.Millisecond)
	if err != nil {
		fmt.Println(err)
	}
}
