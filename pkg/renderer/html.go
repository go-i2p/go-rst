// pkg/renderer/html.go

package renderer

import (
	"bytes"
	"fmt"
	"html"
	"strings"

	"github.com/go-i2p/go-rst/pkg/nodes"
	"github.com/yosssi/gohtml"
)

// HTMLRederer is a renderer that renders nodes to HTML.
type HTMLRenderer struct {
	buffer bytes.Buffer
}

// NewHTMLRederer creates a new HTMLRederer.
func NewHTMLRenderer() *HTMLRenderer {
	return &HTMLRenderer{}
}

// Render renders nodes to HTML.
func (r *HTMLRenderer) Render(nodes []nodes.Node) string {
	r.buffer.Reset()

	r.buffer.WriteString("<!DOCTYPE html>\n<html>\n<head>\n")
	r.renderMeta(nodes)
	r.buffer.WriteString("</head>\n<body>\n")

	for _, node := range nodes {
		r.renderNode(node)
	}

	r.buffer.WriteString("</body>\n</html>")
	return r.buffer.String()
}

func (r *HTMLRenderer) renderMeta(nodelist []nodes.Node) {
	r.buffer.WriteString("<meta charset=\"UTF-8\">\n")

	for _, node := range nodelist {
		switch n := node.(type) {
		case *nodes.MetaNode:
			r.buffer.WriteString(fmt.Sprintf("<meta name=\"%s\" content=\"%s\">\n",
				html.EscapeString(n.Key()),
				html.EscapeString(n.Content())))
		}
	}
}

func (r *HTMLRenderer) renderNode(node nodes.Node) {
	switch n := node.(type) {
	case *nodes.HeadingNode:
		r.buffer.WriteString(fmt.Sprintf("<h%d>%s</h%d>\n",
			n.Level(),
			html.EscapeString(n.Content()),
			n.Level()))

	case *nodes.ParagraphNode:
		r.buffer.WriteString(fmt.Sprintf("<p>%s</p>\n",
			html.EscapeString(n.Content())))

	case *nodes.ListNode:
		tag := "ul"
		if n.IsOrdered() {
			tag = "ol"
		}
		r.buffer.WriteString(fmt.Sprintf("<%s>\n", tag))
		for _, child := range n.Children() {
			if item, ok := child.(*nodes.ListItemNode); ok {
				r.buffer.WriteString(fmt.Sprintf("<li>%s</li>\n",
					html.EscapeString(item.Content())))
			}
		}
		r.buffer.WriteString(fmt.Sprintf("</%s>\n", tag))

	case *nodes.LinkNode:
		r.buffer.WriteString(fmt.Sprintf("<a href=\"%s\" title=\"%s\">%s</a>",
			html.EscapeString(n.URL()),
			html.EscapeString(n.Title()),
			html.EscapeString(n.Content())))

	case *nodes.EmphasisNode:
		r.buffer.WriteString(fmt.Sprintf("<em>%s</em>",
			html.EscapeString(n.Content())))

	case *nodes.StrongNode:
		r.buffer.WriteString(fmt.Sprintf("<strong>%s</strong>",
			html.EscapeString(n.Content())))

	case *nodes.CodeNode:
		r.buffer.WriteString(fmt.Sprintf("<pre><code class=\"language-%s\">%s</code></pre>\n",
			html.EscapeString(n.Language()),
			html.EscapeString(n.Content())))

	case *nodes.TableNode:
		r.renderTable(n)

	case *nodes.DirectiveNode:
		r.renderDirective(n)
	case *nodes.BlockQuoteNode:
		r.buffer.WriteString("<blockquote>")
		r.buffer.WriteString(html.EscapeString(n.Content()))
		if attr := n.Attribution(); attr != "" {
			r.buffer.WriteString("<cite>")
			r.buffer.WriteString(html.EscapeString(attr))
			r.buffer.WriteString("</cite>")
		}
		r.buffer.WriteString("</blockquote>\n")

	}
}

func (r *HTMLRenderer) renderTable(table *nodes.TableNode) {
	r.buffer.WriteString("<table>\n")

	// Render headers
	if len(table.Headers()) > 0 {
		r.buffer.WriteString("<thead><tr>\n")
		for _, header := range table.Headers() {
			r.buffer.WriteString(fmt.Sprintf("<th>%s</th>",
				html.EscapeString(header)))
		}
		r.buffer.WriteString("</tr></thead>\n")
	}

	// Render rows
	r.buffer.WriteString("<tbody>\n")
	for _, row := range table.Rows() {
		r.buffer.WriteString("<tr>\n")
		for _, cell := range row {
			r.buffer.WriteString(fmt.Sprintf("<td>%s</td>",
				html.EscapeString(cell)))
		}
		r.buffer.WriteString("</tr>\n")
	}
	r.buffer.WriteString("</tbody></table>\n")
}

func (r *HTMLRenderer) renderDirective(directive *nodes.DirectiveNode) {
	switch directive.Name() {
	case "image":
		if len(directive.Arguments()) > 0 {
			alt := ""
			if len(directive.Arguments()) > 1 {
				alt = strings.Join(directive.Arguments()[1:], " ")
			}
			r.buffer.WriteString(fmt.Sprintf("<img src=\"%s\" alt=\"%s\">\n",
				html.EscapeString(directive.Arguments()[0]),
				html.EscapeString(alt)))
		}

	case "note":
		r.buffer.WriteString(fmt.Sprintf("<div class=\"note\">%s</div>\n",
			html.EscapeString(directive.RawContent())))

	case "warning":
		r.buffer.WriteString(fmt.Sprintf("<div class=\"warning\">%s</div>\n",
			html.EscapeString(directive.RawContent())))
	}
}

// RenderPretty renders the given nodes as pretty-formatted HTML.
func (r *HTMLRenderer) RenderPretty(nodes []nodes.Node) string {
	// First get the regular HTML output
	rawHTML := r.Render(nodes)

	// Let gohtml handle the formatting
	prettyHTML := gohtml.Format(rawHTML)

	return prettyHTML
}
