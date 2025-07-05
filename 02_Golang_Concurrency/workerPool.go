package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	Jobs()
}

func Jobs() {
	numberOfJobs := 20
	numberOfWorkers := 2

	jobs := make(chan int, numberOfJobs)

	var wg sync.WaitGroup

	for i := 0; i <= numberOfWorkers; i++ {
		wg.Add(1)
		fmt.Println("Starting worker", i)
		go worker1(i, &wg, jobs)
	}

	for j := 1; j <= numberOfJobs; j++ {
		fmt.Println("Job", j)
		jobs <- j
	}

	close(jobs)
	wg.Wait()

}

func worker1(i int, wg *sync.WaitGroup, jobs <-chan int) {
	fmt.Println("worker", i)
	defer wg.Done()

	for j := range jobs {
		fmt.Println("worker", i, "job ", j)
		time.Sleep(2 * time.Second)
	}
}
