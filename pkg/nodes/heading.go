package nodes

import "fmt"

// HeadingNode represents a section heading in RST
type HeadingNode struct {
	*BaseNode
}

// NewHeadingNode creates a new HeadingNode with the given content and level
func NewHeadingNode(content string, level int) *HeadingNode {
	node := &HeadingNode{
		BaseNode: NewBaseNode(NodeHeading),
	}
	node.SetContent(content)
	node.SetLevel(level)
	return node
}

// String representations for debugging
func (n *HeadingNode) String() string {
	return fmt.Sprintf("Heading[%d]: %s", n.Level(), n.Content())
}
