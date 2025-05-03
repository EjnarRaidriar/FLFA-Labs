package main

import (
	"fmt"
	"os"
	"strings"

	"flfa/automaton"
	"flfa/conversion"
	"flfa/grammar"
	"flfa/lexer"
)

func main() {
	// lab1()
	// lab2()
	lab3()
}

func lab3() {
	content, err := os.ReadFile("../choicescript.txt")
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	text := string(content)

	l := lexer.NewLexer(strings.NewReader(text))
	for {
		token := l.Next()
		fmt.Printf("Line %d, Col %d: %s - %q\n", token.Line, token.Column, lexer.TokenToString(token.Type), token.Value)
		if token.Type == lexer.EOF {
			break
		}
	}
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
			{InitialState: "q0", Transition: 'a', NextState: "q0"},
			{InitialState: "q0", Transition: 'a', NextState: "q1"},
			{InitialState: "q1", Transition: 'a', NextState: "q0"},
			{InitialState: "q1", Transition: 'b', NextState: "q2"},
			{InitialState: "q1", Transition: 'c', NextState: "q1"},
			{InitialState: "q2", Transition: 'a', NextState: "q2"},
			{InitialState: "q2", Transition: '&', NextState: "q2"},
		},
		"q0",
		[]string{"q2"},
	)
	fmt.Println(fa.DetermineFA())
	g = conversion.FaToRg(fa)
	g.Println()
	fa = automaton.MakeDFA(*fa)
	fa.Println()
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
