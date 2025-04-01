package nodes

import "fmt"

type TransitionNode struct {
	*BaseNode
	character rune
}

func NewTransitionNode(char rune) *TransitionNode {
	return &TransitionNode{
		BaseNode:  NewBaseNode(NodeTransition),
		character: char,
	}
}

func (n *TransitionNode) Character() rune {
	return n.character
}

func (n *TransitionNode) String() string {
	return fmt.Sprintf("Transition: %c", n.character)
}
