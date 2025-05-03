package lexer

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"unicode"
)

type TokenType int

const (
	EOF TokenType = iota
	TEXT
	COMMAND         // *command
	LABEL           // *label
	CHOICE          // #choice
	OPTION          // *choice
	VARIABLE        // ${variable}
	COMMENT         // comment starting with //
	CONDITIONAL     // *if, *elseif, *else
	END_CONDITIONAL // *endif
	SET             // *set
	GOTO            // *goto
	GOSUB           // *gosub
	RETURN          // *return
	FINISH          // *finish
	ACHIEVE         // *achieve
	SCENE_LIST      // *scene_list
	CREATE          // *create
	TEMP            // *temp
	STAT_CHART      // *stat_chart
	IMAGE           // *image
	SOUND           // *sound
	INPUT_TEXT      // *input_text
	INPUT_NUMBER    // *input_number
	RAND            // *rand
	MULTIREPLACE    // *multireplace
	PARAMS          // Parameters for commands
)

type Token struct {
	Type   TokenType
	Value  string
	Line   int
	Column int
}

type Lexer struct {
	reader  *bufio.Reader
	line    int
	column  int
	current rune
	eof     bool
}

func NewLexer(reader io.Reader) *Lexer {
	lexer := &Lexer{
		reader: bufio.NewReader(reader),
		line:   1,
		column: 0,
	}
	lexer.readChar()
	return lexer
}

func (l *Lexer) readChar() {
	ch, _, err := l.reader.ReadRune()
	if err != nil {
		l.eof = true
		return
	}

	l.current = ch
	l.column++
	if ch == '\n' {
		l.line++
		l.column = 0
	}
}

func (l *Lexer) peekChar() (rune, error) {
	ch, err := l.reader.Peek(1)
	return rune(ch[0]), err
}

func (l *Lexer) skipWhitespace() {
	for unicode.IsSpace(l.current) && !l.eof {
		l.readChar()
	}
}

func (l *Lexer) readIdentifier() string {
	var identifier strings.Builder

	for !l.eof && (unicode.IsLetter(l.current) || unicode.IsDigit(l.current) || l.current == '_') {
		identifier.WriteRune(l.current)
		l.readChar()
	}
	return identifier.String()
}

func (l *Lexer) readParameters() string {
	return strings.TrimSpace(l.readToEndOfLine())
}

func (l *Lexer) readToEndOfLine() string {
	var text strings.Builder
	for !l.eof && l.current != '\n' {
		text.WriteRune(l.current)
		l.readChar()
	}
	if !l.eof {
		l.readChar()
	}
	return text.String()
}

func (l *Lexer) readUntil(delimiter rune) string {
	var text strings.Builder
	for !l.eof && l.current != delimiter && l.current != '\n' {
		text.WriteRune(l.current)
		l.readChar()
	}
	return text.String()
}

func (l *Lexer) Next() Token {
	if l.eof {
		return Token{Type: EOF, Line: l.line, Column: l.column}
	}

	l.skipWhitespace()

	var token Token
	token.Line = l.line
	token.Column = l.column

	switch l.current {
	case '*':
		l.readChar()
		cmdName := l.readIdentifier()
		token.Type = COMMAND
		token.Value = cmdName

		switch strings.ToLower(cmdName) {
		case "if", "elseif", "else":
			token.Type = CONDITIONAL
		case "endif":
			token.Type = END_CONDITIONAL
		case "set":
			token.Type = SET
		case "goto":
			token.Type = GOTO
		case "gosub":
			token.Type = GOSUB
		case "return":
			token.Type = RETURN
		case "finish":
			token.Type = FINISH
		case "achieve":
			token.Type = ACHIEVE
		case "scene_list":
			token.Type = SCENE_LIST
		case "create":
			token.Type = CREATE
		case "temp":
			token.Type = TEMP
		case "stat_chart":
			token.Type = STAT_CHART
		case "image":
			token.Type = IMAGE
		case "sound":
			token.Type = SOUND
		case "input_text":
			token.Type = INPUT_TEXT
		case "input_number":
			token.Type = INPUT_NUMBER
		case "rand":
			token.Type = RAND
		case "multireplace":
			token.Type = MULTIREPLACE
		case "label":
			token.Type = LABEL
		case "choice":
			token.Type = OPTION
		}

		params := l.readParameters()
		if params != "" {
			token.Value = fmt.Sprintf("%s %s", token.Value, params)
		}

	case '#':
		l.readChar()
		token.Type = CHOICE
		token.Value = l.readToEndOfLine()

	case '/':
		if nextChar, err := l.peekChar(); err == nil && nextChar == '/' {
			l.readChar()
			l.readChar()
			token.Type = COMMAND
			token.Value = l.readToEndOfLine()
		} else {
			token.Type = TEXT
			token.Value = string(l.current)
			l.readChar()
		}

	case '$':
		if nextChar, err := l.peekChar(); err == nil && nextChar == '{' {
			l.readChar()
			l.readChar()
			variable := l.readUntil('}')
			l.readChar()
			token.Type = VARIABLE
			token.Value = variable
		} else {
			token.Type = TEXT
			token.Value = string(l.current)
			l.readChar()
		}

	default:
		token.Type = TEXT
		token.Value = l.readToEndOfLine()
	}

	return token
}

func TokenToString(tokenType TokenType) string {
	switch tokenType {
	case EOF:
		return "EOF"
	case TEXT:
		return "TEXT"
	case COMMAND:
		return "COMMAND"
	case LABEL:
		return "LABEL"
	case CHOICE:
		return "CHOICE"
	case OPTION:
		return "OPTION"
	case VARIABLE:
		return "VARIABLE"
	case COMMENT:
		return "COMMENT"
	case CONDITIONAL:
		return "CONDITIONAL"
	case END_CONDITIONAL:
		return "END_CONDITIONAL"
	case SET:
		return "SET"
	case GOTO:
		return "GOTO"
	case GOSUB:
		return "GOSUB"
	case RETURN:
		return "RETURN"
	case FINISH:
		return "FINISH"
	case ACHIEVE:
		return "ACHIEVE"
	case SCENE_LIST:
		return "SCENE_LIST"
	case CREATE:
		return "CREATE"
	case TEMP:
		return "TEMP"
	case STAT_CHART:
		return "STAT_CHART"
	case IMAGE:
		return "IMAGE"
	case SOUND:
		return "SOUND"
	case INPUT_TEXT:
		return "INPUT_TEXT"
	case INPUT_NUMBER:
		return "INPUT_NUMBER"
	case RAND:
		return "RAND"
	case MULTIREPLACE:
		return "MULTIREPLACE"
	case PARAMS:
		return "PARAMS"
	default:
		return "UNKNOWN"
	}
}
