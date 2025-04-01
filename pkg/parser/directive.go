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
