package grammar

import "fmt"

func (g Grammar) Println() {
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
