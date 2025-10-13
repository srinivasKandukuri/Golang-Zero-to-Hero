package main

import "fmt"

func insert(intervals [][]int, newInterval []int) [][]int {
	var results [][]int

	i := 0
	n := len(intervals)

	for i < n && intervals[i][0] <= newInterval[0] {
		results = append(results, intervals[i])
		i++
	}
	return results
}

func main() {
	intervals := [][]int{{1, 3}, {6, 9}}
	newInterval := []int{2, 5}
	fmt.Println("out ", insert(intervals, newInterval))
}
