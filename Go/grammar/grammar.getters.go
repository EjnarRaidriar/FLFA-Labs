package grammar

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
