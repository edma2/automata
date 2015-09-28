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

func (nfa *NFA) NewTransition(q State, input Symbol, qs []State) {
	transition := nfa.delta[q]
	if transition == nil {
		transition = make(Transition)
		nfa.delta[q] = transition
	}
	transition[input] = qs
}

// Return a copy of sorted slice without duplicate strings.
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

func nfaToDfaState(qs []State) State {
	qss := make([]string, len(qs))
	for i := 0; i < len(qs); i++ {
		qss[i] = string(qs[i])
	}
	sort.Strings(qss)
	return State(fmt.Sprintf("{%s}", strings.Join(uniq(qss), ",")))
}

// Compile this NFA to an executable DFA.
func (dfa *NFA) Compile() *DFA {
	return nil
}
