package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	for i := 0; i < 10; i++ {

		// Non-blocking send
		select {
		case ch <- i:
			fmt.Println("Sent ", i)
		default:
			fmt.Println("Channel full")
		}

		// Non-blocking receive
		select {
		case val := <-ch:
			fmt.Println("Received:", val)
		default:
			fmt.Println("Channel empty")
		}
		time.Sleep(300 * time.Millisecond)
	}

}
