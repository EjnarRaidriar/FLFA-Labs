# Topic: Lexer & Scanner Implementation
### Course: Formal Languages & Finite Automata
### Author: Cucoș Emanuil

----

## Theory

### Lexer

A lexer (short for lexical analyzer) is a fundamental component in language processing systems that performs the first stage of the compilation or interpretation process.
It transforms a sequence of characters from the source code into a sequence of tokens, which are meaningful units of the language.

The lexer works by:
1. Reading the input character by character
2. Identifying sequences of characters that form valid tokens according to the language rules
3. Categorizing these tokens into specific types (keywords, identifiers, operators, etc.)
4. Passing these tokens to the next stage of processing (typically a parser)

### Tokens and Lexemes

**Lexeme**: The actual sequence of characters that matches a pattern for a token type

**Token**: A categorized unit of the language with metadata (type, value, position)

For example, in the statement `x = 5`:
- `x` is a lexeme that produces an identifier token
- `=` is a lexeme that produces an assignment operator token
- `5` is a lexeme that produces a number token

## Implementation

The lexer is done for a language called <b>Choice Script</b>.
This language is used for creating text interactive fictions.
It includes various commands for controlling game flow, variables, conditionals, and text content.

### Overview of the Lexer Structure

The lexer implementation consists of:

1. A set of token types defined as constants
2. A `Token` struct to represent individual tokens
3. A `Lexer` struct to handle the scanning process
4. Methods for reading and analyzing the input

### Token Types

The lexer recognizes various token types specific to the interactive fiction language:

```go
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
```

### Token Structure

Each token contains:

```go
type Token struct {
    Type   TokenType // The type of token
    Value  string    // The actual text value of the token
    Line   int       // Line number where the token appears
    Column int       // Column number where the token starts
}
```

### Lexer Structure

The lexer maintains state as it processes the input:

```go
type Lexer struct {
    reader  *bufio.Reader // Input reader
    line    int          // Current line
    column  int          // Current column
    current rune         // Current character
    eof     bool         // End of file flag
}
```

### Functions

The lexer implements several functions:

#### Character Reading

```go
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
```

This function reads the next character from the input, updates the lexer's state, and handles newlines.

#### Peek Ahead

```go
func (l *Lexer) peekChar() (rune, error) {
    ch, err := l.reader.Peek(1)
    return rune(ch[0]), err
}
```

This function allows the lexer to look ahead without consuming the character, which is useful for multi-character tokens.

#### Reading Identifiers

```go
func (l *Lexer) readIdentifier() string {
    var identifier strings.Builder

    for !l.eof && (unicode.IsLetter(l.current) || unicode.IsDigit(l.current) || l.current == '_') {
        identifier.WriteRune(l.current)
        l.readChar()
    }
    return identifier.String()
}
```

This function reads a sequence of characters that form a valid identifier.

#### Next Token

The main function of the lexer is `Next()`, which determines the next token based on the current character:

```go
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
        // Command processing
    case '#':
        // Choice processing
    case '/':
        // Comment processing
    case '$':
        // Variable processing
    default:
        // Text processing
    }

    return token
}
```

This function identifies token types based on initial characters and delegates to specialized functions for further processing.

### Implementation Details

#### Command Processing

The lexer identifies commands by the `*` prefix and categorizes them based on the command name:

```go
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
    // Additional command types...
    }

    params := l.readParameters()
    if params != "" {
        token.Value = fmt.Sprintf("%s %s", token.Value, params)
    }
```

This approach allows for specialized handling of different command types while preserving the original command name.

#### Variable Processing

The lexer recognizes variables in the format `${variable}`:

```go
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
```

#### Comment Processing

Comments are identified by the `//` prefix:

```go
case '/':
    if nextChar, err := l.peekChar(); err == nil && nextChar == '/' {
        l.readChar()
        l.readChar()
        token.Type = COMMENT
        token.Value = l.readToEndOfLine()
    } else {
        token.Type = TEXT
        token.Value = string(l.current)
        l.readChar()
    }
```

#### Text Processing

Any content that doesn't match specific token patterns is treated as text:

```go
default:
    token.Type = TEXT
    token.Value = l.readToEndOfLine()
```

### Utility Functions

The lexer includes several utility functions:

- `skipWhitespace()`: Skips over whitespace characters
- `readParameters()`: Reads command parameters
- `readToEndOfLine()`: Reads characters until the end of the current line
- `readUntil(delimiter)`: Reads characters until a specific delimiter is encountered

### Token String Representation

For debugging and display purposes, the lexer includes a function to convert token types to strings:

```go
func TokenToString(tokenType TokenType) string {
    switch tokenType {
    case EOF:
        return "EOF"
    case TEXT:
        return "TEXT"
    // Additional token types...
    default:
        return "UNKNOWN"
    }
}
```

## Conclusion

The implementation of this lexer demonstrates the fundamental principles of lexical analysis. It efficiently breaks down the input source code into meaningful tokens while preserving important contextual information.
The lexer's ability to recognize various language constructs enables subsequent stages of processing (parsing, interpretation) to work with structured data rather than raw text. This separation of concerns is a key aspect of compiler and interpreter design.
The modular approach taken in this implementation allows for easy extension to support additional token types or language features. The position tracking also provides valuable information for error reporting, enhancing the developer experience when using this language.
