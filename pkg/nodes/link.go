package nodes

import "fmt"

// LinkNode represents a hyperlink
type LinkNode struct {
	*BaseNode
	url   string
	title string
}

// NewLinkNode creates a new LinkNode with the given text, URL, and title
func NewLinkNode(text, url, title string) *LinkNode {
	node := &LinkNode{
		BaseNode: NewBaseNode(NodeLink),
		url:      url,
		title:    title,
	}
	node.SetContent(text)
	return node
}

// URL returns the URL of the link
func (n *LinkNode) URL() string { return n.url }

// Title returns the URL of the link
func (n *LinkNode) Title() string { return n.title }

// String representation for debugging
func (n *LinkNode) String() string {
	return fmt.Sprintf("Link[%s](%s)", n.Content(), n.url)
}
