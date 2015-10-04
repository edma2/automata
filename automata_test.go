package automata

import (
	"fmt"
	"testing"
)

func testString(t *testing.T, actual fmt.Stringer, expected string) {
	if actual.String() != expected {
		t.Errorf("expected %s, got %s", expected, actual)
	}
}

func TestStateSet(t *testing.T) {
	emptyState := NewStateSet()
	var tests = []struct {
		stateSet *StateSet
		expected string
	}{
		{NewStateSet("a", "b", "b", "c"), "{a,b,c}"},
		{NewStateSet("d", "d", "d", "d"), "{d}"},
		{NewStateSet("b", "c", "a"), "{a,b,c}"},
		{emptyState, "{}"},
	}

	for _, test := range tests {
		testString(t, test.stateSet, test.expected)
	}

	combinedStates := emptyState
	for _, test := range tests {
		combinedStates.Include(test.stateSet)
	}
	testString(t, combinedStates, "{a,b,c,d}")
}

func TestTransitionTable(t *testing.T) {
	table := make(TransitionTable)

	testString(t, table.Row("a").Column('1'), "{}")

	table.Row("a").Column('1').Include(NewStateSet("b", "c"))
	testString(t, table.Row("a").Column('1'), "{b,c}")

	table.Row("a").Column('1').Add("d")
	testString(t, table.Row("a").Column('1'), "{b,c,d}")
}

func TestChessboard(t *testing.T) {
	nfa := NewNFA("1", NewStateSet("9"))

	nfa.Add("1", 'r', NewStateSet("2", "4"))
	nfa.Add("2", 'r', NewStateSet("4", "6"))
	nfa.Add("3", 'r', NewStateSet("2", "6"))
	nfa.Add("4", 'r', NewStateSet("2", "8"))
	nfa.Add("5", 'r', NewStateSet("2", "4", "6", "8"))
	nfa.Add("6", 'r', NewStateSet("2", "8"))
	nfa.Add("7", 'r', NewStateSet("4", "8"))
	nfa.Add("8", 'r', NewStateSet("4", "6"))
	nfa.Add("9", 'r', NewStateSet("6", "8"))

	nfa.Add("1", 'b', NewStateSet("5"))
	nfa.Add("2", 'b', NewStateSet("1", "3", "5"))
	nfa.Add("3", 'b', NewStateSet("5"))
	nfa.Add("4", 'b', NewStateSet("1", "5", "7"))
	nfa.Add("5", 'b', NewStateSet("1", "3", "7", "9"))
	nfa.Add("6", 'b', NewStateSet("3", "5", "9"))
	nfa.Add("7", 'b', NewStateSet("5"))
	nfa.Add("8", 'b', NewStateSet("5", "7", "9"))
	nfa.Add("9", 'b', NewStateSet("5"))

	dfa := nfa.Compile()
	fmt.Println(dfa)
	if ok := dfa.Execute("rbb"); !ok {
		t.Error("rbb must be accepted")
	}
}

func TestClosure(t *testing.T) {
	nfa := NewNFA("A", NewStateSet("D"))

	nfa.Add("A", '0', NewStateSet("E"))
	nfa.Add("A", '1', NewStateSet("B"))
	nfa.Add("B", '1', NewStateSet("C"))
	nfa.Add("B", 'ε', NewStateSet("D"))
	nfa.Add("C", '1', NewStateSet("D"))
	nfa.Add("E", '0', NewStateSet("F"))
	nfa.Add("E", 'ε', NewStateSet("B", "C"))
	nfa.Add("F", '0', NewStateSet("D"))

	testString(t, closure(nfa, "A"), "{A}")
	testString(t, closure(nfa, "E"), "{B,C,D,E}")
}
