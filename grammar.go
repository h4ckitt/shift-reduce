package main

import (
	"fmt"
	"strings"
)

type grammar struct {
	substitutions map[string]string
	terminals     map[string]struct{}
}

func newGrammar() grammar {
	return grammar{
		substitutions: make(map[string]string),
		terminals:     make(map[string]struct{}),
	}
}

// Load loads the space separated key-value pairs into the grammar
// Example grammar Is: S -> aS | b
func (g grammar) Load(keyval ...string) {

	if len(keyval)%2 != 0 {
		panic("Invalid number of arguments")
	}
	for i := 0; i < len(keyval); i += 2 {
		values := strings.Split(keyval[i+1], "|")
		for _, val := range values {
			g.substitutions[strings.TrimSpace(val)] = keyval[i]
			g.terminals[keyval[i]] = struct{}{}
		}
	}
}

// Get returns the value for the given key
func (g grammar) get(key string) string {
	if v, ok := g.substitutions[key]; ok {
		return v
	}
	return ""
}

// Print prints the grammar
func (g grammar) print() {
	for key, val := range g.substitutions {
		fmt.Printf("%s -> %s\n", key, val)
	}
}
