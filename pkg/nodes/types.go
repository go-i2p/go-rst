// pkg/nodes/base.go

package nodes

// NodeType represents the type of a node in the RST document structure
type NodeType int

// Node type constants define the possible types of nodes in the RST document tree
const (
	NodeHeading    NodeType = iota // Represents a section heading
	NodeParagraph                  // Represents a text paragraph
	NodeList                       // Represents an ordered or unordered list
	NodeListItem                   // Represents an item within a list
	NodeLink                       // Represents a hyperlink
	NodeEmphasis                   // Represents emphasized (italic) text
	NodeStrong                     // Represents strong (bold) text
	NodeMeta                       // Represents metadata information
	NodeDirective                  // Represents an RST directive
	NodeCode                       // Represents a code block
	NodeTable                      // Represents a table structure
	NodeBlockQuote                 // Represents a block quote
	NodeDoctest                    // Represents a doctest block
	NodeLineBlock                  // Represents a line block
	NodeComment                    // Represents a comment
)

// Node interface defines the common behavior for all RST document nodes
type Node interface {
	// Type returns the NodeType of this node
	Type() NodeType
	// Content returns the textual content of the node
	Content() string
	// SetContent sets the node's textual content
	SetContent(string)
	// Level returns the nesting level of the node
	Level() int
	// SetLevel sets the nesting level of the node
	SetLevel(int)
	// Children returns the node's child nodes
	Children() []Node
	// AddChild adds a child node to this node
	AddChild(Node)
}

// BaseNode provides the basic implementation of the Node interface
// that other node types can embed
type BaseNode struct {
	nodeType NodeType
	content  string
	level    int
	children []Node
}

// NewBaseNode creates a new BaseNode with the specified node type
//
// Parameters:
//   - nodeType: The type of node to create
//
// Returns:
//   - *BaseNode: A new base node instance
func NewBaseNode(nodeType NodeType) *BaseNode {
	return &BaseNode{
		nodeType: nodeType,
		children: make([]Node, 0),
	}
}

// Type returns the NodeType
func (n *BaseNode) Type() NodeType { return n.nodeType }

// Content returns the textual content of the node
func (n *BaseNode) Content() string { return n.content }

// SetContent sets the node's textual content
func (n *BaseNode) SetContent(content string) {
	n.content = content
}

// Level returns the nesting level of the node
func (n *BaseNode) Level() int { return n.level }

// SetLevel sets the nesting level of the node
func (n *BaseNode) SetLevel(level int) {
	n.level = level
}

// Children returns the node's child nodes
func (n *BaseNode) Children() []Node {
	return n.children
}

// AddChild adds a child node to this node
func (n *BaseNode) AddChild(child Node) {
	n.children = append(n.children, child)
}
