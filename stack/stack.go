package stack

import (
	"errors"
)

// Stack represents a stack that holds a slice
type Stack struct {
	items []int
}

//Size function to get size of stack
func (s Stack) Size() int {
	return len(s.items)
}

// Check if Stack is empty
func (s Stack) isEmpty() bool {
	return len(s.items) == 0
}

// Peek function to peek the top item
func (s Stack) Peek() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// Push function to append items in stack
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop function to pop top item in stack
func (s *Stack) Pop() (int, error) {
	if s.isEmpty() {
		return 0, errors.New("stack is empty")
	}

	itemRemoved := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return itemRemoved, nil
}
