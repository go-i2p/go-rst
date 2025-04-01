package nodes

import "fmt"

type TitleNode struct {
	*BaseNode
	level int
}

func NewTitleNode(content string, level int) *TitleNode {
	node := &TitleNode{
		BaseNode: NewBaseNode(NodeTitle),
		level:    level,
	}
	node.SetContent(content)
	return node
}

func (n *TitleNode) Level() int {
	return n.level
}

func (n *TitleNode) String() string {
	return fmt.Sprintf("Title[%d]: %s", n.Level(), n.Content())
}
