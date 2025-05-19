package regex_lexer

import (
	"errors"
	"fmt"
)

type BindingPower struct {
	Left  float32
	Right float32
}

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
