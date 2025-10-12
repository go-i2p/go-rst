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
// This function is designed to never panic on user input, instead degrading gracefully
// by returning partial results or empty slices for malformed content.
func (p *Parser) Parse(content string) []nodes.Node {
	// Add panic recovery at top level to ensure we never crash on bad input
	defer func() {
		if r := recover(); r != nil {
			// Log the panic but return what we have so far
			// In production, this could log to a proper logger
		}
	}()

	// Validate input - empty content returns empty slice
	if content == "" {
		return []nodes.Node{}
	}

	scanner := bufio.NewScanner(strings.NewReader(content))
	var currentNode nodes.Node
	var prevToken Token
	p.nodes = make([]nodes.Node, 0) // Clear existing nodes
	p.context.Reset()               // Reset parser context to prevent state pollution between parses

	for scanner.Scan() {
		line := scanner.Text()
		token := p.lexer.Tokenize(line)

		if newNode := p.processToken(token, prevToken, currentNode, line); newNode != nil {
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

func (p *Parser) processToken(token, prevToken Token, currentNode nodes.Node, originalLine string) nodes.Node {
	switch token.Type {
	case TokenBulletList, TokenEnumList, TokenBlockQuote, TokenComment:
		return p.processListAndBlockTokens(token, currentNode)
	case TokenTransBlock:
		return p.processTranslationBlock(token)
	case TokenHeadingUnderline:
		return p.processHeadingToken(token, prevToken)
	case TokenMeta, TokenCodeBlock, TokenDirective:
		return p.processStructuralTokens(token)
	case TokenEmphasis, TokenStrong:
		return p.processFormattingTokens(token)
	case TokenLineBlock:
		return p.processLineBlockToken(token, currentNode)
	case TokenFootnote, TokenDefinitionList, TokenFieldList:
		return p.processSpecialElementTokens(token, currentNode)
	case TokenText:
		return p.processTextToken(token, currentNode, originalLine)
	case TokenTransition:
		return p.processTransitionToken(token)
	}
	return currentNode
}

// processListAndBlockTokens handles list items, block quotes, and comments.
func (p *Parser) processListAndBlockTokens(token Token, currentNode nodes.Node) nodes.Node {
	switch token.Type {
	case TokenBulletList:
		return p.processListItem(token.Content, token.Args, false, currentNode)
	case TokenEnumList:
		return p.processListItem(token.Content, token.Args, true, currentNode)
	case TokenBlockQuote:
		return p.processBlockQuote(token.Content, token.Args, currentNode)
	case TokenComment:
		return nodes.NewCommentNode(token.Content)
	}
	return currentNode
}

// processTranslationBlock handles translation block tokens with optional translation.
func (p *Parser) processTranslationBlock(token Token) nodes.Node {
	content := strings.TrimSpace(token.Content)
	if p.translator != nil {
		translatedContent := p.translator.Translate(content)
		return nodes.NewParagraphNode(translatedContent)
	}
	return nodes.NewParagraphNode(content)
}

// processHeadingToken handles heading underline tokens.
func (p *Parser) processHeadingToken(token, prevToken Token) nodes.Node {
	if prevToken.Type == TokenText {
		return p.processHeading(prevToken.Content, token.Content)
	}
	return nil
}

// processStructuralTokens handles meta, code block, and directive tokens.
func (p *Parser) processStructuralTokens(token Token) nodes.Node {
	switch token.Type {
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
	}
	return nil
}

// processFormattingTokens handles emphasis and strong formatting tokens.
func (p *Parser) processFormattingTokens(token Token) nodes.Node {
	switch token.Type {
	case TokenEmphasis:
		return p.processEmphasis(token.Content)
	case TokenStrong:
		return p.processStrong(token.Content)
	}
	return nil
}

// processLineBlockToken handles line block tokens and manages existing line blocks.
func (p *Parser) processLineBlockToken(token Token, currentNode nodes.Node) nodes.Node {
	if lineBlock, ok := currentNode.(*nodes.LineBlockNode); ok {
		lines := lineBlock.Lines()
		lines = append(lines, token.Content)
		return nodes.NewLineBlockNode(lines)
	}
	return nodes.NewLineBlockNode([]string{token.Content})
}

// processTextToken handles text tokens based on current parser context.
func (p *Parser) processTextToken(token Token, currentNode nodes.Node, originalLine string) nodes.Node {
	if p.context.inCodeBlock {
		return p.processCodeBlock(originalLine, currentNode)
	}
	if p.context.inMeta {
		return p.processMetaContent(token.Content, currentNode)
	}
	if p.context.inDirective {
		return p.processDirectiveContent(token.Content, currentNode)
	}
	return p.processParagraph(token.Content, currentNode)
}

// processTransitionToken handles transition tokens with optional content.
func (p *Parser) processTransitionToken(token Token) nodes.Node {
	if len(token.Content) > 0 {
		return p.processTransition(token.Content)
	}
	return nodes.NewTransitionNode('-')
}

// processSpecialElementTokens handles footnotes, definition lists, and field lists.
func (p *Parser) processSpecialElementTokens(token Token, currentNode nodes.Node) nodes.Node {
	switch token.Type {
	case TokenFootnote:
		return p.processFootnote(token.Content, token.Args)
	case TokenDefinitionList:
		return p.processDefinitionList(token.Content, currentNode)
	case TokenFieldList:
		return p.processFieldList(token.Content, token.Args, currentNode)
	}
	return currentNode
}

// processFootnote creates a footnote node from the token content and label.
func (p *Parser) processFootnote(content string, args []string) nodes.Node {
	label := ""
	autoNumber := false

	if len(args) > 0 {
		label = args[0]
		// Check if it's an auto-numbered footnote (#) or symbol footnote (*)
		autoNumber = (label == "#" || label == "*")
	}

	return nodes.NewFootnoteNode(label, content, autoNumber)
}

// processDefinitionList handles definition list processing.
// Note: Definition lists require multiline processing which is complex for line-by-line tokenization.
// For now, we treat them as regular paragraphs until multiline support is added.
func (p *Parser) processDefinitionList(content string, currentNode nodes.Node) nodes.Node {
	// Simple implementation: treat as paragraph for now
	// TODO: Implement proper multiline definition list parsing
	return nodes.NewParagraphNode(content)
}

// processFieldList creates or updates a field list node with metadata.
func (p *Parser) processFieldList(content string, args []string, currentNode nodes.Node) nodes.Node {
	fieldName := ""
	if len(args) > 0 {
		fieldName = args[0]
	}

	// Check if we're continuing an existing field list
	if fieldList, ok := currentNode.(*nodes.FieldListNode); ok {
		fieldList.AddField(fieldName, content)
		return fieldList
	}

	// Create new field list node
	fieldList := nodes.NewFieldListNode()
	fieldList.AddField(fieldName, content)
	return fieldList
}
