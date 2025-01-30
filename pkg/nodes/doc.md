# nodes
--
    import "github.com/go-i2p/go-rst/pkg/nodes"


## Usage

#### func  GetIndentedContent

```go
func GetIndentedContent(node Node) string
```
GetIndentedContent Utility function to get node content with proper indentation

#### type BaseNode

```go
type BaseNode struct {
}
```

BaseNode provides the basic implementation of the Node interface that other node
types can embed

#### func  NewBaseNode

```go
func NewBaseNode(nodeType NodeType) *BaseNode
```
NewBaseNode creates a new BaseNode with the specified node type

Parameters:

    - nodeType: The type of node to create

Returns:

    - *BaseNode: A new base node instance

#### func (*BaseNode) AddChild

```go
func (n *BaseNode) AddChild(child Node)
```
AddChild adds a child node to this node

#### func (*BaseNode) Children

```go
func (n *BaseNode) Children() []Node
```
Children returns the node's child nodes

#### func (*BaseNode) Content

```go
func (n *BaseNode) Content() string
```
Content returns the textual content of the node

#### func (*BaseNode) Level

```go
func (n *BaseNode) Level() int
```
Level returns the nesting level of the node

#### func (*BaseNode) SetContent

```go
func (n *BaseNode) SetContent(content string)
```
SetContent sets the node's textual content

#### func (*BaseNode) SetLevel

```go
func (n *BaseNode) SetLevel(level int)
```
SetLevel sets the nesting level of the node

#### func (*BaseNode) Type

```go
func (n *BaseNode) Type() NodeType
```
Type returns the NodeType

#### type BlockQuoteNode

```go
type BlockQuoteNode struct {
	*BaseNode
}
```

BlockQuoteNode represents an indented block quote

#### func  NewBlockQuoteNode

```go
func NewBlockQuoteNode(content, attribution string) *BlockQuoteNode
```
NewBlockQuoteNode creates a new BlockQuoteNode with the given content

#### func (*BlockQuoteNode) Attribution

```go
func (n *BlockQuoteNode) Attribution() string
```
Attribution returns the quote attribution if any

#### func (*BlockQuoteNode) String

```go
func (n *BlockQuoteNode) String() string
```
String representation for debugging

#### type CodeNode

```go
type CodeNode struct {
	*BaseNode
}
```

CodeNode represents a code block

#### func  NewCodeNode

```go
func NewCodeNode(language, content string, lineNumbers bool) *CodeNode
```
NewCodeNode creates a new CodeNode with the given language and content

#### func (*CodeNode) Language

```go
func (n *CodeNode) Language() string
```
Language returns the language of the code block

#### func (*CodeNode) LineNumbers

```go
func (n *CodeNode) LineNumbers() bool
```
LineNumbers returns the line numbers flag of the code block

#### func (*CodeNode) String

```go
func (n *CodeNode) String() string
```
String representation for debugging

#### type DirectiveNode

```go
type DirectiveNode struct {
	*BaseNode
}
```

DirectiveNode represents an RST directive

#### func  NewDirectiveNode

```go
func NewDirectiveNode(name string, args []string) *DirectiveNode
```
NewDirectiveNode creates a new DirectiveNode with the given name and arguments

#### func (*DirectiveNode) Arguments

```go
func (n *DirectiveNode) Arguments() []string
```
Arguments returns the arguments of the directive

#### func (*DirectiveNode) Name

```go
func (n *DirectiveNode) Name() string
```
Name returns the name of the directive

#### func (*DirectiveNode) RawContent

```go
func (n *DirectiveNode) RawContent() string
```
RawContent returns the raw content of the directive

#### func (*DirectiveNode) SetRawContent

```go
func (n *DirectiveNode) SetRawContent(content string)
```
SetRawContent sets the raw content of the directive

#### func (*DirectiveNode) String

```go
func (n *DirectiveNode) String() string
```
String representation for debugging

#### type DoctestNode

```go
type DoctestNode struct {
	*BaseNode
}
```

DoctestNode represents a doctest block with expected output

#### func  NewDoctestNode

```go
func NewDoctestNode(command, expected string) *DoctestNode
```

#### func (*DoctestNode) Command

```go
func (n *DoctestNode) Command() string
```

#### func (*DoctestNode) Expected

```go
func (n *DoctestNode) Expected() string
```

#### func (*DoctestNode) String

```go
func (n *DoctestNode) String() string
```

#### type EmphasisNode

```go
type EmphasisNode struct {
	*BaseNode
}
```

EmphasisNode represents emphasized text (italic)

#### func  NewEmphasisNode

```go
func NewEmphasisNode(content string) *EmphasisNode
```
NewEmphasisNode creates a new EmphasisNode with the given content

#### type HeadingNode

```go
type HeadingNode struct {
	*BaseNode
}
```

HeadingNode represents a section heading in RST

#### func  NewHeadingNode

