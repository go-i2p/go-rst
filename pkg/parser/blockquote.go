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
