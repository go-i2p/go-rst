package nodes

import "fmt"

// LineBlockNode represents poetry-style line blocks that preserve line breaks
type LineBlockNode struct {
	*BaseNode
	lines []string
}

func NewLineBlockNode(lines []string) *LineBlockNode {
	node := &LineBlockNode{
		BaseNode: NewBaseNode(NodeLineBlock),
		lines:    lines,
	}
	return node
}

func (n *LineBlockNode) Lines() []string {
	return n.lines
}

func (n *LineBlockNode) String() string {
	return fmt.Sprintf("LineBlock: %d lines", len(n.lines))
}
