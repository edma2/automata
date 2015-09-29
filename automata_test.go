package automata

import "testing"

func TestStateSet(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"abbc", "{a,b,c}"},
		{"dddd", "{d}"},
		{"bca", "{a,b,c}"},
		{"", "{}"},
	}

	stateSets := make([]*StateSet, len(tests))

	for i, test := range tests {
		states := make([]State, len(test.input))
		for i, runeValue := range test.input {
			states[i] = State(runeValue)
		}
		stateSets[i] = NewStateSet(states)
	}

	for i, ss := range stateSets {
		test := tests[i]
		if actual := ss.String(); actual != test.expected {
			t.Errorf("input: %s, actual: %s, expected: %s", test.input, actual, test.expected)
		}
	}

	combined := NewStateSet([]State{})
	for _, ss := range stateSets {
		combined = combined.Combine(ss)
	}
	if combined.String() != "{a,b,c,d}" {
		t.Errorf("Combine() expected: {a,b,c,d}, actual: %s", combined.String())
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
