package main

import (
	"fmt"
	"math"
)

func main() {
	for i := 0; i <= 100; i++ {
		//fmt.Println(i, isPrimeNum(i))
		if isPrimeNum(i) {
			fmt.Println(i)
		}
	}
}

func isPrimeNum(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
