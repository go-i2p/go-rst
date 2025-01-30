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
