package main

import "fmt"

func main() {
	arr := []int{10, 20, 30, 40, 50, 60}

	firstArr := arr[:len(arr)/2]
	secondArr := arr[len(arr)/2:]

	firstSumChan := make(chan int)
	go sumFirstPart(firstArr, firstSumChan)
	secondSumChan := make(chan int)
	go sumSecondPart(secondArr, secondSumChan)

	fmt.Printf("First Sum: %d\n", <-firstSumChan)
	fmt.Printf("Second Sum: %d\n", <-secondSumChan)
}

func sumFirstPart(firstArr []int, firstSumChan chan int) {
	sum := 0
	for _, value := range firstArr {
		sum += value
	}
	firstSumChan <- sum
	close(firstSumChan)
}

func sumSecondPart(secondArr []int, secondSumChan chan int) {
	sum := 0
	for _, value := range secondArr {
		sum += value
	}
	secondSumChan <- sum
	close(secondSumChan)
}
