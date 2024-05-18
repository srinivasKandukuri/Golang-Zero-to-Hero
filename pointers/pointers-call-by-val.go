package main

import "fmt"

func callByVal(year int) {
	year = year + 1
}

func main() {
	year := 2024
	//yearAdd := &year

	callByVal(year)
	fmt.Println(year)
}
