package conversion

import (
	"flfa/automaton"
	"flfa/functions"
	"flfa/grammar"
	"slices"
	"strings"
)

func RgToFa(g *grammar.Grammar) *automaton.FiniteAutomaton {
	productions := g.GetProductions()
	states := functions.RuneSliceToStringSlice(append(g.GetNonTerminals(), 'X'))
	finalStates := []string{"X"}
	var transitions []automaton.Transition
	for nonTerminal, productions := range productions {
		for _, production := range productions {
			var newTranision automaton.Transition
			newTranision.InitialState = nonTerminal
			newTranision.Transition = rune(production[0])
			if len(production) == 1 {
				newTranision.NextState = "X"
			} else {
				newTranision.NextState = string(production[1])
			}
			transitions = append(transitions, newTranision)
		}
	}
	return automaton.NewFiniteAutomaton(
		states,
		g.GetTerminals(),
		transitions,
		string(g.GetInitialSymbol()),
		finalStates)
}

func FaToRg(f *automaton.FiniteAutomaton) *grammar.Grammar {
	// todo convert all transitions initial state string into single character
	conversionMap := statesToNonTerminals(f.GetStates(), f.GetInitialState())
	var nonTerminals []rune
	productions := make(map[string][]string)
	finalStates := f.GetFinalStates()
	finalNotFinal := false

	for _, transition := range f.GetTransitions() {
		if transition.InitialState == transition.NextState &&
			slices.Contains(finalStates, transition.NextState) {
			finalNotFinal = true
		}
	}

	for _, transition := range f.GetTransitions() {
		var production string
		nonTerminal := conversionMap[transition.InitialState]
		if slices.Contains(finalStates, transition.NextState) {
			if finalNotFinal {
				if transition.InitialState == transition.NextState {
					production = strings.Join([]string{string(transition.Transition), string(conversionMap[transition.NextState])}, "")
					addProduction(&productions, &nonTerminals, nonTerminal, production)
					production = string(transition.Transition)
					addProduction(&productions, &nonTerminals, nonTerminal, production)
				} else {
					production = string(transition.Transition)
					addProduction(&productions, &nonTerminals, nonTerminal, production)
				}
			}
		} else {
			production = strings.Join([]string{string(transition.Transition), string(conversionMap[transition.NextState])}, "")
			addProduction(&productions, &nonTerminals, nonTerminal, production)
		}
	}
	return grammar.NewGrammar(nonTerminals, f.GetAlphabet(), productions, conversionMap[f.GetInitialState()])
}

func addProduction(productions *map[string][]string, nonTerminals *[]rune, nonTerminal rune, production string) {
	newProduction := *productions
	if _, ok := newProduction[string(nonTerminal)]; ok {
		newProduction[string(nonTerminal)] = append(newProduction[string(nonTerminal)], production)
	} else {
		newProduction[string(nonTerminal)] = []string{production}
		*nonTerminals = append(*nonTerminals, nonTerminal)
	}
	*productions = newProduction
}

func statesToNonTerminals(states []string, initialState string) map[string]rune {
	var upperCaseAlphabet []rune = []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I',
		'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R',
		/* 'S',*/ 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	result := map[string]rune{initialState: 'S'}
	counter := 0
	for _, state := range states {
		if state != initialState {
			result[state] = upperCaseAlphabet[counter]
			counter++
		}
	}
	return result
}
