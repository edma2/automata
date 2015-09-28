// Nondeterministic finite automatons (with ε-transitions)
package automata

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
	dfa := NFA{q0: q0, delta: make(TransitionFunc)}
	dfa.final = make(map[State]bool)
	for i := 0; i < len(final); i++ {
		dfa.final[final[i]] = true
	}
	return &dfa
}

func (dfa *NFA) NewTransition(old State, input Symbol, new []State) {
	transition := dfa.delta[old]
	if transition == nil {
		transition = make(Transition)
		dfa.delta[old] = transition
	}
	transition[input] = new
}

// Compile this NFA to an executable DFA.
func (dfa *NFA) Compile() *DFA {
	return nil
}
