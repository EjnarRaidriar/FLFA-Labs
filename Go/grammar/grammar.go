package grammar

import (
	"fmt"
	"math/rand"
	"strings"
)

type Grammar struct {
	nonTerminals  []rune
	terminals     []rune
	productions   map[string][]string
	initialSymbol rune
}

func (g Grammar) GetProductions() map[string][]string {
	return g.productions
}

func (g Grammar) GetTerminals() []rune {
	return g.terminals
}

func (g Grammar) GetNonTerminals() []rune {
	return g.nonTerminals
}

func (g Grammar) GetInitialSymbol() rune {
	return g.initialSymbol
}

func NewGrammar(
	nonTerminals []rune,
	terminals []rune,
	productions map[string][]string,
	initialSymbol rune) *Grammar {
	return &Grammar{nonTerminals, terminals, productions, initialSymbol}
}

func hasNonTerminal(word string, nonTerminals []rune) bool {
	for _, nonnonTerminal := range nonTerminals {
		if strings.ContainsRune(word, nonnonTerminal) {
			return true
		}
	}
	return false
}

func (g Grammar) GenerateString() string {
	var word string = "S"
	for {
		if hasNonTerminal(word, g.nonTerminals) == false {
			break
		}
		for nonTerminal, productions := range g.productions {
			if strings.Contains(word, nonTerminal) {
				var transition string = productions[rand.Intn(len(productions))]
				word = strings.Replace(word, nonTerminal, transition, 1)
			}
		}
	}
	return word
}

// func (g Grammar) DefineGrammar() string {
// 	isType3 := false
// 	isRightRegular := false
// 	isType2 := false
// 	for nonTerminal := range g.productions {

// 	}
// 	return ""
// }

func (g Grammar) Print() {
	fmt.Print("Non-Terminals: [")
	for i, n := range g.nonTerminals {
		if i+1 == len(g.nonTerminals) {
			fmt.Printf("%c", n)
			break
		}
		fmt.Printf("%c ", n)
	}
	fmt.Print("]\n")
	fmt.Print("Terminals: [")
	for i, t := range g.terminals {
		if i+1 == len(g.terminals) {
			fmt.Printf("%c", t)
			break
		}
		fmt.Printf("%c ", t)
	}
	fmt.Print("]\n")
	fmt.Print("Productions:\n")
	for nonTerminal, productions := range g.productions {
		fmt.Printf("%s: %v\n", nonTerminal, productions)
	}
	fmt.Printf("Initials Symbol: %c\n", g.initialSymbol)
}
