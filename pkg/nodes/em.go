package nodes

// EmphasisNode represents emphasized text (italic)
type EmphasisNode struct {
	*BaseNode
}

// NewEmphasisNode creates a new EmphasisNode with the given content
func NewEmphasisNode(content string) *EmphasisNode {
	node := &EmphasisNode{
		BaseNode: NewBaseNode(NodeEmphasis),
	}
	node.SetContent(content)
	return node
}
