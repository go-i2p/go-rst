package nodes

// FootnoteNode represents a reStructuredText footnote element.
// It embeds BaseNode and stores the footnote label and whether it is auto-numbered.
type FootnoteNode struct {
	*BaseNode
	label      string
	autoNumber bool
}

// NewFootnoteNode creates a new FootnoteNode with the given label, content, and autoNumber flag.
//
// Parameters:
//   - label: The footnote label (number, symbol, or auto-number marker)
//   - content: The footnote content text
//   - autoNumber: Whether this footnote is auto-numbered
//
// Returns:
//   - *FootnoteNode: A new footnote node instance
func NewFootnoteNode(label, content string, autoNumber bool) *FootnoteNode {
	node := &FootnoteNode{
		BaseNode:   NewBaseNode(NodeFootnote),
		label:      label,
		autoNumber: autoNumber,
	}
	node.SetContent(content)
	return node
}

// Label returns the footnote label
func (n *FootnoteNode) Label() string {
	return n.label
}

// IsAutoNumber returns whether this footnote is auto-numbered
func (n *FootnoteNode) IsAutoNumber() bool {
	return n.autoNumber
}
