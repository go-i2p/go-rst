package nodes

// MetaNode represents metadata information
type MetaNode struct {
	*BaseNode
	key string
}

// NewMetaNode creates a new MetaNode with the given key and value
func NewMetaNode(key, value string) *MetaNode {
	node := &MetaNode{
		BaseNode: NewBaseNode(NodeMeta),
		key:      key,
	}
	node.SetContent(value)
	return node
}

// Key returns the key of the metadata
func (n *MetaNode) Key() string { return n.key }
