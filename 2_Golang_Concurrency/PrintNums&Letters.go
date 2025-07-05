package main

import (
	"fmt"
	"sync"
)

/*
Problem Description

Use two goroutinealternate printing sequences, one goroutineto print numbers and the other goroutineto print letters. The final effect is as follows:

12AB34CD56EF78GH910IJ1112KL1314MN1516OP1718QR1920ST2122UV2324WX2526YZ2728



Solution

The problem is very simple. Use channels to control the progress of printing. Use two channels to control the printing sequence of numbers and letters respectively. After the numbers are printed, the letters are notified to print through the channel. After the letters are printed, the numbers are notified to print, and the cycle repeats.


Source code analysis

Two notification routines are used here channel. letter is used to notify the goroutine that prints letters to print letters, and number is used to notify the goroutine that prints numbers to print numbers. wait is used to exit the loop after the letters are printed.

You can also use the three channels to control the input of numbers, letters and termination signals respectively.
*/

// without wait group
func main2() {
	var numbers = make(chan bool)
	var letters = make(chan bool)
	var done = make(chan bool)
	go func() {
		i := 1
		for {
			select {
			case <-numbers:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letters <- true
			}
		}
	}()
	go func() {
		j := 'A'
		for {
			select {
			case <-letters:
				if j >= 'Z' {
					done <- true
					return
				} else {
					fmt.Print(string(j))
					j++
					fmt.Print(string(j))
					j++
					numbers <- true

				}
			}
		}
	}()
	numbers <- true // Start the sequence

	for {
		select {
		case <-done:
			return
		}
	}
}

// with wait group
func main() {
	var numbers = make(chan bool)
	var letters = make(chan bool)
	var wg sync.WaitGroup
	go func() {
		i := 1
		for {
			select {
			case <-numbers:
				fmt.Print(i)
				i++
				fmt.Print(i)
				i++
				letters <- true
			}
		}
	}()
	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		j := 'A'
		for {
			select {
			case <-letters:
				if j >= 'Z' {
					wg.Done()
					return
				} else {
					fmt.Print(string(j))
					j++
					fmt.Print(string(j))
					j++
					numbers <- true
				}
			}
		}
	}(&wg)
	numbers <- true // Start the sequence

	wg.Wait()
}
