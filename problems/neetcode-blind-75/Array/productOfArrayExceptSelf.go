package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(arr))
}

func productExceptSelf(nums []int) []int {

	n := len(nums)
	results := make([]int, n)

	// first pass

	prefix := 1

	for i := 0; i < n; i++ {
		results[i] = prefix
		prefix *= nums[i]
	}

	postfix := 1
	// second pass
	for i := n - 1; i >= 0; i-- {
		results[i] *= postfix
		postfix *= nums[i]
	}

	return results
}
