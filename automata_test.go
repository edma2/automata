package automata

import "testing"

func TestStateSet(t *testing.T) {
	emptyState := NewStateSet([]State{})
	var tests = []struct {
		stateSet *StateSet
		expected string
	}{
		{NewStateSet([]State{"a", "b", "b", "c"}), "{a,b,c}"},
		{NewStateSet([]State{"d", "d", "d", "d"}), "{d}"},
		{NewStateSet([]State{"b", "c", "a"}), "{a,b,c}"},
		{emptyState, "{}"},
	}

	for _, test := range tests {
		if actual := test.stateSet.String(); actual != test.expected {
			t.Errorf("input: %s, actual: %s, expected: %s", test.stateSet, actual, test.expected)
		}
	}

	combined := emptyState
	for _, test := range tests {
		combined = combined.Concat(test.stateSet)
	}
	if combined.String() != "{a,b,c,d}" {
		t.Errorf("Concat() expected: {a,b,c,d}, actual: %s", combined.String())
	}
}

//
// func TestChessboard(t *testing.T) {
// 	nfa := NewNFA("1", []State{"9"})
//
// 	nfa.NewTransition("1", 'r', []State{"2", "4"})
// 	nfa.NewTransition("2", 'r', []State{"4", "6"})
// 	nfa.NewTransition("3", 'r', []State{"2", "6"})
// 	nfa.NewTransition("4", 'r', []State{"2", "8"})
// 	nfa.NewTransition("5", 'r', []State{"2", "4", "6", "8"})
// 	nfa.NewTransition("6", 'r', []State{"2", "8"})
// 	nfa.NewTransition("7", 'r', []State{"4", "8"})
// 	nfa.NewTransition("8", 'r', []State{"4", "6"})
// 	nfa.NewTransition("9", 'r', []State{"6", "8"})
//
// 	nfa.NewTransition("1", 'b', []State{"5"})
// 	nfa.NewTransition("2", 'b', []State{"1", "3", "5"})
// 	nfa.NewTransition("3", 'b', []State{"5"})
// 	nfa.NewTransition("4", 'b', []State{"1", "5", "7"})
// 	nfa.NewTransition("5", 'b', []State{"1", "3", "7", "9"})
// 	nfa.NewTransition("6", 'b', []State{"3", "5", "9"})
// 	nfa.NewTransition("7", 'b', []State{"5"})
// 	nfa.NewTransition("8", 'b', []State{"5", "7", "9"})
// 	nfa.NewTransition("9", 'b', []State{"5"})
//
// 	dfa := nfa.Compile()
// 	if ok := dfa.Execute("rbb"); !ok {
// 		t.Error("rbb must be accepted")
// 	}
// }
