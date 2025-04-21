package main

import "fmt"

func main() {
	arr := "()[]{}"
	fmt.Println(validParentheses(arr))
}

func validParentheses(s string) bool {

	var stack []rune

	mapping := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	for _, ch := range s {
		if open, exist := mapping[ch]; exist {
			if len(stack) == 0 || stack[len(stack)-1] != open {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, ch)
		}
	}

	return len(stack) == 0
}
