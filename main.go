package main

import (
	"flag"
	"log"
	"strings"
)

var (
	g     grammar
	s     *stack
	a     *actions
	input string
)

func main() {
	grammarPtr := flag.String("g", "", "Grammar to use")
	inputPtr := flag.String("i", "", "Input string to parse")
	flag.Parse()

	if grammarPtr == nil || inputPtr == nil {
		log.Println("No Or Invalid arguments provided")
		return
	}

	input = *inputPtr

	gr := strings.Split(*grammarPtr, " ")

	if len(gr)%2 != 0 {
		log.Println("Invalid number of arguments")
		return
	}

	g = newGrammar()
	s = new(stack)
	a = newActions()

	g.Load(gr...)
	g.print()

	a.Add("", input, "shift")

	//i := 0

	for {
		if len(input) > 0 {
			s.Push(input[0:1])
		}

		r := tryReduce()
		if r && len(input) > 0 {
			a.Add(s.PeekEntireStack(), input[1:], "reduced")
		} else if len(input) > 0 {
			a.Add(s.PeekEntireStack(), input[1:], "shift")
		} else if len(input) == 0 {
			if len(s.PeekEntireStack()) == 1 {
				log.Println("Accepted")
			} else {
				log.Println("Rejected")
			}
			break
		}

		input = input[1:]

		/*i++

		if i == 4 {
			break
		}*/
	}

	a.Print()

	/*for {
		if a.stack == *inputPtr {
			log.Println("Accepted")
			break
		}
		if a.stack == "" {
			log.Println("Rejected")
			break
		}
		if a.action == "shift" {
			s.Push(a.input[0:1])
			a.input = a.input[1:]
			a.stack = s.PeekEntireStack()
			a.action = "reduce"
		} else {
			v := g.get(s.Peek())
			if v == "" {
				log.Println("Rejected")
				break
			}
			s.Pop()
			s.Push(v)
			a.stack = s.PeekEntireStack()
			a.action = "shift"
		}
		a.Add(a.stack, a.input, a.action)
	}*/
}

func tryReduce() bool {
	//log.Println("Trying to reduce")
	reduced := false
	for reduce() != 0 {
		reduced = true
		continue
	}

	return reduced
}

func reduce() int {
	longest := ""
	stackContents := s.PeekEntireStack()

	for i := len(stackContents) - 1; i >= 0; i-- {
		if v := g.get(stackContents[i:]); v != "" {
			if len(stackContents[i:]) > len(longest) {
				longest = stackContents[i:]
			}
		}
	}

	if longest == "" {
		return 0
	}

	for i := 0; i < len(longest); i++ {
		s.Pop()
	}

	s.Push(g.get(longest))
	a.Add(s.PeekEntireStack(), input, "reduced")

	return len(longest)
}
