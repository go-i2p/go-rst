Project Path: /home/idk/go/src/github.com/go-i2p/go-rst/pkg

Source Tree:

```
pkg
├── parser
│   ├── patterns.go
│   ├── parser_test.go
│   ├── paragraph.go
│   ├── title.go
│   ├── code.go
│   ├── directive.go
│   ├── lineblock.go
│   ├── doc.md
│   ├── context.go
│   ├── blockquote.go
│   ├── list.go
│   ├── strong.go
│   ├── emphasis.go
│   ├── meta.go
│   ├── doctest.go
│   ├── headiing.go
│   ├── subtitle.go
│   ├── table.go
│   ├── lexer.go
│   ├── link.go
│   ├── transition.go
│   └── parser.go
├── translator
│   ├── doc.md
│   └── translator.go
├── renderer
│   ├── doc.md
│   ├── markdown.go
│   ├── html.go
│   └── pdf.go
└── nodes
    ├── comment.go
    ├── paragraph.go
    ├── title.go
    ├── code.go
    ├── directive.go
    ├── lineblock.go
    ├── doc.md
    ├── blockquote.go
    ├── list.go
    ├── strong.go
    ├── meta.go
    ├── doctest.go
    ├── heading.go
    ├── subtitle.go
    ├── em.go
    ├── types.go
    ├── table.go
    ├── link.go
    ├── transition.go
    └── extra_util.go

```

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/patterns.go`:

```````go
// pkg/parser/patterns.go

package parser

import (
	"regexp"
)

// Patterns holds compiled regular expressions for parsing Markdown syntax.
type Patterns struct {
	headingUnderline *regexp.Regexp
	transBlock       *regexp.Regexp
	meta             *regexp.Regexp
	directive        *regexp.Regexp
	codeBlock        *regexp.Regexp
	blockQuote       *regexp.Regexp
	doctest          *regexp.Regexp
	doctestContinue  *regexp.Regexp
	doctestOutput    *regexp.Regexp
	lineBlock        *regexp.Regexp
	comment          *regexp.Regexp
	title            *regexp.Regexp
	subtitle         *regexp.Regexp
	transition       *regexp.Regexp
	bulletList       *regexp.Regexp
	enumList         *regexp.Regexp
}

