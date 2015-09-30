package automata

import (
	"fmt"
	"testing"
)

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
		if actual := test.stateSet.String(); actual != test.expected {
			t.Errorf("input: %s, actual: %s, expected: %s", test.stateSet, actual, test.expected)
		}
	}

	combinedStates := emptyState
	for _, test := range tests {
		combinedStates.Include(test.stateSet)
	}
	if combinedStates.String() != "{a,b,c,d}" {
		t.Errorf("Concat() expected: {a,b,c,d}, actual: %s", combinedStates.String())
	}
}

func TestTransitionTable(t *testing.T) {
	table := make(TransitionTable)
	var actual string

	actual = table.Row("a").Column('1').String()
	if actual != "{}" {
		t.Errorf("get() and states() should return defaults, actual: %s, expected: {}", actual)
	}
	table.Row("a").Column('1').Include(NewStateSet("b", "c"))
	actual = table.Row("a").Column('1').String()
	if actual != "{b,c}" {
		t.Errorf("add() should persist values, actual: %s, expected: {b,c}", actual)
	}
	table.Row("a").Column('1').Add("d")
	actual = table.Row("a").Column('1').String()
	if actual != "{b,c,d}" {
		t.Errorf("add() should persist values, actual: %s, expected: {b,c,d}", actual)
	}
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
