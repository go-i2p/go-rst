# renderer
--
    import "github.com/go-i2p/go-rst/pkg/renderer"


## Usage

#### type HTMLRenderer

```go
type HTMLRenderer struct {
}
```

HTMLRederer is a renderer that renders nodes to HTML.

#### func  NewHTMLRenderer

```go
func NewHTMLRenderer() *HTMLRenderer
```
NewHTMLRederer creates a new HTMLRederer.

#### func (*HTMLRenderer) Render

```go
func (r *HTMLRenderer) Render(nodes []nodes.Node) string
```
Render renders nodes to HTML.

#### func (*HTMLRenderer) RenderPretty

```go
func (r *HTMLRenderer) RenderPretty(nodes []nodes.Node) string
```
RenderPretty renders the given nodes as pretty-formatted HTML.

#### type MarkdownRenderer

```go
type MarkdownRenderer struct {
}
```

MarkdownRenderer implements a Markdown renderer with the same interface as
HTMLRenderer

#### func  NewMarkdownRenderer

```go
func NewMarkdownRenderer() *MarkdownRenderer
```
NewMarkdownRenderer creates a new Markdown renderer

#### func (*MarkdownRenderer) Render

```go
func (r *MarkdownRenderer) Render(nodes []nodes.Node) error
```
Render renders a slice of nodes to Markdown

#### func (*MarkdownRenderer) RenderChildren

```go
func (r *MarkdownRenderer) RenderChildren(node nodes.Node) error
```
RenderChildren renders child nodes

#### func (*MarkdownRenderer) RenderCode

```go
func (r *MarkdownRenderer) RenderCode(node *nodes.CodeNode) error
```
RenderCode renders a code node

#### func (*MarkdownRenderer) RenderDirective

```go
func (r *MarkdownRenderer) RenderDirective(node *nodes.DirectiveNode) error
```
RenderDirective renders a directive node

#### func (*MarkdownRenderer) RenderEmphasis

```go
func (r *MarkdownRenderer) RenderEmphasis(node *nodes.EmphasisNode) error
```
RenderEmphasis renders an emphasis node

#### func (*MarkdownRenderer) RenderHeading

```go
func (r *MarkdownRenderer) RenderHeading(node *nodes.HeadingNode) error
```
RenderHeading renders a heading node

#### func (*MarkdownRenderer) RenderLink

```go
func (r *MarkdownRenderer) RenderLink(node *nodes.LinkNode) error
```
RenderLink renders a link node

#### func (*MarkdownRenderer) RenderList

```go
func (r *MarkdownRenderer) RenderList(node *nodes.ListNode) error
```
RenderList renders a list node

#### func (*MarkdownRenderer) RenderListItem

```go
func (r *MarkdownRenderer) RenderListItem(node *nodes.ListItemNode) error
```
RenderListItem renders a list item node

#### func (*MarkdownRenderer) RenderMeta

```go
func (r *MarkdownRenderer) RenderMeta(node *nodes.MetaNode) error
```
RenderMeta renders a meta node

#### func (*MarkdownRenderer) RenderNode

```go
func (r *MarkdownRenderer) RenderNode(node nodes.Node) error
```
RenderNode renders a single node to Markdown

#### func (*MarkdownRenderer) RenderParagraph

```go
func (r *MarkdownRenderer) RenderParagraph(node *nodes.ParagraphNode) error
```
RenderParagraph renders a paragraph node

#### func (*MarkdownRenderer) RenderStrong

```go
func (r *MarkdownRenderer) RenderStrong(node *nodes.StrongNode) error
```
RenderStrong renders a strong node

#### func (*MarkdownRenderer) RenderTable

```go
func (r *MarkdownRenderer) RenderTable(node *nodes.TableNode) error
```
RenderTable renders a table node

#### func (*MarkdownRenderer) String

```go
func (r *MarkdownRenderer) String() string
```
String returns the rendered markdown as a string

#### type PDFRenderer

```go
type PDFRenderer struct {
}
```

PDFRenderer implements rendering RST nodes to PDF format

#### func  NewPDFRenderer

```go
func NewPDFRenderer() *PDFRenderer
```
NewPDFRenderer creates a new PDF renderer

#### func (*PDFRenderer) Render

```go
func (r *PDFRenderer) Render(nodes []nodes.Node) error
```
Render renders a slice of nodes to PDF

#### func (*PDFRenderer) SaveToFile

```go
func (r *PDFRenderer) SaveToFile(filename string) error
```
SaveToFile saves the PDF to a file
