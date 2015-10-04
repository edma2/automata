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

func TestStrings(t *testing.T) {
	var tests = []struct {
		qs       []state
		expected string
	}{
		{[]state{"a", "b", "b", "c"}, "{a,b,c}"},
		{[]state{"d", "d", "d", "d"}, "{d}"},
		{[]state{"b", "c", "a"}, "{a,b,c}"},
		{[]state{}, "{}"},
	}

	union := make(stateSet)
	for _, test := range tests {
		ss := make(stateSet)
		for _, q := range test.qs {
			ss[q] = true
		}
		union.union(ss)
		testString(t, ss, test.expected)
	}
	testString(t, union, "{a,b,c,d}")
}

func TestTable(t *testing.T) {
	tab := make(ttab)

	testString(t, tab.row("a").col('1'), "{}")

	ss := make(stateSet)
	ss["b"] = true
	ss["c"] = true
	tab.row("a").col('1').union(ss)
	testString(t, tab.row("a").col('1'), "{b,c}")

	tab.row("a").col('1')["d"] = true
	testString(t, tab.row("a").col('1'), "{b,c,d}")
}

func TestChessboard(t *testing.T) {
	nfa := New("1", "9")

	nfa.Add("1", 'r', "2", "4")
	nfa.Add("2", 'r', "4", "6")
	nfa.Add("3", 'r', "2", "6")
	nfa.Add("4", 'r', "2", "8")
	nfa.Add("5", 'r', "2", "4", "6", "8")
	nfa.Add("6", 'r', "2", "8")
	nfa.Add("7", 'r', "4", "8")
	nfa.Add("8", 'r', "4", "6")
	nfa.Add("9", 'r', "6", "8")

	nfa.Add("1", 'b', "5")
	nfa.Add("2", 'b', "1", "3", "5")
	nfa.Add("3", 'b', "5")
	nfa.Add("4", 'b', "1", "5", "7")
	nfa.Add("5", 'b', "1", "3", "7", "9")
	nfa.Add("6", 'b', "3", "5", "9")
	nfa.Add("7", 'b', "5")
	nfa.Add("8", 'b', "5", "7", "9")
	nfa.Add("9", 'b', "5")

	dfa := nfa.Compile()
	//fmt.Println(dfa)
	if ok := dfa.Execute("rbb"); !ok {
		t.Error("rbb must be accepted")
	}
}

func TestClosure(t *testing.T) {
	nfa := New("A", "D")

	nfa.Add("A", '0', "E")
	nfa.Add("A", '1', "B")
	nfa.Add("B", '1', "C")
	nfa.Add("B", 'ε', "D")
	nfa.Add("C", '1', "D")
	nfa.Add("E", '0', "F")
	nfa.Add("E", 'ε', "B", "C")
	nfa.Add("F", '0', "D")

	testString(t, closure(nfa, "A"), "{A}")
	testString(t, closure(nfa, "E"), "{B,C,D,E}")

	nfa2 := noEpsilons(nfa)
	fmt.Println(nfa2)
}
