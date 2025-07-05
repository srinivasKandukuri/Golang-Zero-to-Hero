package main

import "time"

func leakyGoroutine() {
	leakyChan := make(chan int)
	go func() {
		leakyChan <- 1
	}()
}
func main() {
	leakyGoroutine()
	time.Sleep(10 * time.Second)
}
