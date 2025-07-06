package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func fetchData(ctx context.Context) {

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	counter := 0
	for {
		select {
		case <-ticker.C: // Simulate slow operation
			fmt.Println("ticker", counter)
			counter++
			time.Sleep(2 * time.Second)
			if counter == 5 {
				fmt.Println("worker : finished work")
				return
			}
		case <-ctx.Done(): // Context timeout or cancellation
			fmt.Println("Operation timed out:", ctx.Err())
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fetchData(ctx)
	}()

	wg.Wait()
	fmt.Println("Main: Done")
}
