package main

import "testing"

func TestStack_Peek(t *testing.T) {
	s := new(stack)
	s.Push("a")
	s.Push("b")
	s.Push("c")

	if s.Peek() != "c" {
		t.Error("Expected c, got", s.Peek())
	}
}

func TestStack_Pop(t *testing.T) {
	s := new(stack)
	s.Push("a")
	s.Push("b")
	s.Push("c")

	if s.Pop() != "c" {
		t.Error("Expected c, got", s.Pop())
	}

	if s.Pop() != "b" {
		t.Error("Expected b, got", s.Pop())
	}

	if s.Pop() != "a" {
		t.Error("Expected a, got", s.Pop())
	}

	if s.Pop() != "" {
		t.Error("Expected empty string, got", s.Pop())
	}
}

func TestStack_PeekEntireStack(t *testing.T) {
	s := new(stack)
	s.Push("a")
	s.Push("b")
	s.Push("c")

	if s.PeekEntireStack() != "abc" {
		t.Error("Expected abc, got", s.PeekEntireStack())
	}
}
