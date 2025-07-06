package main

import "fmt"

// find the missing elements in sequence
// 1. check next element+1 - current element diff =1 then it is sequence
// else the add that element

func main() {
	arr := []int{1, 7, 3, 5, 7}
	var arr1 []int
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == 1 {
			continue
		} else {
			arr1 = append(arr1, arr[i]+1)
		}
	}

	fmt.Printf("array is %v", arr1)
}
