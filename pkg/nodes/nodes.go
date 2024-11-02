// pkg/nodes/nodes.go

package nodes

import (
	"fmt"
	"strings"
)

// HeadingNode represents a section heading in RST
type HeadingNode struct {
	*BaseNode
}

func NewHeadingNode(content string, level int) *HeadingNode {
	node := &HeadingNode{
		BaseNode: NewBaseNode(NodeHeading),
	}
	node.SetContent(content)
	node.SetLevel(level)
	return node
}

// ParagraphNode represents a text paragraph
type ParagraphNode struct {
	*BaseNode
}

func NewParagraphNode(content string) *ParagraphNode {
	node := &ParagraphNode{
		BaseNode: NewBaseNode(NodeParagraph),
	}
	node.SetContent(content)
	return node
}

// ListNode represents an ordered or unordered list
type ListNode struct {
	*BaseNode
	ordered bool
}

func NewListNode(ordered bool) *ListNode {
	node := &ListNode{
		BaseNode: NewBaseNode(NodeList),
		ordered:  ordered,
	}
	return node
}

func (n *ListNode) IsOrdered() bool {
	return n.ordered
}

// ListItemNode represents an individual list item
type ListItemNode struct {
	*BaseNode
}

func NewListItemNode(content string) *ListItemNode {
	node := &ListItemNode{
		BaseNode: NewBaseNode(NodeListItem),
	}
	node.SetContent(content)
	return node
}

// LinkNode represents a hyperlink
type LinkNode struct {
	*BaseNode
	url   string
	title string
}

func NewLinkNode(text, url, title string) *LinkNode {
	node := &LinkNode{
		BaseNode: NewBaseNode(NodeLink),
		url:      url,
		title:    title,
	}
	node.SetContent(text)
	return node
}

func (n *LinkNode) URL() string   { return n.url }
func (n *LinkNode) Title() string { return n.title }

// EmphasisNode represents emphasized text (italic)
type EmphasisNode struct {
	*BaseNode
}

func NewEmphasisNode(content string) *EmphasisNode {
	node := &EmphasisNode{
		BaseNode: NewBaseNode(NodeEmphasis),
	}
	node.SetContent(content)
	return node
}

// StrongNode represents strong text (bold)
type StrongNode struct {
	*BaseNode
}

func NewStrongNode(content string) *StrongNode {
	node := &StrongNode{
		BaseNode: NewBaseNode(NodeStrong),
	}
	node.SetContent(content)
	return node
}

// MetaNode represents metadata information
type MetaNode struct {
	*BaseNode
	key string
}

func NewMetaNode(key, value string) *MetaNode {
	node := &MetaNode{
		BaseNode: NewBaseNode(NodeMeta),
		key:      key,
	}
	node.SetContent(value)
	return node
}

func (n *MetaNode) Key() string { return n.key }

// DirectiveNode represents an RST directive
type DirectiveNode struct {
	*BaseNode
	name       string
	arguments  []string
	rawContent string
}

func NewDirectiveNode(name string, args []string) *DirectiveNode {
	node := &DirectiveNode{
		BaseNode:   NewBaseNode(NodeDirective),
		name:       name,
		arguments:  args,
		rawContent: "",
	}
	return node
}

func (n *DirectiveNode) Name() string        { return n.name }
func (n *DirectiveNode) Arguments() []string { return n.arguments }
func (n *DirectiveNode) RawContent() string  { return n.rawContent }
func (n *DirectiveNode) SetRawContent(content string) {
	n.rawContent = content
}

// CodeNode represents a code block
type CodeNode struct {
	*BaseNode
	language    string
	lineNumbers bool
}

func NewCodeNode(language string, content string, lineNumbers bool) *CodeNode {
	node := &CodeNode{
		BaseNode:    NewBaseNode(NodeCode),
		language:    language,
		lineNumbers: lineNumbers,
	}
	node.SetContent(content)
	return node
}

func (n *CodeNode) Language() string    { return n.language }
func (n *CodeNode) LineNumbers() bool   { return n.lineNumbers }

// TableNode represents a table structure
type TableNode struct {
	*BaseNode
	headers []string
	rows    [][]string
}

func NewTableNode() *TableNode {
	return &TableNode{
		BaseNode: NewBaseNode(NodeTable),
		headers:  make([]string, 0),
		rows:     make([][]string, 0),
	}
}

func (n *TableNode) SetHeaders(headers []string) {
	n.headers = headers
}

func (n *TableNode) AddRow(row []string) {
	n.rows = append(n.rows, row)
}

func (n *TableNode) Headers() []string   { return n.headers }
func (n *TableNode) Rows() [][]string    { return n.rows }

// Utility function to get node content with proper indentation
func GetIndentedContent(node Node) string {
	content := node.Content()
	if node.Level() > 0 {
		indent := strings.Repeat("    ", node.Level())
		lines := strings.Split(content, "\n")
		for i, line := range lines {
			lines[i] = indent + line
		}
		content = strings.Join(lines, "\n")
	}
	return content
}

// String representations for debugging
func (n *HeadingNode) String() string {
	return fmt.Sprintf("Heading[%d]: %s", n.Level(), n.Content())
}

func (n *ParagraphNode) String() string {
	return fmt.Sprintf("Paragraph: %s", n.Content())
}

func (n *ListNode) String() string {
	listType := "Unordered"
	if n.ordered {
		listType = "Ordered"
	}
	return fmt.Sprintf("%s List with %d items", listType, len(n.Children()))
}

func (n *LinkNode) String() string {
	return fmt.Sprintf("Link[%s](%s)", n.Content(), n.url)
}

func (n *DirectiveNode) String() string {
	return fmt.Sprintf("Directive[%s]: %s", n.name, n.Content())
}

func (n *CodeNode) String() string {
	return fmt.Sprintf("Code[%s]: %d bytes", n.language, len(n.Content()))
}

func (n *TableNode) String() string {
	return fmt.Sprintf("Table: %d columns x %d rows", len(n.headers), len(n.rows))
}

