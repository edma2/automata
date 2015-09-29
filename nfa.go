// Nondeterministic finite automatons (with ε-transitions)
package automata

import (
	"fmt"
	"sort"
	"strings"
)

// A State in the automaton is identified by a string.
// Usually they are human-readable UTF-8 encoded strings.
type State string

// A set of States
type StateSet struct {
	states map[State]bool
}

// An input symbol
type Symbol rune

// A NFA-ε is represented formally by a 5-tuple, (Q, Σ, Δ, q0, F), consisting of
//
// a finite set of states Q
// a finite set of input symbols Σ
// a transition function Δ : Q × (Σ ∪ {ε}) → P(Q)

type Step map[Symbol]*StateSet
type TransitionFunc map[State]Step

func (fn TransitionFunc) get(s State) Step {
	if fn[s] == nil {
		fn[s] = make(Step)
	}
	return fn[s]
}

func (step Step) states(input Symbol) *StateSet {
	if step[input] == nil {
		step[input] = NewStateSet()
	}
	return step[input]
}

func (step Step) add(input Symbol, newStates *StateSet) {
	step[input] = step.states(input).Concat(newStates)
}

// an initial (or start) state q0 ∈ Q
// a set of states F distinguished as accepting (or final) states F ⊆ Q.
type NFA struct {
	transitions TransitionFunc
	startState  State
	finalStates *StateSet
}

type DFA struct {
	NFA
}

// Create a new StateSet - silently ignore duplicates
func NewStateSet(states ...State) *StateSet {
	ss := new(StateSet)
	ss.states = make(map[State]bool)
	for _, s := range states {
		ss.states[s] = true
	}
	return ss
}

// Return states as an array (in no particular order)
func (ss *StateSet) States() []State {
	a := make([]State, len(ss.states))
	i := 0
	for s, _ := range ss.states {
		a[i] = s
		i = i + 1
	}
	return a
}

// Display a StateSet as a string (e.g. "{1,4,6}")
func (ss *StateSet) String() string {
	states := ss.States()
	a := make([]string, len(states))
	for i, s := range states {
		a[i] = string(s)
	}
	sort.Strings(a)
	return fmt.Sprintf("{%s}", strings.Join(a, ","))
}

// Return true if this StateSet contains at least one state also in other.
func (ss *StateSet) ContainsAny(other *StateSet) bool {
	for s, _ := range other.states {
		if ss.states[s] {
			return true
		}
	}
	return false
}

// Combine two StateSets, returning a new StateSet
func (ss *StateSet) Concat(other *StateSet) *StateSet {
	states := append(ss.States(), other.States()...)
	return NewStateSet(states...)
}

func NewNFA(startState State, finalStates *StateSet) *NFA {
	return &NFA{
		startState:  startState,
		finalStates: finalStates,
		transitions: make(TransitionFunc)}
}

func (nfa *NFA) Add(oldState State, input Symbol, newStates *StateSet) {
	nfa.transitions.get(oldState).add(input, newStates)
}

func (nfa *NFA) Compile() *DFA {
	dfa := new(DFA)
	dfa.transitions = make(TransitionFunc)
	dfa.finalStates = NewStateSet()
	powerSetConstruction(nfa, dfa, nil)
	// TODO: compute DFA final states
	return dfa
}

func powerSetConstruction(nfa *NFA, dfa *DFA, ss *StateSet) {
	// if nil, then we're starting from start state
	if ss == nil {
		startStateSet := NewStateSet(nfa.startState)
		dfa.startState = State(startStateSet.String())
		powerSetConstruction(nfa, dfa, startStateSet)
		return
	}
	dfaState := State(ss.String())
	if dfa.transitions[dfaState] != nil {
		return
	}
	if ss.ContainsAny(nfa.finalStates) {
		dfa.finalStates.states[dfaState] = true
	}
	unionStep := make(Step)
	for _, s := range ss.States() {
		for input, newStates := range nfa.transitions.get(s) {
			unionStep.add(input, newStates)
		}
	}
	dfa.transitions[dfaState] = unionStep
	for _, newStates := range unionStep {
		powerSetConstruction(nfa, dfa, newStates)
	}
}

func (dfa *DFA) String() string {
	var lines []string
	for state, step := range dfa.transitions {
		for input, newStates := range step {
			special := ""
			if dfa.finalStates.states[state] {
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
	return dfa.finalStates.states[state]
}
