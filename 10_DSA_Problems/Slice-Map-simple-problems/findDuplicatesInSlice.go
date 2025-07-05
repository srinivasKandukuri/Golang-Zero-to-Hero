package main

import "fmt"

/*
2. ✅ Find Duplicates in a Slice
Why it’s asked:

Tests your ability to use maps for frequency counting.
Example:

Input: []int{1, 2, 3, 4, 2}
Output: true (duplicates exist)
Key Concepts:

Map for counting
Iterating through slices
*/

func main() {
	arr := []int{1, 2, 3, 4, 2}

	fmt.Printf("this slice container duplicate values %t", hasDuplicated2(arr))
}
func hasDuplicates(arr []int) bool {
	var countMap = make(map[int]struct{}, len(arr))

	for _, v := range arr {
		if _, ok := countMap[v]; ok {
			return true
		}
		countMap[v] = struct{}{}
	}
	return false
}

func hasDuplicated2(arr []int) bool {
	seen := make(map[int]bool)
	for _, v := range arr {
		if seen[v] {
			return true
		}
		seen[v] = true
	}
	return false
}
