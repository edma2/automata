package automata_test

import (
	"testing"

	"github.com/edma2/coursera/automata"
)

func TestChessboard(t *testing.T) {
	// Chessboard
	nfa := automata.NewNFA("1", []automata.State{"9"})

	nfa.NewTransition("1", 'r', []automata.State{"2", "4"})
	nfa.NewTransition("2", 'r', []automata.State{"4", "6"})
	nfa.NewTransition("3", 'r', []automata.State{"2", "6"})
	nfa.NewTransition("4", 'r', []automata.State{"2", "8"})
	nfa.NewTransition("5", 'r', []automata.State{"2", "4", "6", "8"})
	nfa.NewTransition("6", 'r', []automata.State{"2", "8"})
	nfa.NewTransition("7", 'r', []automata.State{"4", "8"})
	nfa.NewTransition("8", 'r', []automata.State{"4", "6"})
	nfa.NewTransition("9", 'r', []automata.State{"6", "8"})

	nfa.NewTransition("1", 'b', []automata.State{"5"})
	nfa.NewTransition("2", 'b', []automata.State{"1", "3", "5"})
	nfa.NewTransition("3", 'b', []automata.State{"5"})
	nfa.NewTransition("4", 'b', []automata.State{"1", "5", "7"})
	nfa.NewTransition("5", 'b', []automata.State{"1", "3", "7", "9"})
	nfa.NewTransition("6", 'b', []automata.State{"3", "5", "9"})
	nfa.NewTransition("7", 'b', []automata.State{"5"})
	nfa.NewTransition("8", 'b', []automata.State{"5", "7", "9"})
	nfa.NewTransition("9", 'b', []automata.State{"5"})

	dfa := nfa.Compile()
	if ok := dfa.Execute("rbb"); !ok {
		t.Error("rbb must be accepted")
	}
}
