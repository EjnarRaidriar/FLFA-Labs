package main

import (
	// "fmt"

	"flfa/grammar"

	// "flfa/automaton"
	"flfa/conversion"
)

func main() {
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
	g.Print()
	fa := conversion.RgToFa(g)
	fa.Print()
}
