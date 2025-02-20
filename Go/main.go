package main

import (
	"fmt"

	"flfa/automaton"
	"flfa/conversion"
	"flfa/grammar"
)

func main() {
	// lab1()
	lab2()
}

func lab2() {
	g := grammar.NewGrammar(
		[]rune{'S', 'B', 'C', 'D'},
		[]rune{'a', 'b', 'c'},
		map[string][]string{
			"S": {"aB"},
			"B": {"bS", "aC", "c"},
			"C": {"bD"},
			"D": {"c", "aC"},
		},
		'S',
	)
	fmt.Println(g.DefineGrammar())

	fa := automaton.NewFiniteAutomaton(
		[]string{"q0", "q1", "q2"},
		[]rune{'a', 'b', 'c'},
		[]automaton.Transition{
			{InitialState: "q0", Transition: 'a', NextState: "q1"},
			{InitialState: "q1", Transition: 'c', NextState: "q1"},
			{InitialState: "q1", Transition: 'b', NextState: "q2"},
			{InitialState: "q2", Transition: 'a', NextState: "q2"},
			{InitialState: "q0", Transition: 'a', NextState: "q0"},
			{InitialState: "q1", Transition: 'a', NextState: "q0"},
		},
		"q0",
		[]string{"q2"},
	)
	g = conversion.FaToRg(fa)
	g.Println()
}

func lab1() {
	g := grammar.NewGrammar(
		[]rune{'S', 'B', 'C', 'D'},
		[]rune{'a', 'b', 'c'},
		map[string][]string{
			"S": {"aB"},
			"B": {"bS", "aC", "c"},
			"C": {"bD"},
			"D": {"c", "aC"},
		},
		'S',
	)
	g.Println()
	for i := 0; i < 5; i++ {
		fmt.Printf("%s ", g.GenerateString())
	}
	fmt.Println()
	fa := conversion.RgToFa(g)
	fa.Println()
	var input string
	fmt.Print("Choose a word to check: ")
	fmt.Scanf("%s", &input)
	fmt.Printf("%v\n", fa.CheckWord(input))
}
