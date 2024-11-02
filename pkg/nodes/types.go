// pkg/nodes/base.go

package nodes

type NodeType int

const (
    NodeHeading NodeType = iota
    NodeParagraph
    NodeList
    NodeListItem
    NodeLink
    NodeEmphasis
    NodeStrong
    NodeMeta
    NodeDirective
    NodeCode
    NodeTable
)

type Node interface {
    Type() NodeType
    Content() string
    SetContent(string)
    Level() int
    SetLevel(int)
    Children() []Node
    AddChild(Node)
}

type BaseNode struct {
    nodeType NodeType
    content  string
    level    int
    children []Node
}

func NewBaseNode(nodeType NodeType) *BaseNode {
    return &BaseNode{
        nodeType: nodeType,
        children: make([]Node, 0),
    }
}

func (n *BaseNode) Type() NodeType { return n.nodeType }

func (n *BaseNode) Content() string { return n.content }

func (n *BaseNode) SetContent(content string) {
    n.content = content
}

func (n *BaseNode) Level() int { return n.level }

func (n *BaseNode) SetLevel(level int) {
    n.level = level
}

func (n *BaseNode) Children() []Node {
    return n.children
}

func (n *BaseNode) AddChild(child Node) {
    n.children = append(n.children, child)
}