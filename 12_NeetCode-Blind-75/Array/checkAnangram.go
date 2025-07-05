package main

/*
3. ✅ Check for Anagrams
Why it’s asked:

Checks your skill with maps and sorting.

Example:
Input: "listen", "silent"
Output: true (they are anagrams)


Key Concepts:

Counting characters using a map
Sorting slices of runes

*/

import (
	"fmt"
	"sort"
)

func main() {

	str1 := "car"
	str2 := "tac"

	fmt.Printf("given strings are anagram %t", isAnagram2(str1, str2))
}

func isAnagarm(str1 string, str2 string) bool {

	if len(str1) != len(str2) {
		return false
	}

	s1 := []rune(str1)
	s2 := []rune(str2)

	sort.Slice(s1, func(i, j int) bool {
		return s1[i] < s1[j]
	})

	sort.Slice(s2, func(i, j int) bool {
		return s2[i] < s2[j]
	})
	return string(s1) == string(s2)
}

func isAnagram2(str1, str2 string) bool {

	if len(str1) != len(str2) {
		return false
	}
	countmap := make(map[rune]int)

	for _, char := range str1 {
		countmap[char]++
	}

	for _, char := range str2 {
		countmap[char]--
		if countmap[char] < 0 {
			return false
		}
	}
	fmt.Printf("map %v\n", countmap)

	return true
}