// NewPatterns initializes and returns a new instance of Patterns with compiled regular expressions.
func NewPatterns() *Patterns {
	return &Patterns{
		headingUnderline: regexp.MustCompile(`^[=\-~]+$`),
		transBlock:       regexp.MustCompile(`{%\s*trans\s*%}(.*?){%\s*endtrans\s*%}`),
		meta:             regexp.MustCompile(`^\.\.\s+meta::`),
		directive:        regexp.MustCompile(`^\.\.\s+(\w+)::`),
		codeBlock:        regexp.MustCompile(`^\.\.\s+code::`),
		blockQuote:       regexp.MustCompile(`^(\s{4,})(.*?)(?:\s*--\s*(.*))?$`),
		doctest:          regexp.MustCompile(`^>>> (.+)\n((?:[^>].*\n)*)`),
		doctestContinue:  regexp.MustCompile(`^\.\.\.(.*$)`),
		doctestOutput:    regexp.MustCompile(`^([^>][^>][^>].*)$`),
		lineBlock:        regexp.MustCompile(`^\|(.*)$`),
		comment:          regexp.MustCompile(`^\.\.\s(.*)$`),
		title:            regexp.MustCompile(`^(={3,}|~{3,})\n(.+?)\n(?:={3,}|~{3,})$`),
		subtitle:         regexp.MustCompile(`^(-{3,})\n(.+?)\n(?:-{3,})$`),
		transition:       regexp.MustCompile(`^(\-{4,}|\={4,}|\*{4,})$`),
		bulletList:       regexp.MustCompile(`^(\s*)([-*+])(\s+)(.+)$`),
		enumList:         regexp.MustCompile(`^(\s*)(\d+|[a-zA-Z]|[ivxlcdm]+|[IVXLCDM]+)(\.\s+)(.+)$`),
	}
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/parser_test.go`:

```````go
package parser

// test the functionality of the parser package using the test package
// Use example restructuredText files embedded in the test functions

import (
	"testing"

	"github.com/go-i2p/go-rst/pkg/translator"
)

const (
	simpleDoc  = "example/doc.rst"
	complexDoc = "example/complexDoc.rst"
)

func TestParse(t *testing.T) {
	noopTranslator := translator.NewNoopTranslator()
	parser := NewParser(noopTranslator)
	doc := parser.Parse(simpleDoc)
	if doc == nil {
		t.Errorf("Expected a document, got nil")
	}
}

func TestParseTwo(t *testing.T) {
	noopTranslator := translator.NewNoopTranslator()
	parser := NewParser(noopTranslator)
	doc := parser.Parse(complexDoc)
	if doc == nil {
		t.Errorf("Expected a document, got nil")
	}
}

func TestParseEmpty(t *testing.T) {
	noopTranslator := translator.NewNoopTranslator()
	parser := NewParser(noopTranslator)
	doc := parser.Parse("")
	if len(doc) > 0 {
		t.Errorf("Expected empty, got a document")
	}
}

func TestParseNilTranslatorEmptyInput(t *testing.T) {
	parser := NewParser(nil)
	doc := parser.Parse("")
	if len(doc) > 0 {
		t.Errorf("Expected empty, got a document")
	}
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/paragraph.go`:

```````go
package parser

import (
	"strings"

	"github.com/go-i2p/go-rst/pkg/nodes"
)

func (p *Parser) processParagraph(line string, currentNode nodes.Node) nodes.Node {
	if strings.TrimSpace(line) == "" {
		if currentNode != nil {
			p.nodes = append(p.nodes, currentNode)
		}
		return nil
	}

	if currentNode == nil {
		return nodes.NewParagraphNode(line)
	}

	if currentNode.Type() == nodes.NodeParagraph {
		currentContent := currentNode.Content()
		currentNode.SetContent(currentContent + "\n" + line)
		return currentNode
	}

	return nodes.NewParagraphNode(line)
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/title.go`:

```````go
package parser

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/code.go`:

```````go
package parser

import (
	"strings"

	"github.com/go-i2p/go-rst/pkg/nodes"
)

func (p *Parser) processCodeBlock(line string, currentNode nodes.Node) nodes.Node {
	if currentNode == nil || currentNode.Type() != nodes.NodeCode {
		return currentNode
	}

	p.context.buffer = append(p.context.buffer, line)

	if strings.TrimSpace(line) == "" {
		codeNode := currentNode.(*nodes.CodeNode)
		content := strings.Join(p.context.buffer, "\n")
		codeNode.SetContent(content)
		p.nodes = append(p.nodes, codeNode)
		p.context.Reset()
		return nil
	}

	return currentNode
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/directive.go`:

```````go
package parser

import (
	"strings"

	"github.com/go-i2p/go-rst/pkg/nodes"
)

func (p *Parser) processDirectiveContent(line string, currentNode nodes.Node) nodes.Node {
	if currentNode == nil || currentNode.Type() != nodes.NodeDirective {
		return currentNode
	}

	directiveNode := currentNode.(*nodes.DirectiveNode)
	p.context.buffer = append(p.context.buffer, line)

	if strings.TrimSpace(line) == "" {
		content := strings.Join(p.context.buffer, "\n")
		directiveNode.SetRawContent(content)
		p.nodes = append(p.nodes, directiveNode)
		p.context.Reset()
		return nil
	}

	return currentNode
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/lineblock.go`:

```````go
package parser

import (
	"strings"

	"github.com/go-i2p/go-rst/pkg/nodes"
)

// processLineBlock handles parsing of line block tokens
// Line blocks are used for poetry-style content where line breaks are preserved
func (p *Parser) processLineBlock(content string, currentNode nodes.Node) nodes.Node {
	// If we already have a line block node, append the line
	if currentNode != nil && currentNode.Type() == nodes.NodeLineBlock {
		lineBlockNode := currentNode.(*nodes.LineBlockNode)
		currentLines := lineBlockNode.Lines()
		newLines := append(currentLines, strings.TrimSpace(content))
		return nodes.NewLineBlockNode(newLines)
	}

	// Otherwise create a new line block node
	return nodes.NewLineBlockNode([]string{strings.TrimSpace(content)})
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/doc.md`:

```````md
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

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/context.go`:

```````go
package parser

// ParserContext represents the current state of the parser.
type ParserContext struct {
	inMeta           bool
	inDirective      bool
	currentDirective string
	inCodeBlock      bool
	codeBlockIndent  int
	buffer           []string
}

// NewParserContext creates a new ParserContext instance.
func NewParserContext() *ParserContext {
	return &ParserContext{
		buffer: make([]string, 0),
	}
}

// Reset resets the parser context to its initial state.
func (c *ParserContext) Reset() {
	c.inMeta = false
	c.inDirective = false
	c.currentDirective = ""
	c.inCodeBlock = false
	c.codeBlockIndent = 0
	c.buffer = c.buffer[:0]
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/blockquote.go`:

```````go
package parser

import "github.com/go-i2p/go-rst/pkg/nodes"

func (p *Parser) processBlockQuote(content string, args []string, currentNode nodes.Node) nodes.Node {
	attribution := ""
	if len(args) > 0 {
		attribution = args[0]
	}

	if currentNode != nil && currentNode.Type() == nodes.NodeBlockQuote {
		// Add to existing block quote
		blockQuoteNode := currentNode.(*nodes.BlockQuoteNode)
		blockQuoteNode.AppendContent(content)
		return blockQuoteNode
	}

	return nodes.NewBlockQuoteNode(content, attribution)
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/list.go`:

```````go
package parser

import "github.com/go-i2p/go-rst/pkg/nodes"

func (p *Parser) processListItem(content string, args []string, isOrdered bool, currentNode nodes.Node) nodes.Node {
	indent := 0
	if len(args) > 0 {
		// Calculate indent from the whitespace prefix
		indent = len(args[0])
	}

	/*marker := ""
	if len(args) > 1 {
		marker = args[1]
	}*/

	// Create a new list item node
	listItem := nodes.NewListItemNode(content)

	// If we don't have a current node or it's not a list, create a new list
	if currentNode == nil || currentNode.Type() != nodes.NodeList {
		listNode := nodes.NewListNode(isOrdered)
		listNode.AppendChild(listItem)
		return listNode
	}

	// If we have a list but with different type or indent, create a new list
	listNode := currentNode.(*nodes.ListNode)
	if listNode.IsOrdered() != isOrdered || listNode.Indent() != indent {
		newList := nodes.NewListNode(isOrdered)
		newList.SetIndent(indent)
		newList.AppendChild(listItem)
		return newList
	}

	// Add item to existing list
	listNode.AppendChild(listItem)
	return listNode
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/strong.go`:

```````go
package parser

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/emphasis.go`:

```````go
package parser

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/meta.go`:

```````go
package parser

import (
	"strings"

	"github.com/go-i2p/go-rst/pkg/nodes"
)

func (p *Parser) processMetaContent(line string, currentNode nodes.Node) nodes.Node {
	if currentNode == nil || currentNode.Type() != nodes.NodeMeta {
		return currentNode
	}

	parts := strings.SplitN(line, ":", 2)
	if len(parts) != 2 {
		return currentNode
	}

	key := strings.TrimSpace(parts[0])
	value := strings.TrimSpace(parts[1])

	node := nodes.NewMetaNode(key, value)
	p.nodes = append(p.nodes, node)
	p.context.inMeta = false
	return nil
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/doctest.go`:

```````go
package parser

import "strings"

func processDoctestBlock(line string) Token {
	// Handle doctest blocks
	if strings.HasPrefix(line, ">>> ") {
		return Token{
			Type:    TokenDoctest,
			Content: line[4:], // Remove the leading '>>> '
		}
	}
	return Token{}
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/headiing.go`:

```````go
package parser

import (
	"strings"

	"github.com/go-i2p/go-rst/pkg/nodes"
)

func (p *Parser) processHeading(content, underline string) nodes.Node {
	level := 1
	switch underline[0] {
	case '-':
		level = 2
	case '~':
		level = 3
	}

	node := nodes.NewHeadingNode(strings.TrimSpace(content), level)
	p.nodes = append(p.nodes, node)
	return nil
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/subtitle.go`:

```````go
package parser

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/table.go`:

```````go
package parser

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/lexer.go`:

```````go
// pkg/parser/lexer.go

package parser

import (
	"strings"
)

// TokenType represents the type of a token.
type TokenType int

const (
	TokenText             TokenType = iota // TokenText represents a regular text token.
	TokenHeadingUnderline                  // TokenHeadingUnderline represents a heading underline token.
	TokenTransBlock                        // TokenTransBlock represents a transition block token.
	TokenMeta                              // TokenMeta represents a metadata token.
	TokenDirective                         // TokenDirective represents a directive token.
	TokenCodeBlock                         // TokenCodeBlock represents a code block token.
	TokenBlankLine                         // TokenBlankLine represents a blank line token.
	TokenIndent                            // TokenIndent represents an indent token.
	TokenBlockQuote                        // TokenBlockQuote represents a block quote token.
	TokenComment                           // TokenComment represents a comment token.
	TokenBulletList                        // TokenBulletList represents a bullet list item token.
	TokenEnumList                          // TokenEnumList represents an enumerated list item token.
	TokenDoctest                           // TokenDoctest represents a doctest token.
	TokenLineBlock                         // TokenLineBlock represents a line block token.
	TokenTransition                        // TokenTransition represents a transition token.
)

// Token represents a single token in the input text.
type Token struct {
	Type    TokenType
	Content string
	Args    []string
}

// Lexer represents a lexer for the input text.
type Lexer struct {
	patterns *Patterns
}

// NewLexer creates a new Lexer instance.
func NewLexer() *Lexer {
	return &Lexer{
		patterns: NewPatterns(),
	}
}

// Tokenize tokenizes a single line of input text.
func (l *Lexer) Tokenize(line string) Token {
	// Handle blank lines
	if strings.TrimSpace(line) == "" {
		return Token{Type: TokenBlankLine}
	}

	// Calculate indentation
	indent := 0
	for _, r := range line {
		if r == ' ' {
			indent++
		} else if r == '\t' {
			indent += 4
		} else {
			break
		}
	}

	line = strings.TrimLeft(line, " \t")

	// Check for heading underline
	if l.patterns.headingUnderline.MatchString(line) {
		return Token{
			Type:    TokenHeadingUnderline,
			Content: line,
		}
	}

	// Check for translation blocks
	if matches := l.patterns.transBlock.FindStringSubmatch(line); len(matches) > 1 {
		return Token{
			Type:    TokenTransBlock,
			Content: matches[1],
		}
	}

	// Check for meta directive
	if l.patterns.meta.MatchString(line) {
		return Token{
			Type: TokenMeta,
		}
	}

	// Check for code block
	if l.patterns.codeBlock.MatchString(line) {
		args := parseDirectiveArgs(line)
		return Token{
			Type: TokenCodeBlock,
			Args: args,
		}
	}

	// Check for other directives
	if matches := l.patterns.directive.FindStringSubmatch(line); len(matches) > 1 {
		args := parseDirectiveArgs(line)
		return Token{
			Type:    TokenDirective,
			Content: matches[1],
			Args:    args,
		}
	}

	// Check for block quote
	if matches := l.patterns.blockQuote.FindStringSubmatch(line); len(matches) > 1 {
		attribution := ""
		if len(matches) > 2 {
			attribution = matches[3]
		}
		return Token{
			Type:    TokenBlockQuote,
			Content: matches[2],
			Args:    []string{attribution},
		}
	}

	// Check for comment
	if matches := l.patterns.comment.FindStringSubmatch(line); len(matches) > 1 {
		return Token{
			Type:    TokenComment,
			Content: matches[1],
		}
	}
	// Check for bullet list
	if matches := l.patterns.bulletList.FindStringSubmatch(line); len(matches) > 1 {
		return Token{
			Type:    TokenBulletList,
			Content: matches[4],
			Args:    []string{matches[1], matches[2]}, // indent, bullet type
		}
	}

	// Check for enumerated list
	if matches := l.patterns.enumList.FindStringSubmatch(line); len(matches) > 1 {
		return Token{
			Type:    TokenEnumList,
			Content: matches[4],
			Args:    []string{matches[1], matches[2]}, // indent, marker
		}
	}

	// Check for line block (poetry-style line with | prefix)
	if matches := l.patterns.lineBlock.FindStringSubmatch(line); len(matches) > 0 {
		return Token{
			Type:    TokenLineBlock,
			Content: strings.TrimSpace(matches[1]), // The content after the | character
		}
	}

	// Check for transitions
	if l.patterns.IsTransition(line) {
		transChar := l.patterns.TransitionChar(line)
		return Token{
			Type:    TokenTransition,
			Content: string(transChar),
		}
	}

	// Regular text
	return Token{
		Type:    TokenText,
		Content: line,
	}
}

func parseDirectiveArgs(line string) []string {
	parts := strings.SplitN(line, "::", 2)
	if len(parts) != 2 {
		return nil
	}

	args := strings.Fields(strings.TrimSpace(parts[1]))
	return args
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/link.go`:

```````go
package parser

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/transition.go`:

```````go
package parser

import (
	"strings"

	"github.com/go-i2p/go-rst/pkg/nodes"
)

// IsTransition checks if a line is a transition
func (p *Patterns) IsTransition(line string) bool {
	// A transition is a line with 4+ repeated punctuation characters
	return len(strings.TrimSpace(line)) >= 4 && p.transition.MatchString(strings.TrimSpace(line))
}

// TransitionChar extracts the character used in the transition
func (p *Patterns) TransitionChar(line string) rune {
	trimmed := strings.TrimSpace(line)
	if len(trimmed) > 0 {
		return rune(trimmed[0])
	}
	return '-' // Default to hyphen if empty (shouldn't happen)
}

// processTransition handles the parsing of transition sections
// A transition is a horizontal line separator typically used between sections
func (p *Parser) processTransition(content string) *nodes.TransitionNode {
	// Extract the character used in the transition
	var transChar rune = '-' // Default to hyphen
	if len(content) > 0 {
		transChar = rune(content[0])
	}

	// Create a new transition node with the character
	return nodes.NewTransitionNode(transChar)
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/parser/parser.go`:

```````go
// pkg/parser/parser.go

package parser

import (
	"bufio"
	"strings"

	"github.com/go-i2p/go-rst/pkg/nodes"
	"github.com/go-i2p/go-rst/pkg/translator"
)

// Parser is a struct that holds the state of the parser.
type Parser struct {
	nodes      []nodes.Node
	translator translator.Translator
	context    *ParserContext
	patterns   *Patterns
	lexer      *Lexer
}

// NewParser creates a new Parser instance.
func NewParser(trans translator.Translator) *Parser {
	return &Parser{
		nodes:      make([]nodes.Node, 0),
		translator: trans,
		context:    NewParserContext(),
		patterns:   NewPatterns(),
		lexer:      NewLexer(),
	}
}

// Parse takes a string of reStructuredText content and returns a slice of Node instances.
func (p *Parser) Parse(content string) []nodes.Node {
	scanner := bufio.NewScanner(strings.NewReader(content))
	var currentNode nodes.Node
	var prevToken Token
	p.nodes = make([]nodes.Node, 0) // Clear existing nodes

	for scanner.Scan() {
		line := scanner.Text()
		token := p.lexer.Tokenize(line)

		if newNode := p.processToken(token, prevToken, currentNode); newNode != nil {
			// Only append if we actually have a new node
			if currentNode != nil && currentNode != newNode {
				p.nodes = append(p.nodes, currentNode)
			}
			currentNode = newNode
		}
		prevToken = token
	}

	// Add final node if exists and not already added
	if currentNode != nil && (len(p.nodes) == 0 || p.nodes[len(p.nodes)-1] != currentNode) {
		p.nodes = append(p.nodes, currentNode)
	}

	return p.nodes
}

func (p *Parser) processToken(token, prevToken Token, currentNode nodes.Node) nodes.Node {
	// translatedContent := p.translator.Translate(token.Content)
	// token.Content = translatedContent
	switch token.Type {
	case TokenBulletList:
		return p.processListItem(token.Content, token.Args, false, currentNode)
	case TokenEnumList:
		return p.processListItem(token.Content, token.Args, true, currentNode)
	case TokenBlockQuote:
		return p.processBlockQuote(token.Content, token.Args, currentNode)
	case TokenComment:
		return nodes.NewCommentNode(token.Content)
	case TokenTransBlock:
		// Always create a new node for translation blocks
		translatedContent := p.translator.Translate(strings.TrimSpace(token.Content))
		return nodes.NewParagraphNode(translatedContent)
	case TokenHeadingUnderline:
		if prevToken.Type == TokenText {
			return p.processHeading(prevToken.Content, token.Content)
		}

	case TokenMeta:
		p.context.inMeta = true
		return nodes.NewMetaNode("", "")

	case TokenCodeBlock:
		p.context.inCodeBlock = true
		p.context.codeBlockIndent = 4
		language := ""
		if len(token.Args) > 0 {
			language = token.Args[0]
		}
		return nodes.NewCodeNode(language, "", false)

	case TokenDirective:
		p.context.inDirective = true
		p.context.currentDirective = token.Content
		return nodes.NewDirectiveNode(token.Content, token.Args)

	case TokenLineBlock:
		// Check if we're already in a line block node
		if lineBlock, ok := currentNode.(*nodes.LineBlockNode); ok {
			// Add this line to the existing line block
			lines := lineBlock.Lines()
			lines = append(lines, token.Content)
			newLineBlock := nodes.NewLineBlockNode(lines)
			return newLineBlock
		}
		// Create a new line block node
		return nodes.NewLineBlockNode([]string{token.Content})

	case TokenText:
		if p.context.inCodeBlock {
			return p.processCodeBlock(token.Content, currentNode)
		}
		if p.context.inMeta {
			return p.processMetaContent(token.Content, currentNode)
		}
		if p.context.inDirective {
			return p.processDirectiveContent(token.Content, currentNode)
		}
		return p.processParagraph(token.Content, currentNode)
	case TokenTransition:
		// For transitions, we create a new transition node with the character used
		if len(token.Content) > 0 {
			return p.processTransition(token.Content)
		}
		return nodes.NewTransitionNode('-') // Default to hyphen if empty
	}

	return currentNode
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/translator/doc.md`:

```````md
# translator
--
    import "github.com/go-i2p/go-rst/pkg/translator"


## Usage

#### type NoopTranslator

```go
type NoopTranslator struct{}
```

NoopTranslator implements Translator interface but doesn't translate

#### func  NewNoopTranslator

```go
func NewNoopTranslator() *NoopTranslator
```
NewNoopTranslator returns a new NoopTranslator

#### func (*NoopTranslator) Translate

```go
func (t *NoopTranslator) Translate(text string) string
```
Translate returns the same text it receives(NoopTranslator)

#### type POTranslator

```go
type POTranslator struct {
}
```

POTranslator implements Translator interface using a PO file

#### func  NewPOTranslator

```go
func NewPOTranslator(poFile string) (*POTranslator, error)
```
NewPOTranslator returns a new POTranslator

#### func (*POTranslator) Translate

```go
func (t *POTranslator) Translate(text string) string
```
Translate returns the translated text if it exists in the PO file, otherwise it
returns the original text

#### type Translator

```go
type Translator interface {
	Translate(text string) string
}
```

Translator is an interface for translating text

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/translator/translator.go`:

```````go
package translator

import (
	"github.com/leonelquinteros/gotext"
)

// Translator is an interface for translating text
type Translator interface {
	Translate(text string) string
}

// POTranslator implements Translator interface using a PO file
type POTranslator struct {
	po *gotext.Po
}

// NewPOTranslator returns a new POTranslator
func NewPOTranslator(poFile string) (*POTranslator, error) {
	translator := &POTranslator{
		po: gotext.NewPo(),
	}

	// If no PO file is provided, return a pass-through translator
	if poFile == "" {
		return translator, nil
	}

	// Parse PO file
	translator.po.ParseFile(poFile)

	return translator, nil
}

// Translate returns the translated text if it exists in the PO file, otherwise it returns the original text
func (t *POTranslator) Translate(text string) string {
	if t.po == nil {
		return text
	}
	translated := t.po.Get(text)
	if translated == "" {
		return text
	}
	return translated
}

// NoopTranslator implements Translator interface but doesn't translate
type NoopTranslator struct{}

// NewNoopTranslator returns a new NoopTranslator
func NewNoopTranslator() *NoopTranslator {
	return &NoopTranslator{}
}

// Translate returns the same text it receives(NoopTranslator)
func (t *NoopTranslator) Translate(text string) string {
	return text
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/renderer/doc.md`:

```````md
# renderer
--
    import "github.com/go-i2p/go-rst/pkg/renderer"


## Usage

#### type HTMLRenderer

```go
type HTMLRenderer struct {
}
```

HTMLRederer is a renderer that renders nodes to HTML.

#### func  NewHTMLRenderer

```go
func NewHTMLRenderer() *HTMLRenderer
```
NewHTMLRederer creates a new HTMLRederer.

#### func (*HTMLRenderer) Render

```go
func (r *HTMLRenderer) Render(nodes []nodes.Node) string
```
Render renders nodes to HTML.

#### func (*HTMLRenderer) RenderPretty

```go
func (r *HTMLRenderer) RenderPretty(nodes []nodes.Node) string
```
RenderPretty renders the given nodes as pretty-formatted HTML.

#### type MarkdownRenderer

```go
type MarkdownRenderer struct {
}
```

MarkdownRenderer implements a Markdown renderer with the same interface as
HTMLRenderer

#### func  NewMarkdownRenderer

```go
func NewMarkdownRenderer() *MarkdownRenderer
```
NewMarkdownRenderer creates a new Markdown renderer

#### func (*MarkdownRenderer) Render

```go
func (r *MarkdownRenderer) Render(nodes []nodes.Node) error
```
Render renders a slice of nodes to Markdown

#### func (*MarkdownRenderer) RenderChildren

```go
func (r *MarkdownRenderer) RenderChildren(node nodes.Node) error
```
RenderChildren renders child nodes

#### func (*MarkdownRenderer) RenderCode

```go
func (r *MarkdownRenderer) RenderCode(node *nodes.CodeNode) error
```
RenderCode renders a code node

#### func (*MarkdownRenderer) RenderDirective

```go
func (r *MarkdownRenderer) RenderDirective(node *nodes.DirectiveNode) error
```
RenderDirective renders a directive node

#### func (*MarkdownRenderer) RenderEmphasis

```go
func (r *MarkdownRenderer) RenderEmphasis(node *nodes.EmphasisNode) error
```
RenderEmphasis renders an emphasis node

#### func (*MarkdownRenderer) RenderHeading

```go
func (r *MarkdownRenderer) RenderHeading(node *nodes.HeadingNode) error
```
RenderHeading renders a heading node

#### func (*MarkdownRenderer) RenderLink

```go
func (r *MarkdownRenderer) RenderLink(node *nodes.LinkNode) error
```
RenderLink renders a link node

#### func (*MarkdownRenderer) RenderList

```go
func (r *MarkdownRenderer) RenderList(node *nodes.ListNode) error
```
RenderList renders a list node

#### func (*MarkdownRenderer) RenderListItem

```go
func (r *MarkdownRenderer) RenderListItem(node *nodes.ListItemNode) error
```
RenderListItem renders a list item node

#### func (*MarkdownRenderer) RenderMeta

```go
func (r *MarkdownRenderer) RenderMeta(node *nodes.MetaNode) error
```
RenderMeta renders a meta node

#### func (*MarkdownRenderer) RenderNode

```go
func (r *MarkdownRenderer) RenderNode(node nodes.Node) error
```
RenderNode renders a single node to Markdown

#### func (*MarkdownRenderer) RenderParagraph

```go
func (r *MarkdownRenderer) RenderParagraph(node *nodes.ParagraphNode) error
```
RenderParagraph renders a paragraph node

#### func (*MarkdownRenderer) RenderStrong

```go
func (r *MarkdownRenderer) RenderStrong(node *nodes.StrongNode) error
```
RenderStrong renders a strong node

#### func (*MarkdownRenderer) RenderTable

```go
func (r *MarkdownRenderer) RenderTable(node *nodes.TableNode) error
```
RenderTable renders a table node

#### func (*MarkdownRenderer) String

```go
func (r *MarkdownRenderer) String() string
```
String returns the rendered markdown as a string

#### type PDFRenderer

```go
type PDFRenderer struct {
}
```

PDFRenderer implements rendering RST nodes to PDF format

#### func  NewPDFRenderer

```go
func NewPDFRenderer() *PDFRenderer
```
NewPDFRenderer creates a new PDF renderer

#### func (*PDFRenderer) Render

```go
func (r *PDFRenderer) Render(nodes []nodes.Node) error
```
Render renders a slice of nodes to PDF

#### func (*PDFRenderer) SaveToFile

```go
func (r *PDFRenderer) SaveToFile(filename string) error
```
SaveToFile saves the PDF to a file

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/renderer/markdown.go`:

```````go
// pkg/renderer/markdown.go

package renderer

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/go-i2p/go-rst/pkg/nodes"
)

// MarkdownRenderer implements a Markdown renderer with the same interface as HTMLRenderer
type MarkdownRenderer struct {
	output bytes.Buffer
}

// NewMarkdownRenderer creates a new Markdown renderer
func NewMarkdownRenderer() *MarkdownRenderer {
	return &MarkdownRenderer{}
}

// pkg/renderer/markdown.go

// ... (previous imports and struct definition remain the same)

// Render renders a slice of nodes to Markdown
func (r *MarkdownRenderer) Render(nodes []nodes.Node) error {
	for _, node := range nodes {
		if err := r.RenderNode(node); err != nil {
			return err
		}
	}
	return nil
}

// RenderNode renders a single node to Markdown
func (r *MarkdownRenderer) RenderNode(node nodes.Node) error {
	switch n := node.(type) {
	case *nodes.HeadingNode:
		return r.RenderHeading(n)
	case *nodes.ParagraphNode:
		return r.RenderParagraph(n)
	case *nodes.ListNode:
		return r.RenderList(n)
	case *nodes.ListItemNode:
		return r.RenderListItem(n)
	case *nodes.EmphasisNode:
		return r.RenderEmphasis(n)
	case *nodes.StrongNode:
		return r.RenderStrong(n)
	case *nodes.LinkNode:
		return r.RenderLink(n)
	case *nodes.CodeNode:
		return r.RenderCode(n)
	case *nodes.TableNode:
		return r.RenderTable(n)
	case *nodes.DirectiveNode:
		return r.RenderDirective(n)
	case *nodes.MetaNode:
		return r.RenderMeta(n)
	default:
		return r.RenderChildren(node)
	}
}

// RenderHeading renders a heading node
func (r *MarkdownRenderer) RenderHeading(node *nodes.HeadingNode) error {
	r.output.WriteString("\n")
	r.output.WriteString(strings.Repeat("#", node.Level()))
	r.output.WriteString(" ")
	r.output.WriteString(node.Content())
	r.output.WriteString("\n")
	return r.RenderChildren(node)
}

// RenderParagraph renders a paragraph node
func (r *MarkdownRenderer) RenderParagraph(node *nodes.ParagraphNode) error {
	r.output.WriteString("\n")
	r.output.WriteString(node.Content())
	r.output.WriteString("\n")
	return r.RenderChildren(node)
}

// RenderList renders a list node
func (r *MarkdownRenderer) RenderList(node *nodes.ListNode) error {
	r.output.WriteString("\n")
	return r.RenderChildren(node)
}

// RenderListItem renders a list item node
func (r *MarkdownRenderer) RenderListItem(node *nodes.ListItemNode) error {
	// Default to unordered list items with "-"
	r.output.WriteString("- ")
	r.output.WriteString(node.Content())
	r.output.WriteString("\n")
	return r.RenderChildren(node)
}

// RenderEmphasis renders an emphasis node
func (r *MarkdownRenderer) RenderEmphasis(node *nodes.EmphasisNode) error {
	r.output.WriteString("*")
	r.output.WriteString(node.Content())
	r.output.WriteString("*")
	return r.RenderChildren(node)
}

// RenderStrong renders a strong node
func (r *MarkdownRenderer) RenderStrong(node *nodes.StrongNode) error {
	r.output.WriteString("**")
	r.output.WriteString(node.Content())
	r.output.WriteString("**")
	return r.RenderChildren(node)
}

// RenderLink renders a link node
func (r *MarkdownRenderer) RenderLink(node *nodes.LinkNode) error {
	if title := node.Title(); title != "" {
		r.output.WriteString(fmt.Sprintf("[%s](%s \"%s\")", node.Content(), node.URL(), title))
	} else {
		r.output.WriteString(fmt.Sprintf("[%s](%s)", node.Content(), node.URL()))
	}
	return nil
}

// RenderCode renders a code node
func (r *MarkdownRenderer) RenderCode(node *nodes.CodeNode) error {
	r.output.WriteString("\n```")
	if node.Language() != "" {
		r.output.WriteString(node.Language())
	}
	r.output.WriteString("\n")
	r.output.WriteString(node.Content())
	r.output.WriteString("\n```\n")
	return nil
}

// RenderTable renders a table node
func (r *MarkdownRenderer) RenderTable(node *nodes.TableNode) error {
	headers := node.Headers()
	rows := node.Rows()

	// Headers
	r.output.WriteString("\n|")
	for _, header := range headers {
		r.output.WriteString(fmt.Sprintf(" %s |", header))
	}
	r.output.WriteString("\n|")

	// Separator
	for range headers {
		r.output.WriteString(" --- |")
	}
	r.output.WriteString("\n")

	// Rows
	for _, row := range rows {
		r.output.WriteString("|")
		for _, cell := range row {
			r.output.WriteString(fmt.Sprintf(" %s |", cell))
		}
		r.output.WriteString("\n")
	}
	r.output.WriteString("\n")
	return nil
}

// RenderDirective renders a directive node
func (r *MarkdownRenderer) RenderDirective(node *nodes.DirectiveNode) error {
	switch node.Name() {
	case "image":
		if len(node.Arguments()) > 0 {
			r.output.WriteString(fmt.Sprintf("\n![%s](%s)\n", node.Content(), node.Arguments()[0]))
		}
	case "note":
		r.output.WriteString("\n> **Note**\n")
		r.output.WriteString("> " + strings.Replace(node.RawContent(), "\n", "\n> ", -1) + "\n")
	case "warning":
		r.output.WriteString("\n> **Warning**\n")
		r.output.WriteString("> " + strings.Replace(node.RawContent(), "\n", "\n> ", -1) + "\n")
	default:
		r.output.WriteString(fmt.Sprintf("\n<!-- %s: %s -->\n", node.Name(), node.RawContent()))
	}
	return r.RenderChildren(node)
}

// RenderMeta renders a meta node
func (r *MarkdownRenderer) RenderMeta(node *nodes.MetaNode) error {
	if r.output.Len() == 0 {
		// If at start of document, use YAML front matter
		r.output.WriteString("---\n")
		r.output.WriteString(fmt.Sprintf("%s: %s\n", node.Key(), node.Content()))
		r.output.WriteString("---\n")
	} else {
		// Otherwise use HTML comment
		r.output.WriteString(fmt.Sprintf("\n<!-- %s: %s -->\n", node.Key(), node.Content()))
	}
	return nil
}

// RenderChildren renders child nodes
func (r *MarkdownRenderer) RenderChildren(node nodes.Node) error {
	for _, child := range node.Children() {
		if err := r.RenderNode(child); err != nil {
			return err
		}
	}
	return nil
}

// String returns the rendered markdown as a string
func (r *MarkdownRenderer) String() string {
	return r.output.String()
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/renderer/html.go`:

```````go
// pkg/renderer/html.go

package renderer

import (
	"bytes"
	"fmt"
	"html"
	"strings"

	"github.com/go-i2p/go-rst/pkg/nodes"
	"github.com/yosssi/gohtml"
)

// HTMLRederer is a renderer that renders nodes to HTML.
type HTMLRenderer struct {
	buffer bytes.Buffer
}

// NewHTMLRederer creates a new HTMLRederer.
func NewHTMLRenderer() *HTMLRenderer {
	return &HTMLRenderer{}
}

// Render renders nodes to HTML.
func (r *HTMLRenderer) Render(nodes []nodes.Node) string {
	r.buffer.Reset()

	r.buffer.WriteString("<!DOCTYPE html>\n<html>\n<head>\n")
	r.renderMeta(nodes)
	r.buffer.WriteString("</head>\n<body>\n")

	for _, node := range nodes {
		r.renderNode(node)
	}

	r.buffer.WriteString("</body>\n</html>")
	return r.buffer.String()
}

func (r *HTMLRenderer) renderMeta(nodelist []nodes.Node) {
	r.buffer.WriteString("<meta charset=\"UTF-8\">\n")

	for _, node := range nodelist {
		switch n := node.(type) {
		case *nodes.MetaNode:
			r.buffer.WriteString(fmt.Sprintf("<meta name=\"%s\" content=\"%s\">\n",
				html.EscapeString(n.Key()),
				html.EscapeString(n.Content())))
		}
	}
}

func (r *HTMLRenderer) renderNode(node nodes.Node) {
	switch n := node.(type) {
	case *nodes.HeadingNode:
		r.buffer.WriteString(fmt.Sprintf("<h%d>%s</h%d>\n",
			n.Level(),
			html.EscapeString(n.Content()),
			n.Level()))

	case *nodes.ParagraphNode:
		r.buffer.WriteString(fmt.Sprintf("<p>%s</p>\n",
			html.EscapeString(n.Content())))

	case *nodes.ListNode:
		tag := "ul"
		if n.IsOrdered() {
			tag = "ol"
		}
		r.buffer.WriteString(fmt.Sprintf("<%s>\n", tag))
		for _, child := range n.Children() {
			if item, ok := child.(*nodes.ListItemNode); ok {
				r.buffer.WriteString(fmt.Sprintf("<li>%s</li>\n",
					html.EscapeString(item.Content())))
			}
		}
		r.buffer.WriteString(fmt.Sprintf("</%s>\n", tag))

	case *nodes.LinkNode:
		r.buffer.WriteString(fmt.Sprintf("<a href=\"%s\" title=\"%s\">%s</a>",
			html.EscapeString(n.URL()),
			html.EscapeString(n.Title()),
			html.EscapeString(n.Content())))

	case *nodes.EmphasisNode:
		r.buffer.WriteString(fmt.Sprintf("<em>%s</em>",
			html.EscapeString(n.Content())))

	case *nodes.StrongNode:
		r.buffer.WriteString(fmt.Sprintf("<strong>%s</strong>",
			html.EscapeString(n.Content())))

	case *nodes.CodeNode:
		r.buffer.WriteString(fmt.Sprintf("<pre><code class=\"language-%s\">%s</code></pre>\n",
			html.EscapeString(n.Language()),
			html.EscapeString(n.Content())))

	case *nodes.TableNode:
		r.renderTable(n)

	case *nodes.DirectiveNode:
		r.renderDirective(n)
	case *nodes.BlockQuoteNode:
		r.buffer.WriteString("<blockquote>")
		r.buffer.WriteString(html.EscapeString(n.Content()))
		if attr := n.Attribution(); attr != "" {
			r.buffer.WriteString("<cite>")
			r.buffer.WriteString(html.EscapeString(attr))
			r.buffer.WriteString("</cite>")
		}
		r.buffer.WriteString("</blockquote>\n")
	case *nodes.DoctestNode:
		r.buffer.WriteString("<div class=\"doctest\">")
		r.buffer.WriteString("<pre class=\"doctest-command\">>> ")
		r.buffer.WriteString(html.EscapeString(n.Command()))
		r.buffer.WriteString("</pre>")
		if n.Expected() != "" {
			r.buffer.WriteString("<pre class=\"doctest-output\">")
			r.buffer.WriteString(html.EscapeString(n.Expected()))
			r.buffer.WriteString("</pre>")
		}
		r.buffer.WriteString("</div>\n")
	case *nodes.LineBlockNode:
		r.buffer.WriteString("<div class=\"line-block\">")
		for _, line := range n.Lines() {
			r.buffer.WriteString("<div class=\"line\">")
			r.buffer.WriteString(html.EscapeString(strings.TrimSpace(line)))
			r.buffer.WriteString("</div>\n")
		}
		r.buffer.WriteString("</div>\n")
	case *nodes.CommentNode:
		r.buffer.WriteString("<!-- ")
		r.buffer.WriteString(html.EscapeString(n.Content()))
		r.buffer.WriteString(" -->\n")
	case *nodes.TitleNode:
		r.buffer.WriteString(fmt.Sprintf("<h1 class=\"title\">%s</h1>\n",
			html.EscapeString(n.Content())))
	case *nodes.SubtitleNode:
		r.buffer.WriteString(fmt.Sprintf("<h2 class=\"subtitle\">%s</h2>\n",
			html.EscapeString(n.Content())))
	case *nodes.TransitionNode:
		r.buffer.WriteString("<hr class=\"docutils\">\n")
	}
}

func (r *HTMLRenderer) renderTable(table *nodes.TableNode) {
	r.buffer.WriteString("<table>\n")

	// Render headers
	if len(table.Headers()) > 0 {
		r.buffer.WriteString("<thead><tr>\n")
		for _, header := range table.Headers() {
			r.buffer.WriteString(fmt.Sprintf("<th>%s</th>",
				html.EscapeString(header)))
		}
		r.buffer.WriteString("</tr></thead>\n")
	}

	// Render rows
	r.buffer.WriteString("<tbody>\n")
	for _, row := range table.Rows() {
		r.buffer.WriteString("<tr>\n")
		for _, cell := range row {
			r.buffer.WriteString(fmt.Sprintf("<td>%s</td>",
				html.EscapeString(cell)))
		}
		r.buffer.WriteString("</tr>\n")
	}
	r.buffer.WriteString("</tbody></table>\n")
}

func (r *HTMLRenderer) renderDirective(directive *nodes.DirectiveNode) {
	switch directive.Name() {
	case "image":
		if len(directive.Arguments()) > 0 {
			alt := ""
			if len(directive.Arguments()) > 1 {
				alt = strings.Join(directive.Arguments()[1:], " ")
			}
			r.buffer.WriteString(fmt.Sprintf("<img src=\"%s\" alt=\"%s\">\n",
				html.EscapeString(directive.Arguments()[0]),
				html.EscapeString(alt)))
		}

	case "note":
		r.buffer.WriteString(fmt.Sprintf("<div class=\"note\">%s</div>\n",
			html.EscapeString(directive.RawContent())))

	case "warning":
		r.buffer.WriteString(fmt.Sprintf("<div class=\"warning\">%s</div>\n",
			html.EscapeString(directive.RawContent())))
	}
}

// RenderPretty renders the given nodes as pretty-formatted HTML.
func (r *HTMLRenderer) RenderPretty(nodes []nodes.Node) string {
	// First get the regular HTML output
	rawHTML := r.Render(nodes)

	// Let gohtml handle the formatting
	prettyHTML := gohtml.Format(rawHTML)

	return prettyHTML
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/renderer/pdf.go`:

```````go
// pkg/renderer/pdf.go

package renderer

import (
	"fmt"
	"strings"

	"github.com/go-i2p/go-rst/pkg/nodes"
	"github.com/jung-kurt/gofpdf"
)

// PDFRenderer implements rendering RST nodes to PDF format
type PDFRenderer struct {
	pdf        *gofpdf.Fpdf
	marginLeft float64
	marginTop  float64
	fontSize   float64
	lineHeight float64
	indent     float64
}

// NewPDFRenderer creates a new PDF renderer
func NewPDFRenderer() *PDFRenderer {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	return &PDFRenderer{
		pdf:        pdf,
		marginLeft: 20,
		marginTop:  20,
		fontSize:   12,
		lineHeight: 6,
		indent:     10,
	}
}

// Render renders a slice of nodes to PDF
func (r *PDFRenderer) Render(nodes []nodes.Node) error {
	for _, node := range nodes {
		if err := r.renderNode(node); err != nil {
			return err
		}
	}
	return nil
}

// renderNode handles individual node rendering
func (r *PDFRenderer) renderNode(node nodes.Node) error {
	switch n := node.(type) {
	case *nodes.HeadingNode:
		return r.renderHeading(n)
	case *nodes.ParagraphNode:
		return r.renderParagraph(n)
	case *nodes.ListNode:
		return r.renderList(n)
	case *nodes.CodeNode:
		return r.renderCode(n)
	case *nodes.TableNode:
		return r.renderTable(n)
	case *nodes.DirectiveNode:
		return r.renderDirective(n)
	default:
		return r.renderChildren(node)
	}
}

// SaveToFile saves the PDF to a file
func (r *PDFRenderer) SaveToFile(filename string) error {
	return r.pdf.OutputFileAndClose(filename)
}

func (r *PDFRenderer) renderHeading(node *nodes.HeadingNode) error {
	// Calculate font size based on heading level
	fontSize := r.fontSize + (6 - float64(node.Level())*2)
	r.pdf.SetFont("Arial", "B", fontSize)

	// Add some spacing before heading
	r.pdf.Ln(r.lineHeight)

	r.pdf.Cell(0, r.lineHeight, node.Content())
	r.pdf.Ln(r.lineHeight * 1.5)

	// Reset font
	r.pdf.SetFont("Arial", "", r.fontSize)
	return r.renderChildren(node)
}

func (r *PDFRenderer) renderParagraph(node *nodes.ParagraphNode) error {
	r.pdf.MultiCell(0, r.lineHeight, node.Content(), "", "", false)
	r.pdf.Ln(r.lineHeight)
	return r.renderChildren(node)
}

func (r *PDFRenderer) renderList(node *nodes.ListNode) error {
	currentX := r.pdf.GetX()
	currentY := r.pdf.GetY()

	for i, child := range node.Children() {
		if node.IsOrdered() {
			// Ordered list: use numbers
			r.pdf.SetXY(currentX, currentY)
			r.pdf.Cell(10, r.lineHeight, fmt.Sprintf("%d.", i+1))
			r.pdf.SetX(currentX + r.indent)
		} else {
			// Unordered list: use bullets
			r.pdf.SetXY(currentX, currentY)
			r.pdf.Cell(10, r.lineHeight, "•")
			r.pdf.SetX(currentX + r.indent)
		}

		if listItem, ok := child.(*nodes.ListItemNode); ok {
			r.pdf.MultiCell(0, r.lineHeight, listItem.Content(), "", "", false)
		}

		currentY = r.pdf.GetY() + r.lineHeight
	}

	r.pdf.Ln(r.lineHeight)
	return nil
}

func (r *PDFRenderer) renderCode(node *nodes.CodeNode) error {
	// Set monospace font for code
	r.pdf.SetFont("Courier", "", r.fontSize)

	// Add a light gray background
	startY := r.pdf.GetY()
	content := node.Content()
	lines := strings.Split(content, "\n")

	// Calculate height
	height := float64(len(lines)) * r.lineHeight

	// Draw background
	r.pdf.SetFillColor(240, 240, 240)
	r.pdf.Rect(r.marginLeft-2, startY-2, 170, height+4, "F")

	// Add language identifier if present
	if node.Language() != "" {
		r.pdf.SetFont("Arial", "I", r.fontSize-2)
		r.pdf.Text(r.marginLeft, startY-4, node.Language())
		r.pdf.SetFont("Courier", "", r.fontSize)
	}

	// Render code content
	for _, line := range lines {
		if node.LineNumbers() {
			// Add line numbers if enabled
			lineNum := fmt.Sprintf("%3d ", r.pdf.PageCount())
			r.pdf.Text(r.marginLeft, r.pdf.GetY(), lineNum)
			r.pdf.SetX(r.marginLeft + 15)
		}
		r.pdf.MultiCell(0, r.lineHeight, line, "", "", false)
	}

	// Reset font
	r.pdf.SetFont("Arial", "", r.fontSize)
	r.pdf.Ln(r.lineHeight)
	return nil
}

func (r *PDFRenderer) renderTable(node *nodes.TableNode) error {
	headers := node.Headers()
	rows := node.Rows()

	// Calculate column widths
	colWidth := 170.0 / float64(len(headers))

	// Draw headers
	r.pdf.SetFont("Arial", "B", r.fontSize)
	r.pdf.SetFillColor(240, 240, 240)

	for _, header := range headers {
		r.pdf.CellFormat(colWidth, r.lineHeight, header, "1", 0, "", true, 0, "")
	}
	r.pdf.Ln(-1)

	// Draw rows
	r.pdf.SetFont("Arial", "", r.fontSize)
	r.pdf.SetFillColor(255, 255, 255)

	for _, row := range rows {
		for _, cell := range row {
			r.pdf.CellFormat(colWidth, r.lineHeight, cell, "1", 0, "", false, 0, "")
		}
		r.pdf.Ln(-1)
	}

	r.pdf.Ln(r.lineHeight)
	return nil
}

func (r *PDFRenderer) renderDirective(node *nodes.DirectiveNode) error {
	// Handle special directives
	switch node.Name() {
	case "image":
		// Handle image directive
		if len(node.Arguments()) > 0 {
			imagePath := node.Arguments()[0]
			r.pdf.Image(imagePath, r.pdf.GetX(), r.pdf.GetY(), 0, 0, false, "", 0, "")
			r.pdf.Ln(r.lineHeight)
		}
	case "note", "warning", "important":
		// Handle admonitions
		r.pdf.SetFillColor(245, 245, 245)
		startY := r.pdf.GetY()
		r.pdf.MultiCell(0, r.lineHeight, node.RawContent(), "", "", true)
		r.pdf.SetY(startY + r.lineHeight)
	}

	return r.renderChildren(node)
}

func (r *PDFRenderer) renderChildren(node nodes.Node) error {
	return r.Render(node.Children())
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/comment.go`:

```````go
package nodes

import "fmt"

// CommentNode represents RST comments starting with ..
type CommentNode struct {
	*BaseNode
}

func NewCommentNode(content string) *CommentNode {
	node := &CommentNode{
		BaseNode: NewBaseNode(NodeComment),
	}
	node.SetContent(content)
	return node
}

func (n *CommentNode) String() string {
	return fmt.Sprintf("Comment: %s", n.Content())
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/paragraph.go`:

```````go
package nodes

import "fmt"

// ParagraphNode represents a text paragraph
type ParagraphNode struct {
	*BaseNode
}

// NewParagraphNode creates a new ParagraphNode with the given content
func NewParagraphNode(content string) *ParagraphNode {
	node := &ParagraphNode{
		BaseNode: NewBaseNode(NodeParagraph),
	}
	node.SetContent(content)
	return node
}

// String representation for debugging
func (n *ParagraphNode) String() string {
	return fmt.Sprintf("Paragraph: %s", n.Content())
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/title.go`:

```````go
package nodes

import "fmt"

type TitleNode struct {
	*BaseNode
	level int
}

func NewTitleNode(content string, level int) *TitleNode {
	node := &TitleNode{
		BaseNode: NewBaseNode(NodeTitle),
		level:    level,
	}
	node.SetContent(content)
	return node
}

func (n *TitleNode) Level() int {
	return n.level
}

func (n *TitleNode) String() string {
	return fmt.Sprintf("Title[%d]: %s", n.Level(), n.Content())
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/code.go`:

```````go
package nodes

import "fmt"

// CodeNode represents a code block
type CodeNode struct {
	*BaseNode
	language    string
	lineNumbers bool
}

// NewCodeNode creates a new CodeNode with the given language and content
func NewCodeNode(language, content string, lineNumbers bool) *CodeNode {
	node := &CodeNode{
		BaseNode:    NewBaseNode(NodeCode),
		language:    language,
		lineNumbers: lineNumbers,
	}
	node.SetContent(content)
	return node
}

// Language returns the language of the code block
func (n *CodeNode) Language() string { return n.language }

// LineNumbers returns the line numbers flag of the code block
func (n *CodeNode) LineNumbers() bool { return n.lineNumbers }

// String representation for debugging
func (n *CodeNode) String() string {
	return fmt.Sprintf("Code[%s]: %d bytes", n.language, len(n.Content()))
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/directive.go`:

```````go
package nodes

import "fmt"

// DirectiveNode represents an RST directive
type DirectiveNode struct {
	*BaseNode
	name       string
	arguments  []string
	rawContent string
}

// NewDirectiveNode creates a new DirectiveNode with the given name and arguments
func NewDirectiveNode(name string, args []string) *DirectiveNode {
	node := &DirectiveNode{
		BaseNode:   NewBaseNode(NodeDirective),
		name:       name,
		arguments:  args,
		rawContent: "",
	}
	return node
}

// Name returns the name of the directive
func (n *DirectiveNode) Name() string { return n.name }

// Arguments returns the arguments of the directive
func (n *DirectiveNode) Arguments() []string { return n.arguments }

// RawContent returns the raw content of the directive
func (n *DirectiveNode) RawContent() string { return n.rawContent }

// SetRawContent sets the raw content of the directive
func (n *DirectiveNode) SetRawContent(content string) {
	n.rawContent = content
}

// String representation for debugging
func (n *DirectiveNode) String() string {
	return fmt.Sprintf("Directive[%s]: %s", n.name, n.Content())
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/lineblock.go`:

```````go
package nodes

import "fmt"

// LineBlockNode represents poetry-style line blocks that preserve line breaks
type LineBlockNode struct {
	*BaseNode
	lines []string
}

func NewLineBlockNode(lines []string) *LineBlockNode {
	node := &LineBlockNode{
		BaseNode: NewBaseNode(NodeLineBlock),
		lines:    lines,
	}
	return node
}

func (n *LineBlockNode) Lines() []string {
	return n.lines
}

func (n *LineBlockNode) String() string {
	return fmt.Sprintf("LineBlock: %d lines", len(n.lines))
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/doc.md`:

```````md
# nodes
--
    import "github.com/go-i2p/go-rst/pkg/nodes"


## Usage

#### func  GetIndentedContent

```go
func GetIndentedContent(node Node) string
```
GetIndentedContent Utility function to get node content with proper indentation

#### type BaseNode

```go
type BaseNode struct {
}
```

BaseNode provides the basic implementation of the Node interface that other node
types can embed

#### func  NewBaseNode

```go
func NewBaseNode(nodeType NodeType) *BaseNode
```
NewBaseNode creates a new BaseNode with the specified node type

Parameters:

    - nodeType: The type of node to create

Returns:

    - *BaseNode: A new base node instance

#### func (*BaseNode) AddChild

```go
func (n *BaseNode) AddChild(child Node)
```
AddChild adds a child node to this node

#### func (*BaseNode) Children

```go
func (n *BaseNode) Children() []Node
```
Children returns the node's child nodes

#### func (*BaseNode) Content

```go
func (n *BaseNode) Content() string
```
Content returns the textual content of the node

#### func (*BaseNode) Level

```go
func (n *BaseNode) Level() int
```
Level returns the nesting level of the node

#### func (*BaseNode) SetContent

```go
func (n *BaseNode) SetContent(content string)
```
SetContent sets the node's textual content

#### func (*BaseNode) SetLevel

```go
func (n *BaseNode) SetLevel(level int)
```
SetLevel sets the nesting level of the node

#### func (*BaseNode) Type

```go
func (n *BaseNode) Type() NodeType
```
Type returns the NodeType

#### type BlockQuoteNode

```go
type BlockQuoteNode struct {
	*BaseNode
}
```

BlockQuoteNode represents an indented block quote

#### func  NewBlockQuoteNode

```go
func NewBlockQuoteNode(content, attribution string) *BlockQuoteNode
```
NewBlockQuoteNode creates a new BlockQuoteNode with the given content

#### func (*BlockQuoteNode) Attribution

```go
func (n *BlockQuoteNode) Attribution() string
```
Attribution returns the quote attribution if any

#### func (*BlockQuoteNode) String

```go
func (n *BlockQuoteNode) String() string
```
String representation for debugging

#### type CodeNode

```go
type CodeNode struct {
	*BaseNode
}
```

CodeNode represents a code block

#### func  NewCodeNode

```go
func NewCodeNode(language, content string, lineNumbers bool) *CodeNode
```
NewCodeNode creates a new CodeNode with the given language and content

#### func (*CodeNode) Language

```go
func (n *CodeNode) Language() string
```
Language returns the language of the code block

#### func (*CodeNode) LineNumbers

```go
func (n *CodeNode) LineNumbers() bool
```
LineNumbers returns the line numbers flag of the code block

#### func (*CodeNode) String

```go
func (n *CodeNode) String() string
```
String representation for debugging

#### type CommentNode

```go
type CommentNode struct {
	*BaseNode
}
```

CommentNode represents RST comments starting with ..

#### func  NewCommentNode

```go
func NewCommentNode(content string) *CommentNode
```

#### func (*CommentNode) String

```go
func (n *CommentNode) String() string
```

#### type DirectiveNode

```go
type DirectiveNode struct {
	*BaseNode
}
```

DirectiveNode represents an RST directive

#### func  NewDirectiveNode

```go
func NewDirectiveNode(name string, args []string) *DirectiveNode
```
NewDirectiveNode creates a new DirectiveNode with the given name and arguments

#### func (*DirectiveNode) Arguments

```go
func (n *DirectiveNode) Arguments() []string
```
Arguments returns the arguments of the directive

#### func (*DirectiveNode) Name

```go
func (n *DirectiveNode) Name() string
```
Name returns the name of the directive

#### func (*DirectiveNode) RawContent

```go
func (n *DirectiveNode) RawContent() string
```
RawContent returns the raw content of the directive

#### func (*DirectiveNode) SetRawContent

```go
func (n *DirectiveNode) SetRawContent(content string)
```
SetRawContent sets the raw content of the directive

#### func (*DirectiveNode) String

```go
func (n *DirectiveNode) String() string
```
String representation for debugging

#### type DoctestNode

```go
type DoctestNode struct {
	*BaseNode
}
```

DoctestNode represents a doctest block with expected output

#### func  NewDoctestNode

```go
func NewDoctestNode(command, expected string) *DoctestNode
```

#### func (*DoctestNode) Command

```go
func (n *DoctestNode) Command() string
```

#### func (*DoctestNode) Expected

```go
func (n *DoctestNode) Expected() string
```

#### func (*DoctestNode) String

```go
func (n *DoctestNode) String() string
```

#### type EmphasisNode

```go
type EmphasisNode struct {
	*BaseNode
}
```

EmphasisNode represents emphasized text (italic)

#### func  NewEmphasisNode

```go
func NewEmphasisNode(content string) *EmphasisNode
```
NewEmphasisNode creates a new EmphasisNode with the given content

#### type HeadingNode

```go
type HeadingNode struct {
	*BaseNode
}
```

HeadingNode represents a section heading in RST

#### func  NewHeadingNode

```go
func NewHeadingNode(content string, level int) *HeadingNode
```
NewHeadingNode creates a new HeadingNode with the given content and level

#### func (*HeadingNode) String

```go
func (n *HeadingNode) String() string
```
String representations for debugging

#### type LineBlockNode

```go
type LineBlockNode struct {
	*BaseNode
}
```

LineBlockNode represents poetry-style line blocks that preserve line breaks

#### func  NewLineBlockNode

```go
func NewLineBlockNode(lines []string) *LineBlockNode
```

#### func (*LineBlockNode) Lines

```go
func (n *LineBlockNode) Lines() []string
```

#### func (*LineBlockNode) String

```go
func (n *LineBlockNode) String() string
```

#### type LinkNode

```go
type LinkNode struct {
	*BaseNode
}
```

LinkNode represents a hyperlink

#### func  NewLinkNode

```go
func NewLinkNode(text, url, title string) *LinkNode
```
NewLinkNode creates a new LinkNode with the given text, URL, and title

#### func (*LinkNode) String

```go
func (n *LinkNode) String() string
```
String representation for debugging

#### func (*LinkNode) Title

```go
func (n *LinkNode) Title() string
```
Title returns the URL of the link

#### func (*LinkNode) URL

```go
func (n *LinkNode) URL() string
```
URL returns the URL of the link

#### type ListItemNode

```go
type ListItemNode struct {
	*BaseNode
}
```

ListItemNode represents an individual list item

#### func  NewListItemNode

```go
func NewListItemNode(content string) *ListItemNode
```
NewListItemNode creates a new ListItemNode with the given content

#### type ListNode

```go
type ListNode struct {
	*BaseNode
}
```

ListNode represents an ordered or unordered list

#### func  NewListNode

```go
func NewListNode(ordered bool) *ListNode
```
NewListNode creates a new ListNode with the given ordered flag

#### func (*ListNode) IsOrdered

```go
func (n *ListNode) IsOrdered() bool
```
IsOrdered returns true if the list is ordered, false otherwise

#### func (*ListNode) String

```go
func (n *ListNode) String() string
```
String representation for debugging

#### type MetaNode

```go
type MetaNode struct {
	*BaseNode
}
```

MetaNode represents metadata information

#### func  NewMetaNode

```go
func NewMetaNode(key, value string) *MetaNode
```
NewMetaNode creates a new MetaNode with the given key and value

#### func (*MetaNode) Key

```go
func (n *MetaNode) Key() string
```
Key returns the key of the metadata

#### type Node

```go
type Node interface {
	// Type returns the NodeType of this node
	Type() NodeType
	// Content returns the textual content of the node
	Content() string
	// SetContent sets the node's textual content
	SetContent(string)
	// Level returns the nesting level of the node
	Level() int
	// SetLevel sets the nesting level of the node
	SetLevel(int)
	// Children returns the node's child nodes
	Children() []Node
	// AddChild adds a child node to this node
	AddChild(Node)
}
```

Node interface defines the common behavior for all RST document nodes

#### type NodeType

```go
type NodeType int
```

NodeType represents the type of a node in the RST document structure

```go
const (
	NodeHeading    NodeType = iota // Represents a section heading
	NodeParagraph                  // Represents a text paragraph
	NodeList                       // Represents an ordered or unordered list
	NodeListItem                   // Represents an item within a list
	NodeLink                       // Represents a hyperlink
	NodeEmphasis                   // Represents emphasized (italic) text
	NodeStrong                     // Represents strong (bold) text
	NodeMeta                       // Represents metadata information
	NodeDirective                  // Represents an RST directive
	NodeCode                       // Represents a code block
	NodeTable                      // Represents a table structure
	NodeBlockQuote                 // Represents a block quote
	NodeDoctest                    // Represents a doctest block
	NodeLineBlock                  // Represents a line block
	NodeComment                    // Represents a comment
	NodeTitle                      // Represents a document title
	NodeSubtitle                   // Represents a document subtitle
	NodeTransition
)
```
Node type constants define the possible types of nodes in the RST document tree

#### type ParagraphNode

```go
type ParagraphNode struct {
	*BaseNode
}
```

ParagraphNode represents a text paragraph

#### func  NewParagraphNode

```go
func NewParagraphNode(content string) *ParagraphNode
```
NewParagraphNode creates a new ParagraphNode with the given content

#### func (*ParagraphNode) String

```go
func (n *ParagraphNode) String() string
```
String representation for debugging

#### type StrongNode

```go
type StrongNode struct {
	*BaseNode
}
```

StrongNode represents strong text (bold)

#### func  NewStrongNode

```go
func NewStrongNode(content string) *StrongNode
```
NewStrongNode creates a new StrongNode with the given content

#### type SubtitleNode

```go
type SubtitleNode struct {
	*BaseNode
}
```


#### func  NewSubtitleNode

```go
func NewSubtitleNode(content string) *SubtitleNode
```

#### func (*SubtitleNode) String

```go
func (n *SubtitleNode) String() string
```

#### type TableNode

```go
type TableNode struct {
	*BaseNode
}
```

TableNode represents a table structure

#### func  NewTableNode

```go
func NewTableNode() *TableNode
```
NewTableNode creates a new TableNode

#### func (*TableNode) AddRow

```go
func (n *TableNode) AddRow(row []string)
```
AddRow adds a row to the table

#### func (*TableNode) Headers

```go
func (n *TableNode) Headers() []string
```
Headers returns the headers of the table

#### func (*TableNode) Rows

```go
func (n *TableNode) Rows() [][]string
```
Rows returns the rows of the table

#### func (*TableNode) SetHeaders

```go
func (n *TableNode) SetHeaders(headers []string)
```
SetHeaders sets the headers of the table

#### func (*TableNode) String

```go
func (n *TableNode) String() string
```
String representation for debugging

#### type TitleNode

```go
type TitleNode struct {
	*BaseNode
}
```


#### func  NewTitleNode

```go
func NewTitleNode(content string, level int) *TitleNode
```

#### func (*TitleNode) Level

```go
func (n *TitleNode) Level() int
```

#### func (*TitleNode) String

```go
func (n *TitleNode) String() string
```

#### type TransitionNode

```go
type TransitionNode struct {
	*BaseNode
}
```


#### func  NewTransitionNode

```go
func NewTransitionNode(char rune) *TransitionNode
```

#### func (*TransitionNode) Character

```go
func (n *TransitionNode) Character() rune
```

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/blockquote.go`:

```````go
package nodes

import "fmt"

// BlockQuoteNode represents an indented block quote
type BlockQuoteNode struct {
	*BaseNode
	attribution string
}

func (n *BlockQuoteNode) AppendContent(content string) {
	n.SetContent(n.Content() + "\n" + content)
}

// NewBlockQuoteNode creates a new BlockQuoteNode with the given content
func NewBlockQuoteNode(content, attribution string) *BlockQuoteNode {
	node := &BlockQuoteNode{
		BaseNode:    NewBaseNode(NodeBlockQuote),
		attribution: attribution,
	}
	node.SetContent(content)
	return node
}

// Attribution returns the quote attribution if any
func (n *BlockQuoteNode) Attribution() string {
	return n.attribution
}

// String representation for debugging
func (n *BlockQuoteNode) String() string {
	if n.attribution != "" {
		return fmt.Sprintf("BlockQuote: %s -- %s", n.Content(), n.attribution)
	}
	return fmt.Sprintf("BlockQuote: %s", n.Content())
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/list.go`:

```````go
package nodes

import "fmt"

// ListNode represents an ordered or unordered list
type ListNode struct {
	*BaseNode
	ordered bool
	indent  int
}

func (n *ListNode) Indent() int {
	return n.indent
}

func (n *ListNode) SetIndent(indent int) {
	n.indent = indent
}

// AppendChild adds a list item node as a child to this list node
func (n *ListNode) AppendChild(listItem *ListItemNode) {
	n.BaseNode.AddChild(listItem)
}

// NewListNode creates a new ListNode with the given ordered flag
func NewListNode(ordered bool) *ListNode {
	node := &ListNode{
		BaseNode: NewBaseNode(NodeList),
		ordered:  ordered,
	}
	return node
}

// IsOrdered returns true if the list is ordered, false otherwise
func (n *ListNode) IsOrdered() bool {
	return n.ordered
}

// ListItemNode represents an individual list item
type ListItemNode struct {
	*BaseNode
}

// NewListItemNode creates a new ListItemNode with the given content
func NewListItemNode(content string) *ListItemNode {
	node := &ListItemNode{
		BaseNode: NewBaseNode(NodeListItem),
	}
	node.SetContent(content)
	return node
}

// String representation for debugging
func (n *ListNode) String() string {
	listType := "Unordered"
	if n.ordered {
		listType = "Ordered"
	}
	return fmt.Sprintf("%s List with %d items", listType, len(n.Children()))
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/strong.go`:

```````go
package nodes

// StrongNode represents strong text (bold)
type StrongNode struct {
	*BaseNode
}

// NewStrongNode creates a new StrongNode with the given content
func NewStrongNode(content string) *StrongNode {
	node := &StrongNode{
		BaseNode: NewBaseNode(NodeStrong),
	}
	node.SetContent(content)
	return node
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/meta.go`:

```````go
package nodes

// MetaNode represents metadata information
type MetaNode struct {
	*BaseNode
	key string
}

// NewMetaNode creates a new MetaNode with the given key and value
func NewMetaNode(key, value string) *MetaNode {
	node := &MetaNode{
		BaseNode: NewBaseNode(NodeMeta),
		key:      key,
	}
	node.SetContent(value)
	return node
}

// Key returns the key of the metadata
func (n *MetaNode) Key() string { return n.key }

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/doctest.go`:

```````go
package nodes

// DoctestNode represents a Python doctest block with code and expected output.
type DoctestNode struct {
	BaseNode
	code           string
	expectedOutput string
}

func (n *DoctestNode) Expected() string {
	return n.expectedOutput
}

func (n *DoctestNode) Command() string {
	return n.code
}

// NewDoctestNode creates a new doctest node.
func NewDoctestNode() *DoctestNode {
	return &DoctestNode{
		BaseNode: BaseNode{
			nodeType: NodeDoctest,
		},
		code:           "",
		expectedOutput: "",
	}
}

// SetCode sets the code content of the doctest.
func (n *DoctestNode) SetCode(code string) {
	n.code = code
}

// Code returns the code content of the doctest.
func (n *DoctestNode) Code() string {
	return n.code
}

// SetExpectedOutput sets the expected output of the doctest.
func (n *DoctestNode) SetExpectedOutput(output string) {
	n.expectedOutput = output
}

// ExpectedOutput returns the expected output of the doctest.
func (n *DoctestNode) ExpectedOutput() string {
	return n.expectedOutput
}

// Content returns the code of this node.
func (n *DoctestNode) Content() string {
	return n.code
}

// SetContent sets the code of this node.
func (n *DoctestNode) SetContent(content string) {
	n.code = content
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/heading.go`:

```````go
package nodes

import "fmt"

// HeadingNode represents a section heading in RST
type HeadingNode struct {
	*BaseNode
}

// NewHeadingNode creates a new HeadingNode with the given content and level
func NewHeadingNode(content string, level int) *HeadingNode {
	node := &HeadingNode{
		BaseNode: NewBaseNode(NodeHeading),
	}
	node.SetContent(content)
	node.SetLevel(level)
	return node
}

// String representations for debugging
func (n *HeadingNode) String() string {
	return fmt.Sprintf("Heading[%d]: %s", n.Level(), n.Content())
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/subtitle.go`:

```````go
package nodes

import "fmt"

type SubtitleNode struct {
	*BaseNode
}

func NewSubtitleNode(content string) *SubtitleNode {
	node := &SubtitleNode{
		BaseNode: NewBaseNode(NodeSubtitle),
	}
	node.SetContent(content)
	return node
}

func (n *SubtitleNode) String() string {
	return fmt.Sprintf("Subtitle: %s", n.Content())
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/em.go`:

```````go
package nodes

// EmphasisNode represents emphasized text (italic)
type EmphasisNode struct {
	*BaseNode
}

// NewEmphasisNode creates a new EmphasisNode with the given content
func NewEmphasisNode(content string) *EmphasisNode {
	node := &EmphasisNode{
		BaseNode: NewBaseNode(NodeEmphasis),
	}
	node.SetContent(content)
	return node
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/types.go`:

```````go
// pkg/nodes/base.go

package nodes

// NodeType represents the type of a node in the RST document structure
type NodeType int

// Node type constants define the possible types of nodes in the RST document tree
const (
	NodeHeading    NodeType = iota // Represents a section heading
	NodeParagraph                  // Represents a text paragraph
	NodeList                       // Represents an ordered or unordered list
	NodeListItem                   // Represents an item within a list
	NodeLink                       // Represents a hyperlink
	NodeEmphasis                   // Represents emphasized (italic) text
	NodeStrong                     // Represents strong (bold) text
	NodeMeta                       // Represents metadata information
	NodeDirective                  // Represents an RST directive
	NodeCode                       // Represents a code block
	NodeTable                      // Represents a table structure
	NodeBlockQuote                 // Represents a block quote
	NodeDoctest                    // Represents a doctest block
	NodeLineBlock                  // Represents a line block
	NodeComment                    // Represents a comment
	NodeTitle                      // Represents a document title
	NodeSubtitle                   // Represents a document subtitle
	NodeTransition
)

// Node interface defines the common behavior for all RST document nodes
type Node interface {
	// Type returns the NodeType of this node
	Type() NodeType
	// Content returns the textual content of the node
	Content() string
	// SetContent sets the node's textual content
	SetContent(string)
	// Level returns the nesting level of the node
	Level() int
	// SetLevel sets the nesting level of the node
	SetLevel(int)
	// Children returns the node's child nodes
	Children() []Node
	// AddChild adds a child node to this node
	AddChild(Node)
}

// BaseNode provides the basic implementation of the Node interface
// that other node types can embed
type BaseNode struct {
	nodeType NodeType
	content  string
	level    int
	children []Node
}

// NewBaseNode creates a new BaseNode with the specified node type
//
// Parameters:
//   - nodeType: The type of node to create
//
// Returns:
//   - *BaseNode: A new base node instance
func NewBaseNode(nodeType NodeType) *BaseNode {
	return &BaseNode{
		nodeType: nodeType,
		children: make([]Node, 0),
	}
}

// Type returns the NodeType
func (n *BaseNode) Type() NodeType { return n.nodeType }

// Content returns the textual content of the node
func (n *BaseNode) Content() string { return n.content }

// SetContent sets the node's textual content
func (n *BaseNode) SetContent(content string) {
	n.content = content
}

// Level returns the nesting level of the node
func (n *BaseNode) Level() int { return n.level }

// SetLevel sets the nesting level of the node
func (n *BaseNode) SetLevel(level int) {
	n.level = level
}

// Children returns the node's child nodes
func (n *BaseNode) Children() []Node {
	return n.children
}

// AddChild adds a child node to this node
func (n *BaseNode) AddChild(child Node) {
	n.children = append(n.children, child)
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/table.go`:

```````go
package nodes

import "fmt"

// TableNode represents a table structure
type TableNode struct {
	*BaseNode
	headers []string
	rows    [][]string
}

// NewTableNode creates a new TableNode
func NewTableNode() *TableNode {
	return &TableNode{
		BaseNode: NewBaseNode(NodeTable),
		headers:  make([]string, 0),
		rows:     make([][]string, 0),
	}
}

// SetHeaders sets the headers of the table
func (n *TableNode) SetHeaders(headers []string) {
	n.headers = headers
}

// AddRow adds a row to the table
func (n *TableNode) AddRow(row []string) {
	n.rows = append(n.rows, row)
}

// Headers returns the headers of the table
func (n *TableNode) Headers() []string { return n.headers }

// Rows returns the rows of the table
func (n *TableNode) Rows() [][]string { return n.rows }

// String representation for debugging
func (n *TableNode) String() string {
	return fmt.Sprintf("Table: %d columns x %d rows", len(n.headers), len(n.rows))
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/link.go`:

```````go
package nodes

import "fmt"

// LinkNode represents a hyperlink
type LinkNode struct {
	*BaseNode
	url   string
	title string
}

// NewLinkNode creates a new LinkNode with the given text, URL, and title
func NewLinkNode(text, url, title string) *LinkNode {
	node := &LinkNode{
		BaseNode: NewBaseNode(NodeLink),
		url:      url,
		title:    title,
	}
	node.SetContent(text)
	return node
}

// URL returns the URL of the link
func (n *LinkNode) URL() string { return n.url }

// Title returns the URL of the link
func (n *LinkNode) Title() string { return n.title }

// String representation for debugging
func (n *LinkNode) String() string {
	return fmt.Sprintf("Link[%s](%s)", n.Content(), n.url)
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/transition.go`:

```````go
package nodes

import "fmt"

type TransitionNode struct {
	*BaseNode
	character rune
}

func NewTransitionNode(char rune) *TransitionNode {
	return &TransitionNode{
		BaseNode:  NewBaseNode(NodeTransition),
		character: char,
	}
}

func (n *TransitionNode) Character() rune {
	return n.character
}

func (n *TransitionNode) String() string {
	return fmt.Sprintf("Transition: %c", n.character)
}

```````

`/home/idk/go/src/github.com/go-i2p/go-rst/pkg/nodes/extra_util.go`:

```````go
package nodes

import "strings"

// GetIndentedContent Utility function to get node content with proper indentation
func GetIndentedContent(node Node) string {
	content := node.Content()
	if node.Level() > 0 {
		indent := strings.Repeat("    ", node.Level())
		lines := strings.Split(content, "\n")
		for i, line := range lines {
			lines[i] = indent + line
		}
		content = strings.Join(lines, "\n")
	}
	return content
}

```````