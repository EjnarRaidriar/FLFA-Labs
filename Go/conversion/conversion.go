package conversion

import (
	"flfa/automaton"
	"flfa/functions"
	"flfa/grammar"
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
		g.GetInitialSymbol(),
		finalStates)
}

func FaToRg(f *automaton.FiniteAutomaton) *grammar.Grammar {
	// todo convert all transitions initial state string into single character
	var nonTerminals []rune
	productions := make(map[string][]string)
	for _, transition := range f.GetTransitions() {
		var production string
		if transition.NextState == "X" {
			production = string(transition.Transition)
		} else {
			production = strings.Join([]string{string(transition.Transition), string(transition.NextState)}, "")
		}
		if nonTerminalsList, ok := productions[transition.InitialState]; ok {
			productions[transition.InitialState] = append(nonTerminalsList, production)
		} else {
			productions[transition.InitialState] = []string{production}
			// InitialState[0] is a crutch, do the todo first
			nonTerminals = append(nonTerminals, rune(transition.InitialState[0]))
		}
	}
	return grammar.NewGrammar(nonTerminals, f.GetAlphabet(), productions, f.GetInitialState())
}
