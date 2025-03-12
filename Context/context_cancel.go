package main

import (
	"context"
	"fmt"
	"time"
)

type response struct {
	data int
	err  error
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // Ensure context is cancelled properly
	ch := make(chan response)
	go worker(ctx, 1, ch)
	time.Sleep(3 * time.Second)
	fmt.Println("Canceling the context")
	cancel()
	// Wait for the worker to finish and capture responses
	for res := range ch {
		fmt.Printf("Received from worker: %+v\n", res)
	}
	// Allow the worker to run for 3 seconds
	time.Sleep(2 * time.Second)
}

func worker(ctx context.Context, id int, ch chan response) {
	defer close(ch) // Ensure channel is closed when worker exits
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d received cancellation signal\n", id)
			ch <- response{data: id, err: fmt.Errorf("worker %d shutting down", id)}
			return
		case <-time.After(1 * time.Second):
			fmt.Printf("Worker %d is working\n", id)
			ch <- response{data: id, err: nil}
		}
	}
}
