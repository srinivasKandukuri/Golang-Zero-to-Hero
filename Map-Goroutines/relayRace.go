package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	baton := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go run(&wg, baton)
	baton <- 1
	wg.Wait()
}

func run(wg *sync.WaitGroup, baton chan int) {
	var newRunner int

	runner := <-baton

	fmt.Printf("Runner %d\n", runner)

	if runner != 4 {
		newRunner = runner + 1
		fmt.Printf("Runner is on line %d\n", newRunner)
		go run(wg, baton)
	}
	time.Sleep(100 * time.Millisecond)
	if runner == 4 {
		fmt.Printf("Runner %d finished the rance", runner)
		wg.Done()
		return
	}

	fmt.Printf("Runner %d Exchange with runner %d\n", runner, newRunner)
	baton <- newRunner
}
