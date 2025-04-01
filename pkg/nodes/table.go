package nodes

import "fmt"

// TableNode represents a table structure
type TableNode struct {
	*BaseNode
	headers []string
	rows    [][]string
}

// NewTableNode creates a new TableNode
func NewTableNode() *TableNode {
	return &TableNode{
		BaseNode: NewBaseNode(NodeTable),
		headers:  make([]string, 0),
		rows:     make([][]string, 0),
	}
}

// SetHeaders sets the headers of the table
func (n *TableNode) SetHeaders(headers []string) {
	n.headers = headers
}

// AddRow adds a row to the table
func (n *TableNode) AddRow(row []string) {
	n.rows = append(n.rows, row)
}

// Headers returns the headers of the table
func (n *TableNode) Headers() []string { return n.headers }

// Rows returns the rows of the table
func (n *TableNode) Rows() [][]string { return n.rows }

// String representation for debugging
func (n *TableNode) String() string {
	return fmt.Sprintf("Table: %d columns x %d rows", len(n.headers), len(n.rows))
}
