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
