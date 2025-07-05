package main

/*
1. Count Elements in a Slice
Write a function that takes a slice of integers and returns a map where the key is the integer and the value is how many times it appears.

Input:
[]int{1, 2, 2, 3, 3, 3}


Output:
map[int]int{
    1: 1,
    2: 2,
    3: 3,
}
*/
import (
	"fmt"
	"sort"
)

func main() {
	val := []int{1, 2, 2, 3, 3, 4, 5, 5, 6, 7, 8, 9, 10, 10}

	var countMap = make(map[int]int)

	for _, v := range val {
		countMap[v]++ // this line equals to below lines
		// in go map will take care of the init, if the value doesn't exist then it will create 0 default
		/*

			if _, found := countMap[v]; found {
				countMap[v] += 1
			} else {
				countMap[v] = 1
			}

		*/
	}

	var keys []int
	for k, _ := range countMap {
		keys = append(keys, k)
	}

	fmt.Println("before sort ", keys)
	sort.Ints(keys)
	fmt.Println("after sort", keys)

	for _, k := range keys {
		if countMap[k] > 1 {
			fmt.Printf("%d repeated %d times\n", k, countMap[k])
		}
	}

	fmt.Println(countMap)
}

/*
before sort  [2 3 4 6 7 8 9 1 10 5]
after sort [1 2 3 4 5 6 7 8 9 10]
2 repeated 2 times
3 repeated 2 times
5 repeated 2 times
10 repeated 2 times
map[1:1 2:2 3:2 4:1 5:2 6:1 7:1 8:1 9:1 10:2]
*/
