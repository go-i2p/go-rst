package nodes

import "fmt"

// DirectiveNode represents an RST directive
type DirectiveNode struct {
	*BaseNode
	name       string
	arguments  []string
	rawContent string
}

// NewDirectiveNode creates a new DirectiveNode with the given name and arguments
func NewDirectiveNode(name string, args []string) *DirectiveNode {
	node := &DirectiveNode{
		BaseNode:   NewBaseNode(NodeDirective),
		name:       name,
		arguments:  args,
		rawContent: "",
	}
	return node
}

// Name returns the name of the directive
func (n *DirectiveNode) Name() string { return n.name }

// Arguments returns the arguments of the directive
func (n *DirectiveNode) Arguments() []string { return n.arguments }

// RawContent returns the raw content of the directive
func (n *DirectiveNode) RawContent() string { return n.rawContent }

// SetRawContent sets the raw content of the directive
func (n *DirectiveNode) SetRawContent(content string) {
	n.rawContent = content
}

// String representation for debugging
func (n *DirectiveNode) String() string {
	return fmt.Sprintf("Directive[%s]: %s", n.name, n.Content())
}
