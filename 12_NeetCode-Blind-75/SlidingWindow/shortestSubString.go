/*
shortest substring without repeating characters

there is fixed size window,so it's dynamic shrinking window
Use while-loop inside to shrink when needed.

aabcda

ab
bc
da

aa -> condition fail, char repeated => remove left add right

	 ab ->	condition not fail, minLen = 2 => updated
		abc -> minLen < currentWindowLen -> then  minLen = currentWindowLen
	     bc -> minLen <= currentSubstring => condition match then break the loop
	     bcd -> minLen <= currentSubstring => condition fail -> remove the left -> don't add right => while loop => left++
		  cd -> minLen <= currentSubstring => condition match then break the loop
	      cda -> minLen <= currentSubstring => condition fail -> remove the left -> don't add right => while loop => left++
			da -> minLen <= currentSubstring => condition match then break the loop

minLen = 2
*/
package main

import (
	"fmt"
	"math"
)

func shortestUniqueSubstring(s string) int {
	seen := make(map[byte]bool)
	left := 0
	minLen := math.MaxInt32

	for right := 0; right < len(s); right++ {
		for seen[s[right]] { // Shrink window if duplicate found
			delete(seen, s[left])
			left++
		}
		seen[s[right]] = true              // Add new character
		minLen = min(minLen, right-left+1) // Update min length
	}

	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println(shortestUniqueSubstring("abcabcbb")) // Output: 1
	fmt.Println(shortestUniqueSubstring("aabcbc"))   // Output: 2
	fmt.Println(shortestUniqueSubstring("abcd"))     // Output: 4
}
