package main

import (
	"fmt"
)

/*
📌 Longest Substring Without Repeating Characters (Variable-Length Sliding Window)
🔹 Problem Statement
Given a string s, find the length of the longest substring that contains only unique (non-repeating) characters.

🔹 Constraints
s consists of only English letters (both lowercase and uppercase), digits, and symbols.
0 <= len(s) <= 10⁵
🔹 Example Walkthrough

Input: s = "abcabcbb"
Output: 3
Explanation: The longest substring without repeating characters is "abc", so the length is 3.

Input: s = "bbbbb"
Output: 1
Explanation: The longest substring without repeating characters is "b", so the length is 1.

Input: s = "pwwkew"
Output: 3
Explanation: The longest substring without repeating characters is "wke", so the length is 3.
(Note: "pwke" is not a valid answer because "w" repeats.)



loop the substring
create temp sub string and length
then push each value to substring
if the value contains the temp substring then set the max length of temp substring
then clear the substring
then continue the loop with new string


*/

func main() {
	fmt.Printf("max len %v", longestSubString("pwwkew"))

}

func longestSubString(str string) int {
	maxLen := 0
	left := 0
	seen := make(map[byte]bool)
	for right := 0; right < len(str); right++ {
		for seen[str[right]] {
			delete(seen, str[left])
			left++
		}
		seen[str[right]] = true
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
