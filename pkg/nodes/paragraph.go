package nodes

import "fmt"

// ParagraphNode represents a text paragraph
type ParagraphNode struct {
	*BaseNode
}

// NewParagraphNode creates a new ParagraphNode with the given content
func NewParagraphNode(content string) *ParagraphNode {
	node := &ParagraphNode{
		BaseNode: NewBaseNode(NodeParagraph),
	}
	node.SetContent(content)
	return node
}

// String representation for debugging
func (n *ParagraphNode) String() string {
	return fmt.Sprintf("Paragraph: %s", n.Content())
}
