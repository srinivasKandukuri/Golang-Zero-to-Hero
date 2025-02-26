package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go worker(ch1)

	select {
	case num := <-ch1:
		fmt.Println(num)
	}

	time.Sleep(3 * time.Second)
}

func worker(ch chan int) {
	ch <- 1
	time.Sleep(1 * time.Second)
}
