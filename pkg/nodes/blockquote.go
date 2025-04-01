package nodes

import "fmt"

// BlockQuoteNode represents an indented block quote
type BlockQuoteNode struct {
	*BaseNode
	attribution string
}

// NewBlockQuoteNode creates a new BlockQuoteNode with the given content
func NewBlockQuoteNode(content, attribution string) *BlockQuoteNode {
	node := &BlockQuoteNode{
		BaseNode:    NewBaseNode(NodeBlockQuote),
		attribution: attribution,
	}
	node.SetContent(content)
	return node
}

// Attribution returns the quote attribution if any
func (n *BlockQuoteNode) Attribution() string {
	return n.attribution
}

// String representation for debugging
func (n *BlockQuoteNode) String() string {
	if n.attribution != "" {
		return fmt.Sprintf("BlockQuote: %s -- %s", n.Content(), n.attribution)
	}
	return fmt.Sprintf("BlockQuote: %s", n.Content())
}
