package main

import "fmt"

func main() {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(arr))
}

// brute force
func containerOfWater(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {

		for j := i + 1; j < len(heights); j++ {
			h := min(heights[i], heights[j])
			w := j - i
			area := h * w
			if area > maxArea {
				maxArea = area
			}
		}
	}
	return maxArea
}

func maxArea(heights []int) int {

	left := 0
	right := len(heights) - 1

	maxArea := 0

	for left < right {
		h := min(heights[left], heights[right])
		w := right - left

		area := h * w
		if area > maxArea {
			maxArea = area
		}

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}

	}
	return maxArea
}

func min(a, b int) int {

	if a < b {
		return a
	}
	return b
}
