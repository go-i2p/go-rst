package nodes

import "fmt"

// ListNode represents an ordered or unordered list
type ListNode struct {
	*BaseNode
	ordered bool
	indent  int
}

func (n *ListNode) Indent() int {
	return n.indent
}

func (n *ListNode) SetIndent(indent int) {
	n.indent = indent
}

// AppendChild adds a list item node as a child to this list node
func (n *ListNode) AppendChild(listItem *ListItemNode) {
	n.BaseNode.AddChild(listItem)
}

// NewListNode creates a new ListNode with the given ordered flag
func NewListNode(ordered bool) *ListNode {
	node := &ListNode{
		BaseNode: NewBaseNode(NodeList),
		ordered:  ordered,
	}
	return node
}

// IsOrdered returns true if the list is ordered, false otherwise
func (n *ListNode) IsOrdered() bool {
	return n.ordered
}

// ListItemNode represents an individual list item
type ListItemNode struct {
	*BaseNode
}

// NewListItemNode creates a new ListItemNode with the given content
func NewListItemNode(content string) *ListItemNode {
	node := &ListItemNode{
		BaseNode: NewBaseNode(NodeListItem),
	}
	node.SetContent(content)
	return node
}

// String representation for debugging
func (n *ListNode) String() string {
	listType := "Unordered"
	if n.ordered {
		listType = "Ordered"
	}
	return fmt.Sprintf("%s List with %d items", listType, len(n.Children()))
}
