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

type TransitionFunc map[State]map[Symbol]*StateSet

// an initial (or start) state q0 ∈ Q
// a set of states F distinguished as accepting (or final) states F ⊆ Q.
type NFA struct {
	transitions TransitionFunc
	startState  State
	finalStates *StateSet
}

type DFA struct {
	transitions TransitionFunc
	startStates *StateSet
	finalStates *StateSet
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
	if nfa.transitions[oldState] == nil {
		nfa.transitions[oldState] = make(map[Symbol]*StateSet)
	}
	nfa.transitions[oldState][input] = newStates
}

func (nfa *NFA) Compile() *DFA {
	dfa := new(DFA)
	dfa.transitions = make(TransitionFunc)
	dfa.startStates = NewStateSet(nfa.startState)
	powerSetConstruction(nfa, dfa, dfa.startStates)
	// TODO: compute DFA final states
	return dfa
}

func powerSetConstruction(nfa *NFA, dfa *DFA, ss *StateSet) {
	for _, s := range ss.States() {
		for input, newStates := range nfa.transitions[s] {
			dfaState := State(newStates.String())
			if dfa.transitions[dfaState] == nil {
				dfa.transitions[dfaState] = make(map[Symbol]*StateSet)
			}
			existingNewStates := dfa.transitions[dfaState][input]
			if existingNewStates == nil {
				dfa.transitions[dfaState][input] = newStates
			} else {
				dfa.transitions[dfaState][input] = existingNewStates.Concat(newStates)
			}
		}
	}
}

func (dfa *DFA) Execute(input string) bool {
	return false
}
