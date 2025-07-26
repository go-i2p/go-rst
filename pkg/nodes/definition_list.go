package nodes

// DefinitionListNode represents a reStructuredText definition list element.
// It embeds BaseNode and stores terms and their corresponding definitions.
type DefinitionListNode struct {
	*BaseNode
	terms       []string
	definitions []string
}

// NewDefinitionListNode creates a new DefinitionListNode with empty terms and definitions slices.
//
// Returns:
//   - *DefinitionListNode: A new definition list node instance
func NewDefinitionListNode() *DefinitionListNode {
	return &DefinitionListNode{
		BaseNode:    NewBaseNode(NodeDefinitionList),
		terms:       make([]string, 0),
		definitions: make([]string, 0),
	}
}

// AddDefinition adds a term and its corresponding definition to the list.
//
// Parameters:
//   - term: The term being defined
//   - definition: The definition of the term
func (n *DefinitionListNode) AddDefinition(term, definition string) {
	n.terms = append(n.terms, term)
	n.definitions = append(n.definitions, definition)
}

// Terms returns a copy of the terms slice
func (n *DefinitionListNode) Terms() []string {
	result := make([]string, len(n.terms))
	copy(result, n.terms)
	return result
}

// Definitions returns a copy of the definitions slice
func (n *DefinitionListNode) Definitions() []string {
	result := make([]string, len(n.definitions))
	copy(result, n.definitions)
	return result
}

// Count returns the number of term-definition pairs
func (n *DefinitionListNode) Count() int {
	return len(n.terms)
}
