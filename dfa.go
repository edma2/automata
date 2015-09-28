package automata

// a transition function (δ : Q × Σ → Q)
type dfaTransition map[Symbol]State
type dfaTransitionFunc map[State]dfaTransition

type DFA struct {
	q0    State
	final map[State]bool
	delta dfaTransitionFunc
}

func (dfa *DFA) Execute(input string) bool {
	state := dfa.q0
	for _, runeValue := range input {
		state = dfa.delta[state][Symbol(runeValue)]
		if state == "" {
			return false
		}
	}
	return dfa.final[state]
}
