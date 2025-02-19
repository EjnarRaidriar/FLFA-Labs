package main

import (
	"fmt"

	"flfa/conversion"
	"flfa/grammar"
)

func main() {
	lab1()
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
