package regex

type TokenType int

const (
	EOF TokenType = iota
	CHAR
	INT
	OR
	AND
	RP
	LP
	STAR
	PLUS
	POW
	OPTIONAL
)

type Token struct {
	Type  TokenType
	Value rune
}

func (t *Token) String() string {
	if t.Type == CHAR || t.Type == INT {
		return string(t.Value)
	} else {
		return t.Type.String()
	}
}

func (t TokenType) String() string {
	switch t {
	case EOF:
		return "EOF"
	case CHAR:
		return "CHAR"
	case INT:
		return "INT"
	case OR:
		return "OR"
	case AND:
		return "AND"
	case RP:
		return "RP"
	case LP:
		return "LP"
	case STAR:
		return "STAR"
	case PLUS:
		return "PLUS"
	case POW:
		return "POW"
	case OPTIONAL:
		return "OPTIONAL"
	}
	return "ERR"
}
