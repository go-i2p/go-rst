// pkg/renderer/pdf.go

package renderer

import (
	"fmt"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/go-i2p/go-rst/pkg/nodes"
)

// PDFRenderer implements rendering RST nodes to PDF format
type PDFRenderer struct {
	pdf        *gofpdf.Fpdf
	marginLeft float64
	marginTop  float64
	fontSize   float64
	lineHeight float64
	indent     float64
}

// NewPDFRenderer creates a new PDF renderer
func NewPDFRenderer() *PDFRenderer {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 12)

	return &PDFRenderer{
		pdf:        pdf,
		marginLeft: 20,
		marginTop:  20,
		fontSize:   12,
		lineHeight: 6,
		indent:     10,
	}
}

// Render renders a slice of nodes to PDF
func (r *PDFRenderer) Render(nodes []nodes.Node) error {
	for _, node := range nodes {
		if err := r.renderNode(node); err != nil {
			return err
		}
	}
	return nil
}

// renderNode handles individual node rendering
func (r *PDFRenderer) renderNode(node nodes.Node) error {
	switch n := node.(type) {
	case *nodes.HeadingNode:
		return r.renderHeading(n)
	case *nodes.ParagraphNode:
		return r.renderParagraph(n)
	case *nodes.ListNode:
		return r.renderList(n)
	case *nodes.CodeNode:
		return r.renderCode(n)
	case *nodes.TableNode:
		return r.renderTable(n)
	case *nodes.DirectiveNode:
		return r.renderDirective(n)
	default:
		return r.renderChildren(node)
	}
}

// SaveToFile saves the PDF to a file
func (r *PDFRenderer) SaveToFile(filename string) error {
	return r.pdf.OutputFileAndClose(filename)
}

func (r *PDFRenderer) renderHeading(node *nodes.HeadingNode) error {
	// Calculate font size based on heading level
	fontSize := r.fontSize + (6 - float64(node.Level())*2)
	r.pdf.SetFont("Arial", "B", fontSize)

	// Add some spacing before heading
	r.pdf.Ln(r.lineHeight)

	r.pdf.Cell(0, r.lineHeight, node.Content())
	r.pdf.Ln(r.lineHeight * 1.5)

	// Reset font
	r.pdf.SetFont("Arial", "", r.fontSize)
	return r.renderChildren(node)
}

func (r *PDFRenderer) renderParagraph(node *nodes.ParagraphNode) error {
	r.pdf.MultiCell(0, r.lineHeight, node.Content(), "", "", false)
	r.pdf.Ln(r.lineHeight)
	return r.renderChildren(node)
}

func (r *PDFRenderer) renderList(node *nodes.ListNode) error {
	currentX := r.pdf.GetX()
	currentY := r.pdf.GetY()

	for i, child := range node.Children() {
		if node.IsOrdered() {
			// Ordered list: use numbers
			r.pdf.SetXY(currentX, currentY)
			r.pdf.Cell(10, r.lineHeight, fmt.Sprintf("%d.", i+1))
			r.pdf.SetX(currentX + r.indent)
		} else {
			// Unordered list: use bullets
			r.pdf.SetXY(currentX, currentY)
			r.pdf.Cell(10, r.lineHeight, "•")
			r.pdf.SetX(currentX + r.indent)
		}

		if listItem, ok := child.(*nodes.ListItemNode); ok {
			r.pdf.MultiCell(0, r.lineHeight, listItem.Content(), "", "", false)
		}

		currentY = r.pdf.GetY() + r.lineHeight
	}

	r.pdf.Ln(r.lineHeight)
	return nil
}

func (r *PDFRenderer) renderCode(node *nodes.CodeNode) error {
	// Set monospace font for code
	r.pdf.SetFont("Courier", "", r.fontSize)

	// Add a light gray background
	startY := r.pdf.GetY()
	content := node.Content()
	lines := strings.Split(content, "\n")

	// Calculate height
	height := float64(len(lines)) * r.lineHeight

	// Draw background
	r.pdf.SetFillColor(240, 240, 240)
	r.pdf.Rect(r.marginLeft-2, startY-2, 170, height+4, "F")

	// Add language identifier if present
	if node.Language() != "" {
		r.pdf.SetFont("Arial", "I", r.fontSize-2)
		r.pdf.Text(r.marginLeft, startY-4, node.Language())
		r.pdf.SetFont("Courier", "", r.fontSize)
	}

	// Render code content
	for _, line := range lines {
		if node.LineNumbers() {
			// Add line numbers if enabled
			lineNum := fmt.Sprintf("%3d ", r.pdf.PageCount())
			r.pdf.Text(r.marginLeft, r.pdf.GetY(), lineNum)
			r.pdf.SetX(r.marginLeft + 15)
		}
		r.pdf.MultiCell(0, r.lineHeight, line, "", "", false)
	}

	// Reset font
	r.pdf.SetFont("Arial", "", r.fontSize)
	r.pdf.Ln(r.lineHeight)
	return nil
}

func (r *PDFRenderer) renderTable(node *nodes.TableNode) error {
	headers := node.Headers()
	rows := node.Rows()

	// Calculate column widths
	colWidth := 170.0 / float64(len(headers))

	// Draw headers
	r.pdf.SetFont("Arial", "B", r.fontSize)
	r.pdf.SetFillColor(240, 240, 240)

	for _, header := range headers {
		r.pdf.CellFormat(colWidth, r.lineHeight, header, "1", 0, "", true, 0, "")
	}
	r.pdf.Ln(-1)

	// Draw rows
	r.pdf.SetFont("Arial", "", r.fontSize)
	r.pdf.SetFillColor(255, 255, 255)

	for _, row := range rows {
		for _, cell := range row {
			r.pdf.CellFormat(colWidth, r.lineHeight, cell, "1", 0, "", false, 0, "")
		}
		r.pdf.Ln(-1)
	}

	r.pdf.Ln(r.lineHeight)
	return nil
}

func (r *PDFRenderer) renderDirective(node *nodes.DirectiveNode) error {
	// Handle special directives
	switch node.Name() {
	case "image":
		// Handle image directive
		if len(node.Arguments()) > 0 {
			imagePath := node.Arguments()[0]
			r.pdf.Image(imagePath, r.pdf.GetX(), r.pdf.GetY(), 0, 0, false, "", 0, "")
			r.pdf.Ln(r.lineHeight)
		}
	case "note", "warning", "important":
		// Handle admonitions
		r.pdf.SetFillColor(245, 245, 245)
		startY := r.pdf.GetY()
		r.pdf.MultiCell(0, r.lineHeight, node.RawContent(), "", "", true)
		r.pdf.SetY(startY + r.lineHeight)
	}

	return r.renderChildren(node)
}

func (r *PDFRenderer) renderChildren(node nodes.Node) error {
	return r.Render(node.Children())
}
