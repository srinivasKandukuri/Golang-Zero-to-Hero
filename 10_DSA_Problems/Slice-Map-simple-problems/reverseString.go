package main

import "fmt"

func main() {

	str := "srinivas"

	runes := []rune(str)

	n := len(runes)

	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]

		// Which is equivalent to above line
		// Is a multiple assignment in Go, and it is used here to swap two elements in the runes slice.
		/*	temp := runes[i]
			runes[i] = runes[n-1-i]
			runes[n-1-i] = temp*/
	}

	fmt.Println(string(runes))
}
