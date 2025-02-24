package automaton

import (
	"flfa/functions"
	"slices"
	"strings"
)

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

func (f *FiniteAutomaton) closure(state string) map[string]bool {
	closure := map[string]bool{state: true}
	for element := range closure {
		for _, transition := range f.transitions {
			if transition.InitialState == element &&
				transition.Transition == '&' {
				closure[transition.NextState] = true
			}
		}
	}
	return closure
}

func joinStates(states []string) string {
	if len(states) > 1 {
		return "{" + strings.Join(states, ", ") + "}"
	}
	return states[0]
}

func MakeDFA(f FiniteAutomaton) *FiniteAutomaton {
	if f.DetermineFA() == "DFA" {
		return &f
	}
	dfaTransitions := make([]Transition, 0, 10)
	dfaFinalStates := make([]string, 0, 2)
	dfaStatesMap := make([]map[string]bool, 0)

	closures := make(map[string]map[string]bool)
	for _, state := range f.states {
		closures[state] = f.closure(state)
		var newState map[string]bool = closures[state]
		dfaStatesMap = append(dfaStatesMap, newState)
	}

	i := 0
	for i < len(dfaStatesMap) {
		state := dfaStatesMap[i]
		for _, symbol := range f.alphabet {
			newStateMap := make(map[string]bool)
			for element := range state {
				for _, transition := range f.transitions {
					if transition.Transition == symbol &&
						transition.InitialState == element {
						for closure, ok := range closures[transition.NextState] {
							newStateMap[closure] = ok
						}
					}
				}
			}
			if len(newStateMap) > 0 {
				stateSlice := functions.KeyList(state)
				newStateSlice := functions.KeyList(newStateMap)
				dfaTransitions = append(dfaTransitions, Transition{
					joinStates(stateSlice),
					symbol,
					joinStates(newStateSlice),
				})
				if functions.ContainsMap(dfaStatesMap, newStateMap) == false {
					dfaStatesMap = append(dfaStatesMap, newStateMap)
				}
			}
		}
		i++
	}
	dfaStates := make([]string, 0, len(dfaStatesMap))
	for _, stateMap := range dfaStatesMap {
		state := functions.KeyList(stateMap)
		dfaStates = append(dfaStates, joinStates(state))
		for _, finalState := range f.finalStates {
			if slices.Contains(state, finalState) {
				dfaFinalStates = append(dfaFinalStates, joinStates(state))
			}
		}
	}
	return NewFiniteAutomaton(dfaStates, f.alphabet, dfaTransitions, f.initialState, dfaFinalStates)
}
