package main

import "testing"

func TestGrammar_Load(t *testing.T) {
	g := newGrammar()
	g.Load("a", "b|c|e+j", "d", "e|f")

	if g.get("b") != "a" {
		t.Error("Expected a, got", g.get("b"))
	}

	if g.get("c") != "a" {
		t.Error("Expected a, got", g.get("c"))
	}

	if g.get("e+j") != "a" {
		t.Error("Expected a, got", g.get("e+j"))
	}

	if g.get("e") != "d" {
		t.Error("Expected d, got", g.get("e"))
	}

	if g.get("f") != "d" {
		t.Error("Expected d, got", g.get("f"))
	}
}
