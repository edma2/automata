package automata

import "testing"

var uniqTests = []struct {
	input    []string
	expected []string
}{
	{[]string{"a", "b", "b", "c"}, []string{"a", "b", "c"}},
	{[]string{"a", "a", "a", "a"}, []string{"a"}},
	{[]string{"a", "b", "c"}, []string{"a", "b", "c"}},
	{[]string{}, []string{}},
}

func TestUniq(t *testing.T) {
	for _, test := range uniqTests {
		actual := uniq(test.input)
		if len(actual) != len(test.expected) {
			t.Errorf("input: %s, actual: %s, expected: %s", test.input, actual, test.expected)
		}
		for i, s := range test.expected {
			if actual[i] != s {
				t.Errorf("input: %s, actual: %s, expected: %s", test.input, actual, test.expected)
			}
		}
	}
}

func TestChessboard(t *testing.T) {
	nfa := NewNFA("1", []State{"9"})

	nfa.NewTransition("1", 'r', []State{"2", "4"})
	nfa.NewTransition("2", 'r', []State{"4", "6"})
	nfa.NewTransition("3", 'r', []State{"2", "6"})
	nfa.NewTransition("4", 'r', []State{"2", "8"})
	nfa.NewTransition("5", 'r', []State{"2", "4", "6", "8"})
	nfa.NewTransition("6", 'r', []State{"2", "8"})
	nfa.NewTransition("7", 'r', []State{"4", "8"})
	nfa.NewTransition("8", 'r', []State{"4", "6"})
	nfa.NewTransition("9", 'r', []State{"6", "8"})

	nfa.NewTransition("1", 'b', []State{"5"})
	nfa.NewTransition("2", 'b', []State{"1", "3", "5"})
	nfa.NewTransition("3", 'b', []State{"5"})
	nfa.NewTransition("4", 'b', []State{"1", "5", "7"})
	nfa.NewTransition("5", 'b', []State{"1", "3", "7", "9"})
	nfa.NewTransition("6", 'b', []State{"3", "5", "9"})
	nfa.NewTransition("7", 'b', []State{"5"})
	nfa.NewTransition("8", 'b', []State{"5", "7", "9"})
	nfa.NewTransition("9", 'b', []State{"5"})

	dfa := nfa.Compile()
	if ok := dfa.Execute("rbb"); !ok {
		t.Error("rbb must be accepted")
	}
}
