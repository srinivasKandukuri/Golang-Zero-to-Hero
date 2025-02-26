package main

import "fmt"

func main() {
	channel := make(chan string)

	go func() {
		channel <- "Hello"
	}()
	msg := <-channel
	fmt.Println(msg)
}

//Example of Channel Communication:
