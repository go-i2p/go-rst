// pkg/renderer/markdown.go

package renderer

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/go-i2p/go-rst/pkg/nodes"
)

// MarkdownRenderer implements a Markdown renderer with the same interface as HTMLRenderer
type MarkdownRenderer struct {
	output bytes.Buffer
}

// NewMarkdownRenderer creates a new Markdown renderer
func NewMarkdownRenderer() *MarkdownRenderer {
	return &MarkdownRenderer{}
}

// pkg/renderer/markdown.go

// ... (previous imports and struct definition remain the same)

// Render renders a slice of nodes to Markdown
func (r *MarkdownRenderer) Render(nodes []nodes.Node) error {
	for _, node := range nodes {
		if err := r.RenderNode(node); err != nil {
			return err
		}
	}
	return nil
}

// RenderNode renders a single node to Markdown
func (r *MarkdownRenderer) RenderNode(node nodes.Node) error {
	switch n := node.(type) {
	case *nodes.HeadingNode:
		return r.RenderHeading(n)
	case *nodes.ParagraphNode:
		return r.RenderParagraph(n)
	case *nodes.ListNode:
		return r.RenderList(n)
	case *nodes.ListItemNode:
		return r.RenderListItem(n)
	case *nodes.EmphasisNode:
		return r.RenderEmphasis(n)
	case *nodes.StrongNode:
		return r.RenderStrong(n)
	case *nodes.LinkNode:
		return r.RenderLink(n)
	case *nodes.CodeNode:
		return r.RenderCode(n)
	case *nodes.TableNode:
		return r.RenderTable(n)
	case *nodes.DirectiveNode:
		return r.RenderDirective(n)
	case *nodes.MetaNode:
		return r.RenderMeta(n)
	default:
		return r.RenderChildren(node)
	}
}

// RenderHeading renders a heading node
func (r *MarkdownRenderer) RenderHeading(node *nodes.HeadingNode) error {
	r.output.WriteString("\n")
	r.output.WriteString(strings.Repeat("#", node.Level()))
	r.output.WriteString(" ")
	r.output.WriteString(node.Content())
	r.output.WriteString("\n")
	return r.RenderChildren(node)
}

// RenderParagraph renders a paragraph node
func (r *MarkdownRenderer) RenderParagraph(node *nodes.ParagraphNode) error {
	r.output.WriteString("\n")
	r.output.WriteString(node.Content())
	r.output.WriteString("\n")
	return r.RenderChildren(node)
}

// RenderList renders a list node
func (r *MarkdownRenderer) RenderList(node *nodes.ListNode) error {
	r.output.WriteString("\n")
	return r.RenderChildren(node)
}

// RenderListItem renders a list item node
func (r *MarkdownRenderer) RenderListItem(node *nodes.ListItemNode) error {
	// Default to unordered list items with "-"
	r.output.WriteString("- ")
	r.output.WriteString(node.Content())
	r.output.WriteString("\n")
	return r.RenderChildren(node)
}

// RenderEmphasis renders an emphasis node
func (r *MarkdownRenderer) RenderEmphasis(node *nodes.EmphasisNode) error {
	r.output.WriteString("*")
	r.output.WriteString(node.Content())
	r.output.WriteString("*")
	return r.RenderChildren(node)
}

// RenderStrong renders a strong node
func (r *MarkdownRenderer) RenderStrong(node *nodes.StrongNode) error {
	r.output.WriteString("**")
	r.output.WriteString(node.Content())
	r.output.WriteString("**")
	return r.RenderChildren(node)
}

// RenderLink renders a link node
func (r *MarkdownRenderer) RenderLink(node *nodes.LinkNode) error {
	if title := node.Title(); title != "" {
		r.output.WriteString(fmt.Sprintf("[%s](%s \"%s\")", node.Content(), node.URL(), title))
	} else {
		r.output.WriteString(fmt.Sprintf("[%s](%s)", node.Content(), node.URL()))
	}
	return nil
}

// RenderCode renders a code node
func (r *MarkdownRenderer) RenderCode(node *nodes.CodeNode) error {
	r.output.WriteString("\n```")
	if node.Language() != "" {
		r.output.WriteString(node.Language())
	}
	r.output.WriteString("\n")
	r.output.WriteString(node.Content())
	r.output.WriteString("\n```\n")
	return nil
}

// RenderTable renders a table node
func (r *MarkdownRenderer) RenderTable(node *nodes.TableNode) error {
	headers := node.Headers()
	rows := node.Rows()

	// Headers
	r.output.WriteString("\n|")
	for _, header := range headers {
		r.output.WriteString(fmt.Sprintf(" %s |", header))
	}
	r.output.WriteString("\n|")

	// Separator
	for range headers {
		r.output.WriteString(" --- |")
	}
	r.output.WriteString("\n")

	// Rows
	for _, row := range rows {
		r.output.WriteString("|")
		for _, cell := range row {
			r.output.WriteString(fmt.Sprintf(" %s |", cell))
		}
		r.output.WriteString("\n")
	}
	r.output.WriteString("\n")
	return nil
}

// RenderDirective renders a directive node
func (r *MarkdownRenderer) RenderDirective(node *nodes.DirectiveNode) error {
	switch node.Name() {
	case "image":
		if len(node.Arguments()) > 0 {
			r.output.WriteString(fmt.Sprintf("\n![%s](%s)\n", node.Content(), node.Arguments()[0]))
		}
	case "note":
		r.output.WriteString("\n> **Note**\n")
		r.output.WriteString("> " + strings.Replace(node.RawContent(), "\n", "\n> ", -1) + "\n")
	case "warning":
		r.output.WriteString("\n> **Warning**\n")
		r.output.WriteString("> " + strings.Replace(node.RawContent(), "\n", "\n> ", -1) + "\n")
	default:
		r.output.WriteString(fmt.Sprintf("\n<!-- %s: %s -->\n", node.Name(), node.RawContent()))
	}
	return r.RenderChildren(node)
}

// RenderMeta renders a meta node
func (r *MarkdownRenderer) RenderMeta(node *nodes.MetaNode) error {
	if r.output.Len() == 0 {
		// If at start of document, use YAML front matter
		r.output.WriteString("---\n")
		r.output.WriteString(fmt.Sprintf("%s: %s\n", node.Key(), node.Content()))
		r.output.WriteString("---\n")
	} else {
		// Otherwise use HTML comment
		r.output.WriteString(fmt.Sprintf("\n<!-- %s: %s -->\n", node.Key(), node.Content()))
	}
	return nil
}

// RenderChildren renders child nodes
func (r *MarkdownRenderer) RenderChildren(node nodes.Node) error {
	for _, child := range node.Children() {
		if err := r.RenderNode(child); err != nil {
			return err
		}
	}
	return nil
}

// String returns the rendered markdown as a string
func (r *MarkdownRenderer) String() string {
	return r.output.String()
}
