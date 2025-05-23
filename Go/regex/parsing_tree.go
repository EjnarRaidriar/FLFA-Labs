package regex

import (
	"errors"
	"fmt"
	"strings"
)

type ParseTreeNode struct {
	Token
	LeftSide  *ParseTreeNode
	RightSide *ParseTreeNode
}

func NewParseTree(input string) (ParseTreeNode, error) {
	lexer := NewLexer(input)
	return parseTree(&lexer, 0)
}

func parseTree(l *Lexer, minBP float32) (ParseTreeNode, error) {
	var err error
	lhs := ParseTreeNode{Token: l.Next()}
	if lhs.Type == LP {
		lhs, err = parseTree(l, 0.0)
		if err != nil {
			return ParseTreeNode{}, err
		}
		if l.Next().Type != RP {
			return ParseTreeNode{}, errors.New("expected parenthesis closing")
		}
	}
	if lhs.Type != CHAR && lhs.Type != INT && lhs.RightSide == nil {
		return ParseTreeNode{}, errors.New("new left leaf is not CHAR")
	}
	for {
		op := ParseTreeNode{Token: l.Peek()}
		if op.Type == EOF || op.Type == RP {
			break
		}
		if op.Type == CHAR {
			return ParseTreeNode{}, errors.New("expected operand")
		}
		bindingPower, err := getPostfixBindingPower(op.Type)
		if err == nil {
			if minBP > bindingPower.Left {
				break
			}
			l.Next()
			lhsCopy := lhs
			lhs = ParseTreeNode{Token: op.Token, LeftSide: &lhsCopy}
			continue
		}
		bindingPower = getInfixBindingPower(op.Type)
		if minBP > bindingPower.Left {
			break
		}
		l.Next()
		rhs, err := parseTree(l, bindingPower.Right)
		if err != nil {
			return ParseTreeNode{}, err
		}
		if op.Type == POW && rhs.Type != INT {
			return ParseTreeNode{}, errors.New("expected INT after POW")
		}
		lhsCopy := lhs
		lhs = ParseTreeNode{Token: op.Token, LeftSide: &lhsCopy, RightSide: &rhs}
	}
	return lhs, err
}

func (t *ParseTreeNode) String() string {
	if t == nil {
		return ""
	}

	var result strings.Builder
	var preorderLisp func(node *ParseTreeNode)
	preorderLisp = func(node *ParseTreeNode) {
		if node == nil {
			return
		}

		if node.Type == CHAR || node.Type == INT {
			result.WriteString(fmt.Sprintf("%c", node.Value))
		} else {
			result.WriteString(fmt.Sprintf("(%c ", node.Value))
			preorderLisp(node.LeftSide)
			if node.RightSide != nil {
				result.WriteString(" ")
				preorderLisp(node.RightSide)
			}
			result.WriteString(")")
		}
	}

	preorderLisp(t)

	return result.String()
}
