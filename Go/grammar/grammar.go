package grammar

import (
	"flfa/functions"
	"math/rand"
	"strings"
	"unicode"
	"unicode/utf8"
)

type GrammarType string

const (
	RightRegular GrammarType = "Right Regular"
	LeftRegular  GrammarType = "Left Regular"
	Type_2       GrammarType = "Type 2"
	Type_1       GrammarType = "Type 1"
	Type_0       GrammarType = "Type 0"
)

type Grammar struct {
	nonTerminals  []rune
	terminals     []rune
	productions   map[string][]string
	initialSymbol rune
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

func (g Grammar) DefineGrammar() GrammarType {
	isType3 := true
	isRightLiniar := false
	isType2 := true
	for nonTerminal, productions := range g.productions {
		if utf8.RuneCountInString(nonTerminal) > 1 {
			isType3 = false
			isType2 = false
		}
		for _, prod := range productions {
			prodLen := utf8.RuneCountInString(prod)
			if utf8.RuneCountInString(nonTerminal) > prodLen {
				return Type_0
			}
			if (isType2 || isType3) == false &&
				nonTerminal != string(g.initialSymbol) &&
				prod != "&" {
				return Type_0
			}
			lastLetter := rune(prod[prodLen-1])
			if isType2 && functions.IsLower(prod) == false {
				if functions.IsUpper(prod) {
					isType3 = false
				} else if functions.HasMultipleUpper(prod) {
					isType3 = false
				}
			}
			if isType3 && prodLen > 1 {
				if unicode.IsUpper(lastLetter) {
					isRightLiniar = true
				}
				if isRightLiniar && unicode.IsLower(lastLetter) {
					isType3 = false
				}
			}
		}
	}
	if isType3 && isRightLiniar {
		return RightRegular
	} else if isType3 {
		return LeftRegular
	}
	if isType2 {
		return Type_2
	}
	return Type_1
}