```go
func NewHeadingNode(content string, level int) *HeadingNode
```
NewHeadingNode creates a new HeadingNode with the given content and level

#### func (*HeadingNode) String

```go
func (n *HeadingNode) String() string
```
String representations for debugging

#### type LinkNode

```go
type LinkNode struct {
	*BaseNode
}
```

LinkNode represents a hyperlink

#### func  NewLinkNode

```go
func NewLinkNode(text, url, title string) *LinkNode
```
NewLinkNode creates a new LinkNode with the given text, URL, and title

#### func (*LinkNode) String

```go
func (n *LinkNode) String() string
```
String representation for debugging

#### func (*LinkNode) Title

```go
func (n *LinkNode) Title() string
```
Title returns the URL of the link

#### func (*LinkNode) URL

```go
func (n *LinkNode) URL() string
```
URL returns the URL of the link

#### type ListItemNode

```go
type ListItemNode struct {
	*BaseNode
}
```

ListItemNode represents an individual list item

#### func  NewListItemNode

```go
func NewListItemNode(content string) *ListItemNode
```
NewListItemNode creates a new ListItemNode with the given content

#### type ListNode

```go
type ListNode struct {
	*BaseNode
}
```

ListNode represents an ordered or unordered list

#### func  NewListNode

```go
func NewListNode(ordered bool) *ListNode
```
NewListNode creates a new ListNode with the given ordered flag

#### func (*ListNode) IsOrdered

```go
func (n *ListNode) IsOrdered() bool
```
IsOrdered returns true if the list is ordered, false otherwise

#### func (*ListNode) String

```go
func (n *ListNode) String() string
```
String representation for debugging

#### type MetaNode

```go
type MetaNode struct {
	*BaseNode
}
```

MetaNode represents metadata information

#### func  NewMetaNode

```go
func NewMetaNode(key, value string) *MetaNode
```
NewMetaNode creates a new MetaNode with the given key and value

#### func (*MetaNode) Key

```go
func (n *MetaNode) Key() string
```
Key returns the key of the metadata

#### type Node

```go
type Node interface {
	// Type returns the NodeType of this node
	Type() NodeType
	// Content returns the textual content of the node
	Content() string
	// SetContent sets the node's textual content
	SetContent(string)
	// Level returns the nesting level of the node
	Level() int
	// SetLevel sets the nesting level of the node
	SetLevel(int)
	// Children returns the node's child nodes
	Children() []Node
	// AddChild adds a child node to this node
	AddChild(Node)
}
```

Node interface defines the common behavior for all RST document nodes

#### type NodeType

```go
type NodeType int
```

NodeType represents the type of a node in the RST document structure

```go
const (
	NodeHeading    NodeType = iota // Represents a section heading
	NodeParagraph                  // Represents a text paragraph
	NodeList                       // Represents an ordered or unordered list
	NodeListItem                   // Represents an item within a list
	NodeLink                       // Represents a hyperlink
	NodeEmphasis                   // Represents emphasized (italic) text
	NodeStrong                     // Represents strong (bold) text
	NodeMeta                       // Represents metadata information
	NodeDirective                  // Represents an RST directive
	NodeCode                       // Represents a code block
	NodeTable                      // Represents a table structure
	NodeBlockQuote                 // Represents a block quote
	NodeDoctest                    // Represents a doctest block
)
```
Node type constants define the possible types of nodes in the RST document tree

#### type ParagraphNode

```go
type ParagraphNode struct {
	*BaseNode
}
```

ParagraphNode represents a text paragraph

#### func  NewParagraphNode

```go
func NewParagraphNode(content string) *ParagraphNode
```
NewParagraphNode creates a new ParagraphNode with the given content

#### func (*ParagraphNode) String

```go
func (n *ParagraphNode) String() string
```
String representation for debugging

#### type StrongNode

```go
type StrongNode struct {
	*BaseNode
}
```

StrongNode represents strong text (bold)

#### func  NewStrongNode

```go
func NewStrongNode(content string) *StrongNode
```
NewStrongNode creates a new StrongNode with the given content

#### type TableNode

```go
type TableNode struct {
	*BaseNode
}
```

TableNode represents a table structure

#### func  NewTableNode

```go
func NewTableNode() *TableNode
```
NewTableNode creates a new TableNode

#### func (*TableNode) AddRow

```go
func (n *TableNode) AddRow(row []string)
```
AddRow adds a row to the table

#### func (*TableNode) Headers

```go
func (n *TableNode) Headers() []string
```
Headers returns the headers of the table

#### func (*TableNode) Rows

```go
func (n *TableNode) Rows() [][]string
```
Rows returns the rows of the table

#### func (*TableNode) SetHeaders

```go
func (n *TableNode) SetHeaders(headers []string)
```
SetHeaders sets the headers of the table

#### func (*TableNode) String

```go
func (n *TableNode) String() string
```
String representation for debugging
