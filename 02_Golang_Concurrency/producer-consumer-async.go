package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Inside Asynchronized Producer Consumer")
	number := 5
	datachannel := make(chan int, number)

	waitChannel := make(chan bool)

	go producerAsync(datachannel, number)
	go consumerAsync(datachannel, waitChannel)
	<-waitChannel

}

func producerAsync(datachannel chan int, number int) {
	for i := 0; i < number; i++ {
		datachannel <- i
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Done producerAsync")
	time.Sleep(2 * time.Second)
	datachannel <- 10000
	close(datachannel)
}

func consumerAsync(datachannel chan int, waitChannel chan bool) {

	for data := range datachannel {
		fmt.Printf("Consumer %d\n", data)
	}
	fmt.Println("Done Consuming")
	waitChannel <- true
}
