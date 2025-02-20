package automaton

func (f FiniteAutomaton) GetInitialState() string {
	return f.initialState
}

func (f FiniteAutomaton) GetAlphabet() []rune {
	return f.alphabet
}

func (f FiniteAutomaton) GetStates() []string {
	return f.states
}

func (f FiniteAutomaton) GetTransitions() []Transition {
	return f.transitions
}

func (f FiniteAutomaton) GetFinalStates() []string {
	return f.finalStates
}
