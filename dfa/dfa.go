package dfa

// A deterministic finite automaton M is a 5-tuple, (Q, Σ, δ, q0, F), consisting of

// a finite set of states (Q)
type state string

// a finite set of input symbols called the alphabet (Σ)
type symbol rune

// a transition function (δ : Q × Σ → Q)
type transitionTable map[symbol]state
type transitionFunc map[state]transitionTable

// a start state (q0 ∈ Q)
// a set of accept states (F ⊆ Q)
type DFA struct {
	delta  transitionFunc
	start  state
	accept map[state]bool
}

func NewDFA(startState string, acceptStates []string) *DFA {
	dfa := new(DFA)
	dfa.delta = make(transitionFunc)
	dfa.start = state(startState)
	dfa.accept = make(map[state]bool)
	for i := 0; i < len(acceptStates); i++ {
		dfa.accept[state(acceptStates[i])] = true
	}
	return dfa
}

func (dfa *DFA) NewTransition(old string, input rune, new string) {
	transition := dfa.delta[state(old)]
	if transition == nil {
		transition = make(map[symbol]state)
		dfa.delta[state(old)] = transition
	}
	transition[symbol(input)] = state(new)
}

func (dfa *DFA) Execute(input string) bool {
	state := dfa.start
	for _, runeValue := range input {
		state = dfa.delta[state][symbol(runeValue)]
		// empty transition
		if state == "" {
			return false
		}
	}
	return dfa.accept[state]
}
