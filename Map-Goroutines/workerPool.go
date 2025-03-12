package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	numberOfJobs := 5
	workers := 2
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	var wg sync.WaitGroup
	for i := 1; i <= workers; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := range jobs {
				fmt.Println("worker", i, "job ", j)
				time.Sleep(2 * time.Second)
				results <- j * 2
			}
		}(i)
	}

	//send 5 jobs to channel
	for i := 1; i <= numberOfJobs; i++ {
		jobs <- i
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println("result", res)
	}
}
