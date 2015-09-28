package main

import (
	"fmt"

	"github.com/edma2/coursera/automata/nfa"
)

func main() {
	// Chessboard
	a := nfa.NewNFA("1", []nfa.State{"9"})

	a.NewTransition("1", 'r', []nfa.State{"2", "4"})
	a.NewTransition("2", 'r', []nfa.State{"4", "6"})
	a.NewTransition("3", 'r', []nfa.State{"2", "6"})
	a.NewTransition("4", 'r', []nfa.State{"2", "8"})
	a.NewTransition("5", 'r', []nfa.State{"2", "4", "6", "8"})
	a.NewTransition("6", 'r', []nfa.State{"2", "8"})
	a.NewTransition("7", 'r', []nfa.State{"4", "8"})
	a.NewTransition("8", 'r', []nfa.State{"4", "6"})
	a.NewTransition("9", 'r', []nfa.State{"6", "8"})

	a.NewTransition("1", 'b', []nfa.State{"5"})
	a.NewTransition("2", 'b', []nfa.State{"1", "3", "5"})
	a.NewTransition("3", 'b', []nfa.State{"5"})
	a.NewTransition("4", 'b', []nfa.State{"1", "5", "7"})
	a.NewTransition("5", 'b', []nfa.State{"1", "3", "7", "9"})
	a.NewTransition("6", 'b', []nfa.State{"3", "5", "9"})
	a.NewTransition("7", 'b', []nfa.State{"5"})
	a.NewTransition("8", 'b', []nfa.State{"5", "7", "9"})
	a.NewTransition("9", 'b', []nfa.State{"5"})

	fmt.Println(a)
}
