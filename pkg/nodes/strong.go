package nodes

// StrongNode represents strong text (bold)
type StrongNode struct {
	*BaseNode
}

// NewStrongNode creates a new StrongNode with the given content
func NewStrongNode(content string) *StrongNode {
	node := &StrongNode{
		BaseNode: NewBaseNode(NodeStrong),
	}
	node.SetContent(content)
	return node
}
