// Nondeterministic finite automatons (with ε-transitions)
package automata

import (
	"fmt"
	"sort"
	"strings"
)

// A State in the automaton is identified by a human-readable UTF-8
// encoded string.
type State string

// A set of States
// States in a StateSet are always unique.
type StateSet struct {
	states map[State]bool
}

// Create a new StateSet - silently ignore duplicates
func NewStateSet(states ...State) *StateSet {
	stateSet := new(StateSet)
	stateSet.states = make(map[State]bool)
	for _, state := range states {
		stateSet.states[state] = true
	}
	return stateSet
}

// Return states as an array (in no particular order)
func (stateSet *StateSet) States() []State {
	a := make([]State, len(stateSet.states))
	i := 0
	for state, _ := range stateSet.states {
		a[i] = state
		i = i + 1
	}
	return a
}

func (stateSet *StateSet) Contains(state State) bool {
	return stateSet.states[state]
}

// Display a StateSet as a string (e.g. "{1,4,6}")
func (stateSet *StateSet) String() string {
	states := stateSet.States()
	a := make([]string, len(states))
	for i, state := range states {
		a[i] = string(state)
	}
	sort.Strings(a)
	return fmt.Sprintf("{%s}", strings.Join(a, ","))
}

// Return true if this StateSet contains at least one state also in other.
func (stateSet *StateSet) ContainsAny(other *StateSet) bool {
	for state, _ := range other.states {
		if stateSet.Contains(state) {
			return true
		}
	}
	return false
}

// Combine two StateSets, returning a new StateSet
func (stateSet *StateSet) Concat(other *StateSet) *StateSet {
	states := append(stateSet.States(), other.States()...)
	return NewStateSet(states...)
}

// An input symbol
type Symbol rune

// A NFA-ε is represented formally by a 5-tuple, (Q, Σ, Δ, q0, F), consisting of
//
// a finite set of states Q
// a finite set of input symbols Σ
// a transition function Δ : Q × (Σ ∪ {ε}) → P(Q)

type Row map[Symbol]*StateSet
type TransitionTable map[State]Row

func (table TransitionTable) get(state State) Row {
	if table[state] == nil {
		table[state] = make(Row)
	}
	return table[state]
}

func (row Row) states(input Symbol) *StateSet {
	if row[input] == nil {
		row[input] = NewStateSet()
	}
	return row[input]
}

// TODO: move to StateSet?
func (row Row) add(input Symbol, newStates *StateSet) {
	row[input] = row.states(input).Concat(newStates)
}

// an initial (or start) state q0 ∈ Q
// a set of states F distinguished as accepting (or final) states F ⊆ Q.
type NFA struct {
	transitions TransitionTable
	startState  State
	finalStates *StateSet
}

type DFA struct {
	NFA
}

func NewNFA(startState State, finalStates *StateSet) *NFA {
	return &NFA{
		startState:  startState,
		finalStates: finalStates,
		transitions: make(TransitionTable)}
}

func (nfa *NFA) Add(oldState State, input Symbol, newStates *StateSet) {
	nfa.transitions.get(oldState).add(input, newStates)
}

func (nfa *NFA) Compile() *DFA {
	dfa := new(DFA)
	dfa.transitions = make(TransitionTable)
	dfa.finalStates = NewStateSet()
	powerSetConstruction(nfa, dfa, nil)
	// TODO: compute DFA final states
	return dfa
}

func powerSetConstruction(nfa *NFA, dfa *DFA, stateSet *StateSet) {
	// if nil, then we're starting from start state
	if stateSet == nil {
		startStateSet := NewStateSet(nfa.startState)
		dfa.startState = State(startStateSet.String()) // TODO: stateSet.FoldState()?
		powerSetConstruction(nfa, dfa, startStateSet)
		return
	}
	dfaState := State(stateSet.String())
	if dfa.transitions[dfaState] != nil {
		return
	}
	if stateSet.ContainsAny(nfa.finalStates) {
		dfa.finalStates.states[dfaState] = true
	}
	unionRow := make(Row)
	for _, state := range stateSet.States() {
		for input, newStates := range nfa.transitions.get(state) {
			unionRow.add(input, newStates)
		}
	}
	dfa.transitions[dfaState] = unionRow
	for _, newStates := range unionRow {
		powerSetConstruction(nfa, dfa, newStates)
	}
}

func (dfa *DFA) String() string {
	var lines []string
	for state, row := range dfa.transitions {
		for input, newStates := range row {
			special := ""
			if dfa.finalStates.Contains(state) {
				special = special + "*"
			}
			if dfa.startState == state {
				special = special + ">"
			}
			lines = append(lines, fmt.Sprintf("%s %s	%c	%s", special, state, input, newStates))
		}
	}
	return strings.Join(lines, "\n")
}

func (dfa *DFA) Execute(input string) bool {
	state := dfa.startState
	for _, runeValue := range input {
		newStates := dfa.transitions[state][Symbol(runeValue)]
		state = State(newStates.String())
	}
	return dfa.finalStates.Contains(state)
}
