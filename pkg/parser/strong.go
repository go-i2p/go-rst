package parser

import (
	"github.com/go-i2p/go-rst/pkg/nodes"
)

// processStrong handles the parsing of strong (bold) text
func (p *Parser) processStrong(content string) *nodes.StrongNode {
	// If translator is available, translate the content
	if p.translator != nil {
		content = p.translator.Translate(content)
	}

	// Create a new strong node with the content
	return nodes.NewStrongNode(content)
}
