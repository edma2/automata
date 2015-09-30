// Nondeterministic finite automatons (with ε-transitions)
package automata

import (
	"fmt"
	"sort"
	"strings"
)

// A state in the automaton is identified by a sequence of bytes.
type State string

// A set of States.
// States in a StateSet are always unique.
type StateSet struct {
	states map[State]bool
}

// Returns a new state set, silently ignoring duplicate states.
func NewStateSet(states ...State) *StateSet {
	stateSet := new(StateSet)
	stateSet.states = make(map[State]bool)
	for _, state := range states {
		stateSet.states[state] = true
	}
	return stateSet
}

// Returns the states in this state set as an array.
// The states in the array are not in any particular order.
func (stateSet *StateSet) States() []State {
	a := make([]State, len(stateSet.states))
	i := 0
	for state, _ := range stateSet.states {
		a[i] = state
		i = i + 1
	}
	return a
}

func (stateSet *StateSet) Fold() State {
	return State(stateSet.String())
}

// Returns true if this state set contains state.
func (stateSet *StateSet) Contains(state State) bool {
	return stateSet.states[state]
}

// Returns this state set as a string (e.g. "{1,4,6}").
func (stateSet *StateSet) String() string {
	states := stateSet.States()
	a := make([]string, len(states))
	for i, state := range states {
		a[i] = string(state)
	}
	sort.Strings(a)
	return fmt.Sprintf("{%s}", strings.Join(a, ","))
}

// Returns true if this states set contains at least one state that
// is also in another state set.
func (stateSet *StateSet) ContainsAny(other *StateSet) bool {
	for state, _ := range other.states {
		if stateSet.Contains(state) {
			return true
		}
	}
	return false
}

// Includes another state set's states to this one.
func (stateSet *StateSet) Include(other *StateSet) {
	stateSet.Add(other.States()...)
}

func (stateSet *StateSet) Add(states ...State) {
	for _, state := range states {
		stateSet.states[state] = true
	}
}

// An input symbol.
// Automata input is always a sequence of runes (a Unicode string).
type Symbol rune

// A transition table.
type TransitionTable map[State]Row

// A row in a transition table. It maps an input symbol to the next set of
// possible states.
type Row map[Symbol]*StateSet

// Returns the row for a given state in this table.
func (table TransitionTable) Row(state State) Row {
	if table[state] == nil {
		table[state] = make(Row)
	}
	return table[state]
}

// Returns the column (a state set) for a given symbol in this row.
func (row Row) Column(input Symbol) *StateSet {
	if row[input] == nil {
		row[input] = NewStateSet()
	}
	return row[input]
}

type NFA struct {
	transitions TransitionTable
	startState  State
	finalStates *StateSet
}

type DFA struct {
	transitions map[State]map[Symbol]State
	startState  State
	finalStates *StateSet
}

// Returns a new NFA with given start state and final states.
func NewNFA(startState State, finalStates *StateSet) *NFA {
	return &NFA{
		startState:  startState,
		finalStates: finalStates,
		transitions: make(TransitionTable)}
}

// Adds a new transition to this NFA.
func (nfa *NFA) Add(oldState State, input Symbol, newStates *StateSet) {
	nfa.transitions.Row(oldState).Column(input).Include(newStates)
}

// Adds a new transition to this DFA.
func (dfa *DFA) Add(oldState State, input Symbol, newState State) {
	if dfa.transitions[oldState] == nil {
		dfa.transitions[oldState] = make(map[Symbol]State)
	}
	dfa.transitions[oldState][input] = newState
}

// Compiles this NFA to a DFA.
func (nfa *NFA) Compile() *DFA {
	dfa := new(DFA)
	dfa.transitions = make(map[State]map[Symbol]State)
	dfa.finalStates = NewStateSet()
	powersetConstruction(nfa, dfa, NewStateSet(nfa.startState))
	return dfa
}

// Implements the powerset construction algorithm.
func powersetConstruction(nfa *NFA, dfa *DFA, stateSet *StateSet) {
	dfaState := stateSet.Fold()
	if dfa.transitions[dfaState] != nil {
		return
	}
	if stateSet.ContainsAny(nfa.finalStates) {
		dfa.finalStates.Add(dfaState)
	}
	if dfa.startState == "" {
		dfa.startState = dfaState
	}
	union := make(Row)
	for _, state := range stateSet.States() {
		for input, newStates := range nfa.transitions.Row(state) {
			union.Column(input).Include(newStates)
		}
	}
	for input, newStates := range union {
		dfa.Add(dfaState, input, newStates.Fold())
	}
	for _, newStates := range union {
		powersetConstruction(nfa, dfa, newStates)
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
			lines = append(lines, fmt.Sprintf("%s		%s		%c	%s", special, state, input, newStates))
		}
	}
	return strings.Join(lines, "\n")
}

func (dfa *DFA) Execute(input string) bool {
	state := dfa.startState
	for _, runeValue := range input {
		state = dfa.transitions[state][Symbol(runeValue)]
	}
	return dfa.finalStates.Contains(state)
}
