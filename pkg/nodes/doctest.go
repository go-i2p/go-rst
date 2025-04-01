package nodes

import "fmt"

// DoctestNode represents a doctest block with expected output
type DoctestNode struct {
	*BaseNode
	command  string
	expected string
}

func NewDoctestNode(command, expected string) *DoctestNode {
	return &DoctestNode{
		BaseNode: NewBaseNode(NodeDoctest),
		command:  command,
		expected: expected,
	}
}

func (n *DoctestNode) Command() string {
	return n.command
}

func (n *DoctestNode) Expected() string {
	return n.expected
}

func (n *DoctestNode) String() string {
	return fmt.Sprintf("Doctest: %s -> %s", n.command, n.expected)
}
