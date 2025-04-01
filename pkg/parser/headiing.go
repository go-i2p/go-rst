package parser

import (
	"strings"

	"github.com/go-i2p/go-rst/pkg/nodes"
)

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
