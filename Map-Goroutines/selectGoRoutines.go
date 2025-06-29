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
			fmt.Println("1 sec")
			if i == 10 {
				return
			}
		}
	}

}

func worker(ch chan int) {
	time.Sleep(3 * time.Second)
	ch <- 1
	time.Sleep(1 * time.Second)
}
