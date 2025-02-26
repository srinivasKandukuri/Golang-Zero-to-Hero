package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go workerG(i, &wg)
	}

	wg.Wait()
	fmt.Println("all workers completed")
}

func workerG(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("worker", id, "started")
}
