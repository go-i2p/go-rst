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
