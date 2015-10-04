// Nondeterministic finite automatons (with ε-transitions)
package automata

import (
	"fmt"
	"sort"
	"strings"
)

// A state in the automaton is identified by a sequence of bytes.
type state string

// A set of states.
// States in a set are always unique.
type stateSet map[state]bool

// Fold this state set into a single state.
func (ss stateSet) singleton() state {
	return state(ss.String())
}

// Returns this state set as a string (e.g. "{1,4,6}").
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

// Returns true if this states set contains at least one state that
// is also in another state set.
func (ss stateSet) exists(other stateSet) bool {
	for q := range other {
		if ss[q] {
			return true
		}
	}
	return false
}

// Includes another state set's states to this one.
func (ss stateSet) union(other stateSet) {
	for q := range other {
		ss[q] = true
	}
}

// An input symbol.
// Automata input is always a sequence of runes (a Unicode string).
type symbol rune

// A transition table.
type ttab map[state]row

// A row in a transition table. It maps an input symbol to the next set of
// possible states.
type row map[symbol]stateSet

// Returns the row for a given state in this table.
func (tab ttab) row(q state) row {
	if tab[q] == nil {
		tab[q] = make(row)
	}
	return tab[q]
}

// Returns the column (a state set) for a given symbol in this row.
func (r row) col(a symbol) stateSet {
	if r[a] == nil {
		r[a] = make(stateSet)
	}
	return r[a]
}

type NFA struct {
	delta ttab
	q0    state
	final stateSet
}

// CL(q) = set of states you can reach from state q following only arcs labeled ε.
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
	for q := range nfa.delta.row(q).col('ε') {
		closure0(nfa, q, cl)
	}
}

// Returns a new NFA with given start state and final states.
func New(q0 state, finals ...state) *NFA {
	final := make(stateSet)
	for _, q := range finals {
		final[q] = true
	}
	return &NFA{q0: q0, final: final, delta: make(ttab)}
}

// Adds a new transition to this NFA.
func (nfa *NFA) Add(q state, a symbol, qs ...state) {
	ss := make(stateSet)
	for _, q := range qs {
		ss[q] = true
	}
	nfa.delta.row(q).col(a).union(ss)
}

// Compiles this NFA to a DFA-equivalent NFA.
func (nfa *NFA) Compile() *NFA {
	ss := make(stateSet)
	ss[nfa.q0] = true
	dfa := New(ss.singleton())
	powerset(nfa, dfa, ss)
	return dfa
}

// Implements the powerset construction algorithm.
func powerset(nfa *NFA, dfa *NFA, ss stateSet) {
	q := ss.singleton()
	if _, ok := dfa.delta[q]; ok {
		return // visited
	}
	if ss.exists(nfa.final) {
		dfa.final[q] = true
	}
	urow := make(row)
	for q := range ss {
		for a, next := range nfa.delta.row(q) {
			urow.col(a).union(next)
		}
	}
	for a, next := range urow {
		dfa.delta.row(q).col(a)[next.singleton()] = true
	}
	for _, next := range urow {
		powerset(nfa, dfa, next)
	}
}

func (dfa *NFA) String() string {
	var lines []string
	for q, r := range dfa.delta {
		mark := ""
		if dfa.final[q] {
			mark = mark + "*"
		}
		if dfa.q0 == q {
			mark = mark + ">"
		}
		for a, next := range r {
			lines = append(lines, fmt.Sprintf("%s		%s		%c	%s", mark, q, a, next))
		}
	}
	return strings.Join(lines, "\n")
}

// Get the one and only state in this state set.
func (ss stateSet) get1() state {
	for q := range ss {
		return q
	}
	return "" // not reached
}

func (dfa *NFA) Execute(input string) bool {
	q := dfa.q0
	for _, runeValue := range input {
		a := symbol(runeValue)
		q = dfa.delta[q][a].get1()
	}
	return dfa.final[q]
}
