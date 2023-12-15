package main

import "log"

// actions is a linked list of actions taken by the parser
type actions struct {
	stack  string
	input  string
	action string
	next   *actions
}

func newActions() *actions {
	return &actions{}
}

func (a *actions) Add(stack, input, action string) {
	if a.stack == "" {
		a.stack = stack
		a.input = input
		a.action = action
		return
	}
	a.next = newActions()
	a.next.Add(stack, input, action)
}

func (a *actions) Print() {
	if a == nil {
		return
	}
	log.Printf("|	%s	|	%s	|	%s	|\n", a.stack, a.input, a.action)
	log.Println("--------------------------------------------------")
	a.next.Print()
}
