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
