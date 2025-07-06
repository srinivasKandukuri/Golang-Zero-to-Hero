package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Printf("Produced %d\n", i)
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		n := <-ch
		fmt.Printf("Consumed %d\n", n)
	}
}

func main() {
	ch := make(chan int)
	wg := new(sync.WaitGroup)
	wg.Add(2)
	go producer(ch, wg)
	go consumer(ch, wg)
	wg.Wait()

}
