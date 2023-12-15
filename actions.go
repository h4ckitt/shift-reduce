package main

import "log"

// actions is a linked list of actions taken by the parser
type actions struct {
	stack  string
	input  string
	action string
	next   *actions
}

var (
	head    *actions
	current *actions
)

func AddEventToList(stack, input, action string) {
	if head == nil {
		head = &actions{stack, input, action, nil}
		current = head
		return
	}
	current.next = &actions{stack, input, action, nil}
	current = current.next
}

func PrintList() {
	head.Print()
}

func (a *actions) Print() {
	if a == nil {
		return
	}

	log.Printf("|	%s	|	%s	|	%s	|\n", a.stack, a.input, a.action)
	log.Println("---------------------------------------------")
	a.next.Print()
}
