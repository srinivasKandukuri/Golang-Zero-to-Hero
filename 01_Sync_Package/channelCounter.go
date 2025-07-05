package main

import (
	"fmt"
	"sync"
)

type Counterr struct {
	value int
}

func (c *Counterr) inc() {
	c.value++
}

func (c *Counterr) getValue() int {
	return c.value
}

func main() {

	c := &Counterr{}
	incChan := make(chan struct{})
	var wg sync.WaitGroup
	//var counterMg sync.WaitGroup
	resultsChan := make(chan int)
	//counterMg.Add(1)
	go func() {
		//defer counterMg.Done()
		for range incChan {
			c.inc()
		}
		resultsChan <- c.getValue()
	}()

	worker := func() {
		defer wg.Done()

		for j := 0; j < 1000; j++ {
			incChan <- struct{}{} // sending increment
		}
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go worker()
	}

	wg.Wait()
	close(incChan)
	//counterMg.Wait()
	finalvalue := <-resultsChan
	fmt.Println("final counter", finalvalue)
}
