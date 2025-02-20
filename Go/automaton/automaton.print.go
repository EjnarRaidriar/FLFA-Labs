package automaton

import (
	"flfa/functions"
	"fmt"
)

func (p Transition) print() {
	fmt.Printf("(%s, %c) -> %s", p.InitialState, p.Transition, p.NextState)
}

func (f FiniteAutomaton) printTransitions() {
	fmt.Println("Transitions:")
	for _, n := range f.transitions {
		n.print()
		fmt.Println()
	}
}

func (f FiniteAutomaton) Println() {
	fmt.Printf("States: %v\n", f.states)
	fmt.Print("Alphabet: ")
	functions.PrintRuneList(f.alphabet)
	fmt.Println()
	f.printTransitions()
	fmt.Printf("Initial state: %s\n", f.initialState)
	fmt.Printf("Final states: %v\n", f.finalStates)
}
