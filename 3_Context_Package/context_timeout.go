package main

import (
	"context"
	"fmt"
	"time"
)

func fetchData(ctx context.Context, id int) {
	select {
	case <-time.After(3 * time.Second): // Simulate slow operation
		fmt.Println("Data fetched successfully", id)
	case <-ctx.Done(): // Context timeout or cancellation
		fmt.Println("Operation timed out:", ctx.Err(), id)
		return
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	for i := 1; i <= 3; i++ {
		go fetchData(ctx, i)
	}

	time.Sleep(5 * time.Second) // Wait longer to see timeout
	fmt.Println("Main: Done")
}
