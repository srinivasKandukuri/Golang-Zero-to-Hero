package main

import (
	"fmt"
)

func maxSumSubarray(arr []int, k int) int {
	window := 0
	n := len(arr)

	for i := 0; i < k; i++ {
		window += arr[i]
	}

	maxSum := window

	// Slide the window
	for i := k; i < n; i++ {
		window += arr[i] - arr[i-k] // Add new element, remove old element
		if window > maxSum {
			maxSum = window
		}
	}
	return maxSum
}

func maxSumBruteForce(arr []int, k int) int {

	n := len(arr)
	maxSum := 0

	for i := 0; i < n-k; i++ {

		currentSum := 0

		for j := i; j < i+k; j++ {
			currentSum += arr[j]
		}

		if maxSum < currentSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

func main() {

	arr := []int{2, 1, 5, 1, 3, 2}

	k := 3

	fmt.Println("Max sum of subarray of size", k, "is:", maxSumSubarray(arr, k))
}
