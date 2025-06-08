package main

import (
	"fmt"
	"time"
)

func main() {

	// declare channel
	var ch = make(chan int, 1)

	// sending value to channel
	ch <- 1

	// receive value from channel
	val := <-ch
	fmt.Println(val)

	//closing a channel

	close(ch)

	// looping channel

	for val := range ch {
		fmt.Println(val)
	}

	//one way channel direction
	logCh := make(chan string)
	go logger(logCh)
	sendData(logCh)
	time.Sleep(time.Second)
	close(logCh)
}

// one way direction of receiving data from channel
func logger(ch <-chan string) {
	for ch1 := range ch {
		fmt.Println("Log", ch1)
	}
}

// one way direction of channel sending data
func sendData(ch chan<- string) {
	ch <- "starting system"
	ch <- "processing data"
	ch <- "ending system"
}
