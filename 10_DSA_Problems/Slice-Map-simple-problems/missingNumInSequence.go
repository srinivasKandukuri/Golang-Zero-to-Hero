package main

import "fmt"

func main() {
	arr := []int{1, 7, 3, 5, 7}
	arr1 := []int{}
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1]-arr[i] == 1 {
			continue
		} else {
			arr1 = append(arr1, arr[i]+1)
		}
	}

	fmt.Printf("array is %v", arr1)
}
