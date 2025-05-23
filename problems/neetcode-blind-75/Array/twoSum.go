package main

import "fmt"

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.

Example 1:

Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
Example 2:

Input: nums = [3,2,4], target = 6
Output: [1,2]
Example 3:

Input: nums = [3,3], target = 6
Output: [0,1]
*/

func main() {
	arr := []int{2, 7, 11, 15}
	fmt.Println(twoSum(arr, 9))
}
func twoSum(nums []int, target int) []int {
	n := len(nums)

	seen := make(map[int]int)

	for i := 0; i < n; i++ {
		compliment := target - nums[i]

		if j, ok := seen[compliment]; ok {
			return []int{i, j}
		}
		seen[nums[i]] = i
	}

	return []int{}
}
