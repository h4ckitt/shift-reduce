package main

import (
	"flag"
	"log"
	"strings"
)

var (
	g     grammar
	s     *stack
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

	g.Load(gr...)

	AddEventToList("", input, "shift")

	for {
		if len(input) > 0 {
			s.Push(input[0:1])
		}

		if tryReduce() && len(input) > 0 {
			AddEventToList(s.PeekEntireStack(), input[1:], "reduced")
		} else if len(input) > 0 {
			AddEventToList(s.PeekEntireStack(), input[1:], "shift")
		} else if len(input) == 0 {
			if s.PeekEntireStack() == s.Peek() {
				log.Println("Accepted")
			} else {
				log.Println("Rejected")
			}
			break
		}

		input = input[1:]
	}

	PrintList()
}

func tryReduce() bool {
	reduced := false
	for s.Reduce() != 0 {
		reduced = true
		continue
	}

	return reduced
}
