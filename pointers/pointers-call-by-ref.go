package main

import "fmt"

func callByRef(year *int) {
	*year = *year + 1
}

func main() {
	year := 2024
	yearAdd := &year

	callByRef(yearAdd)
	fmt.Println(year)
}
