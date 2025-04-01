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
	}

	return currentNode
}
