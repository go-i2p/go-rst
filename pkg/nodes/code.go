package nodes

import "fmt"

// CodeNode represents a code block
type CodeNode struct {
	*BaseNode
	language    string
	lineNumbers bool
}

// NewCodeNode creates a new CodeNode with the given language and content
func NewCodeNode(language, content string, lineNumbers bool) *CodeNode {
	node := &CodeNode{
		BaseNode:    NewBaseNode(NodeCode),
		language:    language,
		lineNumbers: lineNumbers,
	}
	node.SetContent(content)
	return node
}

// Language returns the language of the code block
func (n *CodeNode) Language() string { return n.language }

// LineNumbers returns the line numbers flag of the code block
func (n *CodeNode) LineNumbers() bool { return n.lineNumbers }

// String representation for debugging
func (n *CodeNode) String() string {
	return fmt.Sprintf("Code[%s]: %d bytes", n.language, len(n.Content()))
}
