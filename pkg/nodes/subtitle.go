package nodes

import "fmt"

type SubtitleNode struct {
	*BaseNode
}

func NewSubtitleNode(content string) *SubtitleNode {
	node := &SubtitleNode{
		BaseNode: NewBaseNode(NodeSubtitle),
	}
	node.SetContent(content)
	return node
}

func (n *SubtitleNode) String() string {
	return fmt.Sprintf("Subtitle: %s", n.Content())
}
