package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch := make(chan string, 3)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range ch {
			fmt.Println("Received : ", msg)
			time.Sleep(500 * time.Millisecond)
		}

	}()

	ch <- "Hello World start"
	ch <- "Hello World process"
	ch <- "Hello World end"
	close(ch)
	wg.Wait()

}

func main1() {
	logCh := make(chan string, 3) // buffered
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range logCh {
			fmt.Println("Log:", msg)
			time.Sleep(500 * time.Millisecond) // simulate slow logging
		}
	}()

	logCh <- "Start"
	logCh <- "Processing"
	logCh <- "Finished"
	close(logCh)
	wg.Wait()
}
