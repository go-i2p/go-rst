package nodes

import "fmt"

// CommentNode represents RST comments starting with ..
type CommentNode struct {
	*BaseNode
}

func NewCommentNode(content string) *CommentNode {
	node := &CommentNode{
		BaseNode: NewBaseNode(NodeComment),
	}
	node.SetContent(content)
	return node
}

func (n *CommentNode) String() string {
	return fmt.Sprintf("Comment: %s", n.Content())
}
