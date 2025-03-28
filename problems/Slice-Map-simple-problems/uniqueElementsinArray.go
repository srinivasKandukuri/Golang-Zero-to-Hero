package main

import "fmt"

/*
3. Unique Elements in a Slice
Given a slice of integers, return a map where the key is the unique element and the value is true.

Input:
[]int{1, 2, 2, 3}


Output:
map[int]bool{
    1: true,
    2: true,
    3: true,
}
*/

func main() {

	arr := []int{1, 2, 3, 4, 4, 2}
	uniqueMap := make(map[int]struct{})

	for _, v := range arr {
		uniqueMap[v] = struct{}{}
	}
	fmt.Printf("%v values", uniqueMap)
}

/*

time complexity : O(n)
	best case O(N)
	worst case O(n)

space complexity : O(n)



*/
