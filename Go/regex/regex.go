package regex

import (
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
	"time"
)

func GenerateString(regEx string) (string, error) {
	tree, err := NewParseTree(regEx)
	if err != nil {
		return "", err
	}
	return generateStringFromNode(&tree), nil
}

func generateStringFromNode(n *ParseTreeNode) string {
	var result strings.Builder
	source := rand.NewPCG(uint64(time.Now().Nanosecond()), uint64(time.Now().Nanosecond()))
	switch n.Type {
	case INT, CHAR:
		return string(n.Value)
	case OR:
		return or(source, generateStringFromNode(n.LeftSide), generateStringFromNode(n.RightSide))
	case AND:
		return and(generateStringFromNode(n.LeftSide), generateStringFromNode(n.RightSide))
	case POW:
		return pow(generateStringFromNode(n.LeftSide), generateStringFromNode(n.RightSide))
	case STAR:
		return star(source, generateStringFromNode(n.LeftSide))
	case PLUS:
		return plus(source, generateStringFromNode(n.LeftSide))
	case OPTIONAL:
		return optional(source, generateStringFromNode(n.LeftSide))
	}

	return result.String()
}

func or(source *rand.PCG, str1 string, str2 string) string {
	r := rand.New(source).Float64()
	if r < 0.5 {
		return str1
	} else {
		return str2
	}
}

func and(str1 string, str2 string) string {
	return str1 + str2
}

func pow(str1 string, str2 string) string {
	times, err := strconv.Atoi(str2)
	if err != nil {
		fmt.Println("Conversion error: ", err)
		return ""
	}
	var result strings.Builder
	for range times {
		result.WriteString(str1)
	}
	return result.String()
}

func star(source *rand.PCG, str string) string {
	var result strings.Builder
	r := rand.New(source).IntN(6)
	for range r {
		result.WriteString(str)
	}
	return result.String()
}

func plus(source *rand.PCG, str string) string {
	var result strings.Builder
	r := rand.New(source).IntN(5) + 1
	for range r {
		result.WriteString(str)
	}
	return result.String()
}

func optional(source *rand.PCG, str string) string {
	r := rand.New(source).Float64()
	if r < 0.5 {
		return str
	} else {
		return ""
	}
}
