// Package automata provides an interface for defining NFAs (with
// ε-transitions), compiling them to DFAs, and executing DFAs against
// input strings.
package automata

import (
	"fmt"
	"sort"
	"strings"
)

// A state is a string that uniquely identifies a state of an automaton.
type state string

// A stateSet implements a set of states.
type stateSet map[state]bool

// Singleton collapses sets of states into singleton states.
// Example: calling singleton on a set with states "a", "b", "c"
// gives you a new state "{a,b,c}".
func (ss stateSet) singleton() state {
	return state(ss.String())
}

func (ss stateSet) String() string {
	a := make([]string, len(ss))
	i := 0
	for q := range ss {
		a[i] = string(q)
		i = i + 1
	}
	sort.Strings(a)
	return fmt.Sprintf("{%s}", strings.Join(a, ","))
}

// Intersects compares two sets and returns true if the intersection
// is not empty.
func (ss stateSet) intersects(other stateSet) bool {
	for q := range other {
		if ss[q] {
			return true
		}
	}
	return false
}

// Union adds each state from another set.
func (ss stateSet) union(other stateSet) {
	for q := range other {
		ss[q] = true
	}
}

// A symbol is what automata inputs are composed of.
type symbol rune

// A ttab is a transition table in an automaton, mapping states to rows.
type ttab map[state]row

// A row maps an input symbol to the next set of possible states.
type row map[symbol]stateSet

// Row returns the row for a given state in the table. If the row
// doesn't exist, an empty one is created and returned.
func (tab ttab) row(q state) row {
	if tab[q] == nil {
		tab[q] = make(row)
	}
	return tab[q]
}

// Column returns the next set of possible states for an input symbol.
// If the set doesn't exist, an empty one is created and returned.
func (r row) col(a symbol) stateSet {
	if r[a] == nil {
		r[a] = make(stateSet)
	}
	return r[a]
}

// A Nondeterministic Finite Automaton with ε-transitions.  NFA's are
// not executable, only DFA's are (see below).
type NFA struct {
	delta ttab
	q0    state
	final stateSet
}

// A (read-only) Deterministic Finite Automaton which has been compiled
// from an NFA.  DFA's are executable, but can only be built using the
// NFA interface.
type DFA struct {
	nfa *NFA
}

// Closure of q is defined as the set of states you can reach from
// state q following only arcs labeled ε.
func closure(nfa *NFA, q state) stateSet {
	cl := make(stateSet)
	closure0(nfa, q, cl)
	return cl
}

func closure0(nfa *NFA, q state, cl stateSet) {
	if cl[q] {
		return
	}
	cl[q] = true
	if ss, ok := nfa.delta.row(q)['ε']; ok {
		for q := range ss {
			closure0(nfa, q, cl)
		}
	}
}

func rowUnion(nfa *NFA, ss stateSet) row {
	urow := make(row)
	for q := range ss {
		for a, next := range nfa.delta.row(q) {
			if a != 'ε' {
				urow.col(a).union(next)
			}
		}
	}
	return urow
}

func noEpsilons(nfa *NFA) *NFA {
	nfa2 := NewNFA(nfa.q0)
	for q, _ := range nfa.delta {
		cl := closure(nfa, q)
		nfa2.delta[q] = rowUnion(nfa, cl)
		if cl.intersects(nfa.final) {
			nfa2.final.union(cl)
		}
	}
	return nfa2
}

// NewNFA returns a new NFA with given start state and final states.
func NewNFA(q0 state, finals ...state) *NFA {
	final := make(stateSet)
	for _, q := range finals {
		final[q] = true
	}
	return &NFA{q0: q0, final: final, delta: make(ttab)}
}

// Add adds a new transition to the NFA.
func (nfa *NFA) Add(q state, a symbol, qs ...state) {
	ss := make(stateSet)
	for _, q := range qs {
		ss[q] = true
	}
	nfa.delta.row(q).col(a).union(ss)
}

// Compile compiles the NFA to a DFA.
func (nfa *NFA) Compile() *DFA {
	ss := make(stateSet)
	ss[nfa.q0] = true
	dfa := new(DFA)
	dfa.nfa = NewNFA(ss.singleton())
	powerset(noEpsilons(nfa), dfa.nfa, ss)
	return dfa
}

func powerset(nfa *NFA, dfa *NFA, ss stateSet) {
	q := ss.singleton()
	if _, ok := dfa.delta[q]; ok {
		return // visited
	}
	if ss.intersects(nfa.final) {
		dfa.final[q] = true
	}
	urow := rowUnion(nfa, ss)
	for a, next := range urow {
		dfa.delta.row(q).col(a)[next.singleton()] = true
	}
	for _, next := range urow {
		powerset(nfa, dfa, next)
	}
}

func (dfa *DFA) String() string {
	return dfa.nfa.String()
}

func (nfa *NFA) String() string {
	var lines []string
	for q, r := range nfa.delta {
		mark := ""
		if nfa.final[q] {
			mark = mark + "*"
		}
		if nfa.q0 == q {
			mark = mark + ">"
		}
		for a, next := range r {
			lines = append(lines, fmt.Sprintf("%s		%s		%c	%s", mark, q, a, next))
		}
	}
	return strings.Join(lines, "\n")
}

// Get1 returns the one and only state in the state set.
func (ss stateSet) get1() state {
	for q := range ss {
		return q
	}
	return "" // not reached
}

// Execute runs the DFA aginst an input string, and returns true if
// the string matches (i.e. is a member of the language defined by the
// automaton).
func (dfa *DFA) Execute(input string) bool {
	q := dfa.nfa.q0
	for _, runeValue := range input {
		a := symbol(runeValue)
		q = dfa.nfa.delta[q][a].get1()
	}
	return dfa.nfa.final[q]
}
