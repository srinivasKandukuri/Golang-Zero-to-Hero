package main

import "fmt"

func main() {
	fmt.Println(string(firstNonRepeatingChar("swiss"))) // Output: "w"
}

func firstNonRepeatingChar(s string) rune {
	count := make(map[rune]int)

	for _, ch := range s {
		count[ch]++
	}

	for _, ch := range s {
		if count[ch] == 1 {
			return ch
		}
	}
	return 0
}
