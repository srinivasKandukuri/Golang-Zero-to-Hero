package main

import (
	"fmt"
	"time"
)

func PrintMessage(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {
	go PrintMessage("hello this is go routines")
	time.Sleep(time.Millisecond * 10000)
	PrintMessage("without Go routine")
}
