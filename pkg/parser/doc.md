# parser
--
    import "github.com/go-i2p/go-rst/pkg/parser"


## Usage

#### type Lexer

```go
type Lexer struct {
}
```

Lexer represents a lexer for the input text.

#### func  NewLexer

```go
func NewLexer() *Lexer
```
NewLexer creates a new Lexer instance.

#### func (*Lexer) Tokenize

```go
func (l *Lexer) Tokenize(line string) Token
```
Tokenize tokenizes a single line of input text.

#### type Parser

```go
type Parser struct {
}
```

Parser is a struct that holds the state of the parser.

#### func  NewParser

```go
func NewParser(trans translator.Translator) *Parser
```
NewParser creates a new Parser instance.

#### func (*Parser) Parse

```go
func (p *Parser) Parse(content string) []nodes.Node
```
Parse takes a string of reStructuredText content and returns a slice of Node
instances.

#### type ParserContext

```go
type ParserContext struct {
}
```

ParserContext represents the current state of the parser.

#### func  NewParserContext

```go
func NewParserContext() *ParserContext
```
NewParserContext creates a new ParserContext instance.

#### func (*ParserContext) Reset

```go
func (c *ParserContext) Reset()
```
Reset resets the parser context to its initial state.

#### type Patterns

```go
type Patterns struct {
}
```

Patterns holds compiled regular expressions for parsing Markdown syntax.

#### func  NewPatterns

```go
func NewPatterns() *Patterns
```
NewPatterns initializes and returns a new instance of Patterns with compiled
regular expressions.

#### type Token

```go
type Token struct {
	Type    TokenType
	Content string
	Args    []string
}
```

Token represents a single token in the input text.

#### type TokenType

```go
type TokenType int
```

TokenType represents the type of a token.

```go
const (
	TokenText             TokenType = iota // TokenText represents a regular text token.
	TokenHeadingUnderline                  // TokenHeadingUnderline represents a heading underline token.
	TokenTransBlock                        // TokenTransBlock represents a transition block token.
	TokenMeta                              // TokenMeta represents a metadata token.
	TokenDirective                         // TokenDirective represents a directive token.
	TokenCodeBlock                         // TokenCodeBlock represents a code block token.
	TokenBlankLine                         // TokenBlankLine represents a blank line token.
	TokenIndent                            // TokenIndent represents an indent token.
)
```
