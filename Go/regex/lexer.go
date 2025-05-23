package regex

import (
	"fmt"
	"strings"
	"unicode"
)

type Lexer struct {
	Tokens []Token
}

func NewLexer(input string) Lexer {
	lexer := Lexer{Tokens: make([]Token, 0, len(input))}

	for _, ch := range input {
		if unicode.IsLetter(ch) {
			// AND when: "AB", ")B", "A*B", "A+B", "A?B"
			addAndIfNeeded(&lexer)
			lexer.Tokens = append(lexer.Tokens, Token{Type: CHAR, Value: ch})
			continue
		}

		if unicode.IsNumber(ch) {
			addAndIfNeeded(&lexer)
			lexer.Tokens = append(lexer.Tokens, Token{Type: INT, Value: ch})
			continue
		}

		tokenType, ok := map[rune]TokenType{
			'|': OR,
			'(': LP,
			')': RP,
			'*': STAR,
			'+': PLUS,
			'?': OPTIONAL,
			'^': POW,
		}[ch]

		if !ok {
			fmt.Printf("Error, found %c", ch)
			return Lexer{}
		}

		if tokenType == LP {
			// AND when: "A(", ")(", "*(", "+(", "?("
			addAndIfNeeded(&lexer)
		}

		lexer.Tokens = append(lexer.Tokens, Token{Type: tokenType, Value: ch})
	}
	return lexer
}

func addAndIfNeeded(l *Lexer) {
	if len(l.Tokens) > 0 {
		if l.Tokens[len(l.Tokens)-1].Type == CHAR ||
			l.Tokens[len(l.Tokens)-1].Type == INT ||
			l.Tokens[len(l.Tokens)-1].Type == RP ||
			l.Tokens[len(l.Tokens)-1].Type == PLUS ||
			l.Tokens[len(l.Tokens)-1].Type == OPTIONAL ||
			l.Tokens[len(l.Tokens)-1].Type == STAR {
			l.Tokens = append(l.Tokens, Token{Type: AND, Value: '_'})
		}
	}
}

func (l *Lexer) Next() Token {
	if len(l.Tokens) > 0 {
		nextToken := l.Tokens[0]
		l.Tokens = l.Tokens[1:]
		return nextToken
	}
	return Token{Type: EOF, Value: ' '}
}

func (l *Lexer) Peek() Token {
	if len(l.Tokens) > 0 {
		nextToken := l.Tokens[0]
		return nextToken
	}
	return Token{Type: EOF, Value: ' '}
}

func (l *Lexer) String() string {
	var result strings.Builder
	for _, t := range l.Tokens {
		result.WriteString(fmt.Sprintf("%s ", t.String()))
	}
	return result.String()
}
