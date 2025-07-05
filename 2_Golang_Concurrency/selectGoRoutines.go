package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)

	go worker(ch1)

	i := 1
	for {
		select {
		case num := <-ch1:
			fmt.Println(num)
		case <-time.After(2 * time.Second):
			i++
			fmt.Println("current i value", i)
			if i == 10 {
				return
			}
		}
	}

}

func worker(ch chan<- int) {
	ch <- 1
	time.Sleep(1 * time.Second)
}
