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
