package automaton

type Transition struct {
	InitialState string
	Transition   rune
	NextState    string
}

type FiniteAutomaton struct {
	states       []string
	alphabet     []rune
	transitions  []Transition
	initialState string
	finalStates  []string
}

func NewFiniteAutomaton(
	states []string,
	alpahbet []rune,
	transitions []Transition,
	initialState string,
	finalStates []string) *FiniteAutomaton {

	return &FiniteAutomaton{
		states,
		alpahbet,
		transitions,
		initialState,
		finalStates}
}

func (f *FiniteAutomaton) CheckWord(word string) bool {
	var states map[string]bool = map[string]bool{f.initialState: true}
	for _, letter := range word {
		var nextStates map[string]bool = make(map[string]bool)
		for state := range states {
			for _, transition := range f.transitions {
				if state == transition.InitialState && letter == transition.Transition {
					nextStates[transition.NextState] = true
				}
			}
		}
		states = nextStates
	}
	for _, finalState := range f.finalStates {
		if states[finalState] {
			return true
		}
	}
	return false
}

func (f *FiniteAutomaton) DetermineFA() string {
	transitions := make(map[string]rune)
	isNFA := false
	for _, transition := range f.transitions {
		if transition.Transition == '&' {
			return "&-NFA"
		}
		if initialState, ok := transitions[transition.InitialState]; ok &&
			initialState == transition.Transition {
			isNFA = true
		} else {
			transitions[transition.InitialState] = transition.Transition
		}
	}
	if isNFA {
		return "NFA"
	}
	return "DFA"
}
