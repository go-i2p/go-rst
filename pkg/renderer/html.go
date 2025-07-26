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
		r.renderHeading(n)
	case *nodes.ParagraphNode:
		r.renderParagraph(n)
	case *nodes.ListNode:
		r.renderList(n)
	case *nodes.LinkNode:
		r.renderLink(n)
	case *nodes.EmphasisNode:
		r.renderEmphasis(n)
	case *nodes.StrongNode:
		r.renderStrong(n)
	case *nodes.CodeNode:
		r.renderCode(n)
	case *nodes.TableNode:
		r.renderTable(n)
	case *nodes.DirectiveNode:
		r.renderDirective(n)
	case *nodes.BlockQuoteNode:
		r.renderBlockQuote(n)
	case *nodes.DoctestNode:
		r.renderDoctest(n)
	case *nodes.LineBlockNode:
		r.renderLineBlock(n)
	case *nodes.CommentNode:
		r.renderComment(n)
	case *nodes.TitleNode:
		r.renderTitle(n)
	case *nodes.SubtitleNode:
		r.renderSubtitle(n)
	case *nodes.TransitionNode:
		r.renderTransition(n)
	case *nodes.FootnoteNode:
		r.renderFootnote(n)
	case *nodes.DefinitionListNode:
		r.renderDefinitionList(n)
	case *nodes.FieldListNode:
		r.renderFieldList(n)
	}
}

// renderHeading renders a heading node as HTML.
func (r *HTMLRenderer) renderHeading(n *nodes.HeadingNode) {
	r.buffer.WriteString(fmt.Sprintf("<h%d>%s</h%d>\n",
		n.Level(),
		html.EscapeString(n.Content()),
		n.Level()))
}

// renderParagraph renders a paragraph node as HTML.
func (r *HTMLRenderer) renderParagraph(n *nodes.ParagraphNode) {
	r.buffer.WriteString(fmt.Sprintf("<p>%s</p>\n",
		html.EscapeString(n.Content())))
}

// renderList renders a list node as HTML.
func (r *HTMLRenderer) renderList(n *nodes.ListNode) {
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
}

// renderLink renders a link node as HTML.
func (r *HTMLRenderer) renderLink(n *nodes.LinkNode) {
	r.buffer.WriteString(fmt.Sprintf("<a href=\"%s\" title=\"%s\">%s</a>",
		html.EscapeString(n.URL()),
		html.EscapeString(n.Title()),
		html.EscapeString(n.Content())))
}

// renderEmphasis renders an emphasis node as HTML.
func (r *HTMLRenderer) renderEmphasis(n *nodes.EmphasisNode) {
	r.buffer.WriteString(fmt.Sprintf("<em>%s</em>",
		html.EscapeString(n.Content())))
}

// renderStrong renders a strong node as HTML.
func (r *HTMLRenderer) renderStrong(n *nodes.StrongNode) {
	r.buffer.WriteString(fmt.Sprintf("<strong>%s</strong>",
		html.EscapeString(n.Content())))
}

// renderCode renders a code node as HTML.
func (r *HTMLRenderer) renderCode(n *nodes.CodeNode) {
	r.buffer.WriteString(fmt.Sprintf("<pre><code class=\"language-%s\">%s</code></pre>\n",
		html.EscapeString(n.Language()),
		html.EscapeString(n.Content())))
}

// renderTable renders a table node as HTML.
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

// renderDirective renders a directive node as HTML.
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

// renderBlockQuote renders a block quote node as HTML.
func (r *HTMLRenderer) renderBlockQuote(n *nodes.BlockQuoteNode) {
	r.buffer.WriteString("<blockquote>")
	r.buffer.WriteString(html.EscapeString(n.Content()))
	if attr := n.Attribution(); attr != "" {
		r.buffer.WriteString("<cite>")
		r.buffer.WriteString(html.EscapeString(attr))
		r.buffer.WriteString("</cite>")
	}
	r.buffer.WriteString("</blockquote>\n")
}

// renderDoctest renders a doctest node as HTML.
func (r *HTMLRenderer) renderDoctest(n *nodes.DoctestNode) {
	r.buffer.WriteString("<div class=\"doctest\">")
	r.buffer.WriteString("<pre class=\"doctest-command\">>> ")
	r.buffer.WriteString(html.EscapeString(n.Command()))
	r.buffer.WriteString("</pre>")
	if n.Expected() != "" {
		r.buffer.WriteString("<pre class=\"doctest-output\">")
		r.buffer.WriteString(html.EscapeString(n.Expected()))
		r.buffer.WriteString("</pre>")
	}
	r.buffer.WriteString("</div>\n")
}

