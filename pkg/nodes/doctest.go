package nodes

// DoctestNode represents a Python doctest block with code and expected output.
type DoctestNode struct {
	BaseNode
	code           string
	expectedOutput string
}

// NewDoctestNode creates a new doctest node.
func NewDoctestNode() *DoctestNode {
	return &DoctestNode{
		BaseNode: BaseNode{
			nodeType: NodeDoctest,
		},
		code:           "",
		expectedOutput: "",
	}
}

// SetCode sets the code content of the doctest.
func (n *DoctestNode) SetCode(code string) {
	n.code = code
}

// Code returns the code content of the doctest.
func (n *DoctestNode) Code() string {
	return n.code
}

// SetExpectedOutput sets the expected output of the doctest.
func (n *DoctestNode) SetExpectedOutput(output string) {
	n.expectedOutput = output
}

// ExpectedOutput returns the expected output of the doctest.
func (n *DoctestNode) ExpectedOutput() string {
	return n.expectedOutput
}

// Content returns the code of this node.
func (n *DoctestNode) Content() string {
	return n.code
}

// SetContent sets the code of this node.
func (n *DoctestNode) SetContent(content string) {
	n.code = content
}
