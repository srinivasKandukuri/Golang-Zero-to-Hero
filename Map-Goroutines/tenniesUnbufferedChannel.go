package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var wg sync.WaitGroup

func main() {

	court := make(chan int)
	wg.Add(2)
	go Player("Sk", court)
	go Player("Pk", court)
	court <- 1
	wg.Wait()
}

func Player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court

		if !ok {
			fmt.Println("channel is closed")
			return
		}
		fmt.Println("Player ", name, "ball", ball)

		n := rand.Intn(100)
		if n%13 == 0 {
			fmt.Printf("%s missed! Closing court.\n", name)
			close(court)
			return
		}

		ball++
		court <- ball
	}
}
