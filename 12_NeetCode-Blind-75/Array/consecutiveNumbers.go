package main

import "fmt"

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
