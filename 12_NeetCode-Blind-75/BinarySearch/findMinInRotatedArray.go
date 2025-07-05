package main

import "fmt"

func main() {
	arr := []int{4, 5, 6, 7, 0, 1, 2}

	fmt.Printf("min number in given array of int is %d\n", findMin(arr))
}

func findMin(nums []int) int {

	left := 0
	right := len(nums) - 1

	for left <= right {

		fmt.Printf("left : %d\n", left)
		fmt.Printf("right : %d\n", right)

		mid := (left + right) / 2

		fmt.Printf("mid : %d\n", mid)

		fmt.Printf("num left : %d\n", nums[left])
		fmt.Printf("num right : %d\n", nums[right])
		fmt.Printf("num mid : %d\n", nums[mid])
		if nums[left] < nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return nums[left]
}

// min
