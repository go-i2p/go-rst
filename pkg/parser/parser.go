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
		content := strings.TrimSpace(token.Content)
		if p.translator != nil {
			translatedContent := p.translator.Translate(content)
			return nodes.NewParagraphNode(translatedContent)
		}
		// If no translator is available, return the original content
		return nodes.NewParagraphNode(content)
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

	case TokenEmphasis:
		// Process the emphasized text
		return p.processEmphasis(token.Content)

	case TokenStrong:
		// Process the strong (bold) text
		return p.processStrong(token.Content)

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
			return p.processCodeBlock(originalLine, currentNode) // Use original line to preserve indentation
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
