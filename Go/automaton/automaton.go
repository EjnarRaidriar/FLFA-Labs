package automaton

import (
	"fmt"

	"flfa/functions"
)

type Transition struct {
	InitialState string
	Transition   string
	NextState    string
}

type FiniteAutomaton struct {
	states       []string
	alphabet     []rune
	transitions  []Transition
	initialState rune
	finalStates  []string
}

func NewFiniteAutomaton(
	states []string,
	alpahbet []rune,
	transitions []Transition,
	initialState rune,
	finalStates []string) *FiniteAutomaton {

	return &FiniteAutomaton{
		states,
		alpahbet,
		transitions,
		initialState,
		finalStates}
}

func (f FiniteAutomaton) GetInitialState() rune {
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

func (p Transition) print() {
	fmt.Printf("(%s, %s) -> %s", p.InitialState, p.Transition, p.NextState)
}

func (f FiniteAutomaton) printTransitions() {
	fmt.Println("Transitions:")
	for _, n := range f.transitions {
		n.print()
		fmt.Println()
	}
}

func (f FiniteAutomaton) Print() {
	fmt.Printf("States: %v %v\n", f.states, len(f.states))
	fmt.Print("Alphabet: ")
	functions.PrintRuneList(f.alphabet)
	fmt.Println()
	f.printTransitions()
	fmt.Printf("Initial state: %c\n", f.initialState)
	fmt.Printf("Final states: %v\n", f.finalStates)
}