// renderLineBlock renders a line block node as HTML.
func (r *HTMLRenderer) renderLineBlock(n *nodes.LineBlockNode) {
	r.buffer.WriteString("<div class=\"line-block\">")
	for _, line := range n.Lines() {
		r.buffer.WriteString("<div class=\"line\">")
		r.buffer.WriteString(html.EscapeString(strings.TrimSpace(line)))
		r.buffer.WriteString("</div>\n")
	}
	r.buffer.WriteString("</div>\n")
}

// renderComment renders a comment node as HTML.
func (r *HTMLRenderer) renderComment(n *nodes.CommentNode) {
	r.buffer.WriteString("<!-- ")
	r.buffer.WriteString(html.EscapeString(n.Content()))
	r.buffer.WriteString(" -->\n")
}

// renderTitle renders a title node as HTML.
func (r *HTMLRenderer) renderTitle(n *nodes.TitleNode) {
	r.buffer.WriteString(fmt.Sprintf("<h1 class=\"title\">%s</h1>\n",
		html.EscapeString(n.Content())))
}

// renderSubtitle renders a subtitle node as HTML.
func (r *HTMLRenderer) renderSubtitle(n *nodes.SubtitleNode) {
	r.buffer.WriteString(fmt.Sprintf("<h2 class=\"subtitle\">%s</h2>\n",
		html.EscapeString(n.Content())))
}

// renderTransition renders a transition node as HTML.
func (r *HTMLRenderer) renderTransition(n *nodes.TransitionNode) {
	r.buffer.WriteString("<hr class=\"docutils\">\n")
}

// renderFootnote renders a footnote node as HTML.
// It creates a footnote reference with a link to the footnote definition.
func (r *HTMLRenderer) renderFootnote(n *nodes.FootnoteNode) {
	label := html.EscapeString(n.Label())
	content := html.EscapeString(n.Content())

	if n.IsAutoNumber() {
		// Auto-numbered footnote
		r.buffer.WriteString(fmt.Sprintf(
			"<div class=\"footnote\" id=\"footnote-%s\">\n"+
				"<span class=\"label\">%s</span> %s\n"+
				"</div>\n",
			label, label, content))
	} else {
		// Named or numbered footnote
		r.buffer.WriteString(fmt.Sprintf(
			"<div class=\"footnote\" id=\"footnote-%s\">\n"+
				"<span class=\"label\">[%s]</span> %s\n"+
				"</div>\n",
			label, label, content))
	}
}

// renderDefinitionList renders a definition list node as HTML.
// It creates a <dl> element with <dt> terms and <dd> definitions.
func (r *HTMLRenderer) renderDefinitionList(n *nodes.DefinitionListNode) {
	if n.Count() == 0 {
		return
	}

	r.buffer.WriteString("<dl>\n")

	terms := n.Terms()
	definitions := n.Definitions()

	for i := 0; i < n.Count(); i++ {
		term := html.EscapeString(terms[i])
		definition := html.EscapeString(definitions[i])

		r.buffer.WriteString(fmt.Sprintf("<dt>%s</dt>\n", term))
		r.buffer.WriteString(fmt.Sprintf("<dd>%s</dd>\n", definition))
	}

	r.buffer.WriteString("</dl>\n")
}

// renderFieldList renders a field list node as HTML.
// It creates a table structure for metadata fields with proper semantic markup.
func (r *HTMLRenderer) renderFieldList(n *nodes.FieldListNode) {
	if !n.HasFields() {
		return
	}

	r.buffer.WriteString("<table class=\"docinfo\">\n<tbody>\n")

	for key, value := range n.Fields() {
		escapedKey := html.EscapeString(key)
		escapedValue := html.EscapeString(value)

		r.buffer.WriteString(fmt.Sprintf(
			"<tr><th class=\"docinfo-name\">%s:</th>\n"+
				"<td>%s</td></tr>\n",
			escapedKey, escapedValue))
	}

	r.buffer.WriteString("</tbody>\n</table>\n")
}

// RenderPretty renders the given nodes as pretty-formatted HTML.
func (r *HTMLRenderer) RenderPretty(nodes []nodes.Node) string {
	// First get the regular HTML output
	rawHTML := r.Render(nodes)

	// Let gohtml handle the formatting
	prettyHTML := gohtml.Format(rawHTML)

	return prettyHTML
}
