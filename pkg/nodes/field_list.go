package nodes

// FieldListNode represents a reStructuredText field list, typically used for document metadata.
// It embeds BaseNode and stores fields as a map of key-value pairs.
type FieldListNode struct {
	*BaseNode
	fields map[string]string
}

// NewFieldListNode creates a new FieldListNode with an empty fields map.
//
// Returns:
//   - *FieldListNode: A new field list node instance
func NewFieldListNode() *FieldListNode {
	return &FieldListNode{
		BaseNode: NewBaseNode(NodeFieldList),
		fields:   make(map[string]string),
	}
}

// AddField adds a field key-value pair to the field list.
//
// Parameters:
//   - key: The field name
//   - value: The field value
func (n *FieldListNode) AddField(key, value string) {
	n.fields[key] = value
}

// GetField retrieves the value for a given field key.
//
// Parameters:
//   - key: The field name to look up
//
// Returns:
//   - string: The field value
//   - bool: Whether the field exists
func (n *FieldListNode) GetField(key string) (string, bool) {
	value, exists := n.fields[key]
	return value, exists
}

// Fields returns a copy of the fields map
func (n *FieldListNode) Fields() map[string]string {
	result := make(map[string]string)
	for k, v := range n.fields {
		result[k] = v
	}
	return result
}

// HasFields returns whether the field list has any fields
func (n *FieldListNode) HasFields() bool {
	return len(n.fields) > 0
}
