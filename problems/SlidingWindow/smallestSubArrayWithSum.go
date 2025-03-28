package main

import (
	"fmt"
)

/*
Problem Statement
Given an array of positive integers arr and a target sum S, find the length of the smallest contiguous subarray whose sum is greater than or equal to S.
If no such subarray exists, return 0.

Constraints
1 ≤ len(arr) ≤ 10⁵
1 ≤ arr[i] ≤ 10⁴
1 ≤ S ≤ 10⁹


Example 1
Input:
arr = []int{2, 3, 1, 2, 4, 3}
S = 7

Output:
2

Example 2
Input:
arr = []int{1, 4, 4}
S = 4

Output:
1
*/

func main() {

	arr := []int{2, 3, 1, 2, 4, 3}
	S := 7
	fmt.Printf("smallest sub array with sum %d", smallestSubArraySumOpt(arr, S))
}

func smallestSubArraySum(arr []int, s int) int {

	n := len(arr)
	minLength := n + 1
	for i := 0; i < n; i++ {
		tempSum := 0
		for j := i; j < n; j++ {
			tempSum += arr[j]
			if tempSum >= s {
				if j-i+1 <= minLength {
					minLength = j - i + 1
				}
				break
			}
		}
	}
	if minLength == n+1 {
		return 0
	}

	return minLength
}

func smallestSubArraySumOpt(arr []int, s int) int {
	n := len(arr)
	minLength := n + 1

	left := 0
	sum := 0
	for right := 0; right < n; right++ {
		sum += arr[right]
		for sum >= s {
			if right-left+1 < minLength {
				minLength = right - left + 1
			}
			sum -= arr[left]
			left++
		}
	}
	if minLength == n+1 {
		return 0
	}
	return minLength
}
