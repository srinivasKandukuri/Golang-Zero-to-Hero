package main

// without generics
type IntStack struct {
	data []int
}

func (s *IntStack) Push(x int) {
	s.data = append(s.data, x)
}

func (s *IntStack) Pop() int {
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}

// with generics
type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(x T) {
	s.data = append(s.data, x)
}

func (s *Stack[T]) Pop() T {
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}
