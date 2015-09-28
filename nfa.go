// Nondeterministic finite automatons (with ε-transitions)
package automata

import (
	"fmt"
	"sort"
	"strings"
)

// A NFA-ε is represented formally by a 5-tuple, (Q, Σ, Δ, q0, F), consisting of
//
// a finite set of states Q
type State string

// a finite set of input symbols Σ
type Symbol rune

// a transition function Δ : Q × (Σ ∪ {ε}) → P(Q)
type Transition map[Symbol][]State
type TransitionFunc map[State]Transition

// TODO: ε-transitions

// an initial (or start) state q0 ∈ Q
// a set of states F distinguished as accepting (or final) states F ⊆ Q.
type NFA struct {
	q0    State
	final map[State]bool
	delta TransitionFunc
}

func NewNFA(q0 State, final []State) *NFA {
	nfa := NFA{q0: q0, delta: make(TransitionFunc)}
	nfa.final = make(map[State]bool)
	for i := 0; i < len(final); i++ {
		nfa.final[final[i]] = true
	}
	return &nfa
}

func (nfa *NFA) NewTransition(old State, input Symbol, new []State) {
	transition := nfa.delta[old]
	if transition == nil {
		transition = make(Transition)
		nfa.delta[old] = transition
	}
	transition[input] = new
}

// Return a copy of a without duplicate strings.
// Assumes a is already sorted.
func uniq(a []string) []string {
	newa := make([]string, len(a))
	i := -1
	for _, s := range a {
		if i < 0 || newa[i] != s {
			i = i + 1
			newa[i] = s
		}
	}
	if i < 0 {
		return []string{}
	} else {
		return newa[0 : i+1]
	}
}

func nfaToDfaState(states []State) State {
	strStates := make([]string, len(states))
	for i := 0; i < len(states); i++ {
		strStates[i] = string(states[i])
	}
	sort.Strings(strStates)
	return State(fmt.Sprintf("{%s}", strings.Join(uniq(strStates), ",")))
}

// Compile this NFA to an executable DFA.
func (dfa *NFA) Compile() *DFA {
	return nil
}
