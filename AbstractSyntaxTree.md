# Topic: Parser & Building an Abstract Syntax Tree
### Course: Formal Languages & Finite Autopata
### Author: Cucoș Emanuil

----

## Theory

### Parser

**Parser** is a tool for formal analysis of a sentence or other string of words.
A parse tree contains semantic information that could be used in subsequent stages of compilation.
The structure of a parse tree can be represented with an **Abstract Syntax Tree** (AST).
The tree is generated from lexer's tokens using rules that belong to a language.
The purpose of the tree is to specify the order of operations that form a word or sequence of words.

### Pratt Parsing Algorithm

**Prat Parsing Algorithm** is a top-down-operator-precedence parsing discovered by Vaughan Pratt. This algorithm uses precedences, also called binding powers, to build the AST. Binding powers indicate which operation should be executed first.

## Implementation

### Tokens

```go
type Token struct {
	Type  TokenType
	Value rune
}
```

```go
const (
	EOF      TokenType = iota
	CHAR               // any alphabetic character
	INT                // any digit
	OR                 // "|"
	AND                // concatination of char, int or parentheses
	RP                 // ")"
	LP                 // "("
	STAR               // "*"
	PLUS               // "+"
	POW                // "^"
	OPTIONAL           // "?"
)
```

### Lexer

Lexer contains a list of tokens and two ways to access them.

```go
type Lexer struct {
	Tokens []Token
}
```

First way to access a token is to take it from the lexer.
A taken token means that it was already processed.

```go
func (l *Lexer) Next() Token
```

Second way to access a token is to peek ahead without taking it from the lexer.

```go
func (l *Lexer) Peek() Token
```

### Binding Powers

Bindings powers are one of the most important feature of the Pratt Parsing.
They specify two which operator an atom coresponds to.

There is infix and postfix bindig powers.
Infix binding power is for operators that require atoms on both sides.
Postfix binding power is for operatos that should have an atom only on the left side.

```go
func getInfixBindingPower(tokenType TokenType) BindingPower {
	bindingPower, ok := map[TokenType]BindingPower{
		OR:  {1.0, 1.1},
		AND: {2.0, 2.1},
		POW: {3.1, 3.0},
	}[tokenType]
	if !ok {
		fmt.Printf("Error in binding lookup: %+v\n", tokenType)
		return BindingPower{}
	}
	return bindingPower
}

func getPostfixBindingPower(tokenType TokenType) (BindingPower, error) {
	bindingPower, ok := map[TokenType]BindingPower{
		STAR:     {4, 0},
		PLUS:     {4, 0},
		OPTIONAL: {4, 0},
	}[tokenType]
	if !ok {
		return BindingPower{}, errors.New("wrong token")
	}
	return bindingPower, nil
}
```

### Parser

The parser is the most important part of the code.
It takes the input string, forms a lexer and then recursively forms the AST.

This is the structure of the AST:

```go
type ParseTreeNode struct {
	Token
	LeftSide  *ParseTreeNode
	RightSide *ParseTreeNode
}
```

This is the pseudocode of the parser:  

```go
func parseTree(l *Lexer, minBP float32) (ParseTreeNode, error) {
  // Consume token from lexer as left hand sign (lhs);

  // If lhs is left parenthesis
  // then parse tree deeper and assign it to this lhs;
  // |_ if next token is not a right parenthesis
  // |_ then there is an error;

  // If lhs is not int or char, then error;

  // Infinite loop;
  // |_ Peek into lexer and assign it to operand;
  // |_ If operand token is EOF or right parenthesis, then exit loop;
  // |_ If operand is a CHAR, then error;
  // |_ Get postfix binding power (pbp);
  // |_ If pbp doesn't have error (it is quantization operand);
  // |- |_ if minPB > left pbp, then exit infinite loop;
  // |- |_ consume lexer token;
  // |- |_ copy lhs and
  // |- |_ assign it to a new lhs where
  // |- |_ head is the operand and left part is the copy;
  // |- |_ start the loop again
  // |_ If minPB > infix binding power, then exit loop;
  // |_ Consume lexer token;
  // |_ Parse tree deeper and assign it to right hand sign (rhs);
  // |_ If operand is power and rhs is not INT, then error;
  // |_ Copy lhs;
  // |_ assign to new lhs:
  // |_ head - operator; left side - copy; right side - rhs; (end of loop)

  // return lhs
}
```

Note: quantifiers have a limit of 5 occurences to limit the output word.

## Conclusion

This work implements a string generation from a regex using Pratt parsing algorithm.
The lexer tokenizes the input and is used by the parser to generate AST.
The parser is the most complex module which includes both language rules and handles errors.
The algorithm uses binding powers to handle operator precedence ensuring no ambiguity and deterministic output.
Quantifiers are limited to maximum of 5 occurences.
