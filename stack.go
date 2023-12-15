package main

import "strings"

type stack []string

// Push pushes a value onto the stack
func (s *stack) Push(v string) {
	*s = append(*s, v)
}

// Pop returns the top value from the stack
func (s *stack) Pop() string {
	l := len(*s)
	if l == 0 {
		return ""
	}
	res := (*s)[l-1]
	*s = (*s)[:l-1]
	return res
}

// Peek returns the top value from the stack without removing it
func (s *stack) Peek() string {
	l := len(*s)
	if l == 0 {
		return ""
	}
	return (*s)[l-1]
}

// PeekEntireStack returns the entire stack as a string
func (s *stack) PeekEntireStack() string {
	return strings.Join(*s, "")
}
