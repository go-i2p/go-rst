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

// NewHeadingNode creates a new HeadingNode with the given content and level
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

// NewParagraphNode creates a new ParagraphNode with the given content
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

// NewListNode creates a new ListNode with the given ordered flag
func NewListNode(ordered bool) *ListNode {
	node := &ListNode{
		BaseNode: NewBaseNode(NodeList),
		ordered:  ordered,
	}
	return node
}

// IsOrdered returns true if the list is ordered, false otherwise
func (n *ListNode) IsOrdered() bool {
	return n.ordered
}

// ListItemNode represents an individual list item
type ListItemNode struct {
	*BaseNode
}

// NewListItemNode creates a new ListItemNode with the given content
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

// NewLinkNode creates a new LinkNode with the given text, URL, and title
func NewLinkNode(text, url, title string) *LinkNode {
	node := &LinkNode{
		BaseNode: NewBaseNode(NodeLink),
		url:      url,
		title:    title,
	}
	node.SetContent(text)
	return node
}

// URL returns the URL of the link
func (n *LinkNode) URL() string { return n.url }

// Title returns the URL of the link
func (n *LinkNode) Title() string { return n.title }

// EmphasisNode represents emphasized text (italic)
type EmphasisNode struct {
	*BaseNode
}

// NewEmphasisNode creates a new EmphasisNode with the given content
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

// NewStrongNode creates a new StrongNode with the given content
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

// DirectiveNode represents an RST directive
type DirectiveNode struct {
	*BaseNode
	name       string
	arguments  []string
	rawContent string
}

// NewDirectiveNode creates a new DirectiveNode with the given name and arguments
func NewDirectiveNode(name string, args []string) *DirectiveNode {
	node := &DirectiveNode{
		BaseNode:   NewBaseNode(NodeDirective),
		name:       name,
		arguments:  args,
		rawContent: "",
	}
	return node
}

// Name returns the name of the directive
func (n *DirectiveNode) Name() string { return n.name }

// Arguments returns the arguments of the directive
func (n *DirectiveNode) Arguments() []string { return n.arguments }

// RawContent returns the raw content of the directive
func (n *DirectiveNode) RawContent() string { return n.rawContent }

// SetRawContent sets the raw content of the directive
func (n *DirectiveNode) SetRawContent(content string) {
	n.rawContent = content
}

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

// GetIndentedContent Utility function to get node content with proper indentation
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

// String representation for debugging
func (n *ParagraphNode) String() string {
	return fmt.Sprintf("Paragraph: %s", n.Content())
}

// String representation for debugging
func (n *ListNode) String() string {
	listType := "Unordered"
	if n.ordered {
		listType = "Ordered"
	}
	return fmt.Sprintf("%s List with %d items", listType, len(n.Children()))
}

// String representation for debugging
func (n *LinkNode) String() string {
	return fmt.Sprintf("Link[%s](%s)", n.Content(), n.url)
}

// String representation for debugging
func (n *DirectiveNode) String() string {
	return fmt.Sprintf("Directive[%s]: %s", n.name, n.Content())
}

// String representation for debugging
func (n *CodeNode) String() string {
	return fmt.Sprintf("Code[%s]: %d bytes", n.language, len(n.Content()))
}

// String representation for debugging
func (n *TableNode) String() string {
	return fmt.Sprintf("Table: %d columns x %d rows", len(n.headers), len(n.rows))
}

// BlockQuoteNode represents an indented block quote
type BlockQuoteNode struct {
	*BaseNode
	attribution string
}

// NewBlockQuoteNode creates a new BlockQuoteNode with the given content
func NewBlockQuoteNode(content, attribution string) *BlockQuoteNode {
	node := &BlockQuoteNode{
		BaseNode:    NewBaseNode(NodeBlockQuote),
		attribution: attribution,
	}
	node.SetContent(content)
	return node
}

// Attribution returns the quote attribution if any
func (n *BlockQuoteNode) Attribution() string {
	return n.attribution
}

// String representation for debugging
func (n *BlockQuoteNode) String() string {
	if n.attribution != "" {
		return fmt.Sprintf("BlockQuote: %s -- %s", n.Content(), n.attribution)
	}
	return fmt.Sprintf("BlockQuote: %s", n.Content())
}

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

// LineBlockNode represents poetry-style line blocks that preserve line breaks
type LineBlockNode struct {
	*BaseNode
	lines []string
}

func NewLineBlockNode(lines []string) *LineBlockNode {
	node := &LineBlockNode{
		BaseNode: NewBaseNode(NodeLineBlock),
		lines:    lines,
	}
	return node
}

func (n *LineBlockNode) Lines() []string {
	return n.lines
}

func (n *LineBlockNode) String() string {
	return fmt.Sprintf("LineBlock: %d lines", len(n.lines))
}

// CommentNode represents RST comments starting with ..
type CommentNode struct {
	*BaseNode
}

func NewCommentNode(content string) *CommentNode {
	node := &CommentNode{
		BaseNode: NewBaseNode(NodeComment),
	}
	node.SetContent(content)
	return node
}

func (n *CommentNode) String() string {
	return fmt.Sprintf("Comment: %s", n.Content())
}

type TitleNode struct {
	*BaseNode
	level int
}

func NewTitleNode(content string, level int) *TitleNode {
	node := &TitleNode{
		BaseNode: NewBaseNode(NodeTitle),
		level:    level,
	}
	node.SetContent(content)
	return node
}

func (n *TitleNode) Level() int {
	return n.level
}

func (n *TitleNode) String() string {
	return fmt.Sprintf("Title[%d]: %s", n.Level(), n.Content())
}

type SubtitleNode struct {
	*BaseNode
}

func NewSubtitleNode(content string) *SubtitleNode {
	node := &SubtitleNode{
		BaseNode: NewBaseNode(NodeSubtitle),
	}
	node.SetContent(content)
	return node
}

func (n *SubtitleNode) String() string {
	return fmt.Sprintf("Subtitle: %s", n.Content())
}
