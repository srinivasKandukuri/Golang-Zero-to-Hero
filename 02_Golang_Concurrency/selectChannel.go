package main

import (
	"fmt"
	"time"
)

// 🔄 Select and Multiplexing
// Design a select statement to read from multiple channels with timeout.
func main() {

	var ch1 = make(chan int)
	var ch2 = make(chan int)

	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 1
	}()

	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- 2
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received from input1", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received from input2", msg2)
	case <-time.After(1500 * time.Millisecond):
		fmt.Println("timeout")
	}

	close(ch1)
	close(ch2)

}
