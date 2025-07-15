package parser

import (
	"github.com/go-i2p/go-rst/pkg/nodes"
)

// processEmphasis handles the parsing of emphasized (italic) text
func (p *Parser) processEmphasis(content string) *nodes.EmphasisNode {
	// If translator is available, translate the content
	if p.translator != nil {
		content = p.translator.Translate(content)
	}

	// Create a new emphasis node with the content
	return nodes.NewEmphasisNode(content)
}
