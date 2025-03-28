package main

import "fmt"

/*
Given a map where the key is a string and the value is a slice of strings, reverse the slices for each key.

Input:
map[string][]string{
    "a": {"x", "y", "z"},
    "b": {"1", "2"},
}

Output:

map[string][]string{
    "a": {"z", "y", "x"},
    "b": {"2", "1"},
}

*/

var sliceMap = map[string][]string{
	"a": {"x", "y", "z"},
	"b": {"1", "2"},
}

func main() {

	for k, v := range sliceMap {
		sliceMap[k] = reverseSlice(v)
	}
	fmt.Printf("new slice of %v\n", sliceMap)
}

func reverseSlice(arr []string) []string {

	n := len(arr)
	for i := 0; i < n/2; i++ {

		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
		/*	var temp = arr[i]
			arr[i] = arr[n-1-i]
			arr[n-1-i] = temp*/

	}
	fmt.Println(arr)
	return arr
}
