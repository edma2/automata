// Nondeterministic finite automatons (with ε-transitions)
package automata

import (
	"fmt"
	"sort"
	"strings"
)

// States are represented by Go runes.
// Each rune is an integer value identifying a Unicode code point.
type State rune

// A set of States
type StateSet struct {
	states map[State]bool
}

// A NFA-ε is represented formally by a 5-tuple, (Q, Σ, Δ, q0, F), consisting of
//
// a finite set of states Q
// a finite set of input symbols Σ
// a transition function Δ : Q × (Σ ∪ {ε}) → P(Q)
// an initial (or start) state q0 ∈ Q
// a set of states F distinguished as accepting (or final) states F ⊆ Q.
type NFA struct {
}

// Create a new StateSet - silently ignore duplicates
func NewStateSet(states []State) *StateSet {
	ss := new(StateSet)
	ss.states = make(map[State]bool)
	for _, s := range states {
		ss.states[s] = true
	}
	return ss
}

// Display a StateSet as a string (e.g. "{1,4,6}")
func (ss *StateSet) String() string {
	a := make([]string, len(ss.states))
	i := 0
	for s, _ := range ss.states {
		a[i] = string(s)
		i = i + 1
	}
	sort.Strings(a)
	return fmt.Sprintf("{%s}", strings.Join(a, ","))
}

// Combine two StateSets, returning a new StateSet
func (ss *StateSet) Concat(other *StateSet) *StateSet {
	states := make([]State, len(ss.states)+len(other.states))
	i := 0
	for s, _ := range ss.states {
		states[i] = s
		i = i + 1
	}
	for s, _ := range other.states {
		states[i] = s
		i = i + 1
	}
	return NewStateSet(states)
}
