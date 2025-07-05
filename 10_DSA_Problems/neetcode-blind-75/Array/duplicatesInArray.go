package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 1, 2, 6, 7, 8}
	fmt.Println(findDuplicatesInArray(arr))
}

func findDuplicatesInArray(arr []int) []int {
	seen := make(map[int]bool)
	sli := []int{}

	for _, num := range arr {
		if seen[num] {
			sli = append(sli, num)
		} else {
			seen[num] = true
		}
	}
	return sli
}
