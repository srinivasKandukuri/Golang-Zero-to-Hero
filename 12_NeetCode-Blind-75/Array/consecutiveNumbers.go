package main

import "fmt"

/*
128. Longest Consecutive Sequence - Explanation
https://neetcode.io/solutions/longest-consecutive-sequence

Given an array of integers nums, return the length of the longest consecutive sequence of elements that can be formed.

A consecutive sequence is a sequence of elements in which each element is exactly 1 greater than the previous element. The elements do not have to be consecutive in the original array.

You must write an algorithm that runs in O(n) time.

Example 1:

Input: nums = [2,20,4,10,3,4,5]

Output: 4
Explanation: The longest consecutive sequence is [2, 3, 4, 5].

Example 2:

Input: nums = [0,3,2,5,4,6,1,1]

Output: 7
Constraints:

0 <= nums.length <= 1000
-10^9 <= nums[i] <= 10^9
*/

func main() {
	arr := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(arr))
}
func longestConsecutive(nums []int) int {
	longest := 0

	set := make(map[int]struct{}, len(nums))

	for _, n := range nums {
		set[n] = struct{}{}
	}

	for n := range set {
		if _, exist := set[n-1]; exist {
			continue
		}

		num := n
		streak := 1

		for {
			_, ok := set[num+1]
			if !ok {
				break
			}
			num++
			streak++
		}

		if streak > longest {
			longest = streak
		}
	}

	return longest
}
