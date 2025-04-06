package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(arr))
}

func groupAnagrams(strs []string) [][]string {

	groupMap := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortedSlice(str)
		groupMap[sortedStr] = append(groupMap[sortedStr], str)
	}

	res := [][]string{}

	for _, group := range groupMap {
		res = append(res, group)
	}
	return res

}

func sortedSlice(str string) string {

	r := []rune(str)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return string(r)
}
