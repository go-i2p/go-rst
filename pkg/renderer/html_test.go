package renderer

import (
	"strings"
	"testing"

	"github.com/go-i2p/go-rst/pkg/nodes"
)

func TestHTMLRenderer_RenderFootnote(t *testing.T) {
	renderer := NewHTMLRenderer()

	t.Run("Numbered footnote", func(t *testing.T) {
		footnote := nodes.NewFootnoteNode("1", "This is footnote 1", false)
		renderer.buffer.Reset()
		renderer.renderFootnote(footnote)

		output := renderer.buffer.String()
		expected := `<div class="footnote" id="footnote-1">
<span class="label">[1]</span> This is footnote 1
</div>
`
		if output != expected {
			t.Errorf("Expected:\n%s\nGot:\n%s", expected, output)
		}
	})

	t.Run("Auto-numbered footnote", func(t *testing.T) {
		footnote := nodes.NewFootnoteNode("#", "Auto footnote content", true)
		renderer.buffer.Reset()
		renderer.renderFootnote(footnote)

		output := renderer.buffer.String()
		expected := `<div class="footnote" id="footnote-#">
<span class="label">#</span> Auto footnote content
</div>
`
		if output != expected {
			t.Errorf("Expected:\n%s\nGot:\n%s", expected, output)
		}
	})

	t.Run("Symbol footnote", func(t *testing.T) {
		footnote := nodes.NewFootnoteNode("*", "Symbol footnote", false)
		renderer.buffer.Reset()
		renderer.renderFootnote(footnote)

		output := renderer.buffer.String()
		expected := `<div class="footnote" id="footnote-*">
<span class="label">[*]</span> Symbol footnote
</div>
`
		if output != expected {
			t.Errorf("Expected:\n%s\nGot:\n%s", expected, output)
		}
	})

	t.Run("HTML escaping in footnote", func(t *testing.T) {
		footnote := nodes.NewFootnoteNode("<script>", "Content with <tags> & \"quotes\"", false)
		renderer.buffer.Reset()
		renderer.renderFootnote(footnote)

		output := renderer.buffer.String()
		if !strings.Contains(output, "&lt;script&gt;") {
			t.Error("Expected HTML escaping of footnote label")
		}
		if !strings.Contains(output, "&lt;tags&gt;") || !strings.Contains(output, "&amp;") {
			t.Error("Expected HTML escaping of footnote content")
		}
	})
}

func TestHTMLRenderer_RenderDefinitionList(t *testing.T) {
	renderer := NewHTMLRenderer()

	t.Run("Empty definition list", func(t *testing.T) {
		defList := nodes.NewDefinitionListNode()
		renderer.buffer.Reset()
		renderer.renderDefinitionList(defList)

		output := renderer.buffer.String()
		if output != "" {
			t.Errorf("Expected empty output for empty definition list, got: %s", output)
		}
	})

	t.Run("Single definition", func(t *testing.T) {
		defList := nodes.NewDefinitionListNode()
		defList.AddDefinition("Term 1", "Definition of term 1")
		renderer.buffer.Reset()
		renderer.renderDefinitionList(defList)

		output := renderer.buffer.String()
		expected := `<dl>
<dt>Term 1</dt>
<dd>Definition of term 1</dd>
</dl>
`
		if output != expected {
			t.Errorf("Expected:\n%s\nGot:\n%s", expected, output)
		}
	})

	t.Run("Multiple definitions", func(t *testing.T) {
		defList := nodes.NewDefinitionListNode()
		defList.AddDefinition("First Term", "First definition")
		defList.AddDefinition("Second Term", "Second definition")
		renderer.buffer.Reset()
		renderer.renderDefinitionList(defList)

		output := renderer.buffer.String()
		expected := `<dl>
<dt>First Term</dt>
<dd>First definition</dd>
<dt>Second Term</dt>
<dd>Second definition</dd>
</dl>
`
		if output != expected {
			t.Errorf("Expected:\n%s\nGot:\n%s", expected, output)
		}
	})

	t.Run("HTML escaping in definition list", func(t *testing.T) {
		defList := nodes.NewDefinitionListNode()
		defList.AddDefinition("<script>alert('xss')</script>", "Content with <tags> & \"quotes\"")
		renderer.buffer.Reset()
		renderer.renderDefinitionList(defList)

		output := renderer.buffer.String()
		if !strings.Contains(output, "&lt;script&gt;") {
			t.Error("Expected HTML escaping of term")
		}
		if !strings.Contains(output, "&lt;tags&gt;") || !strings.Contains(output, "&amp;") {
			t.Error("Expected HTML escaping of definition")
		}
	})
}

func TestHTMLRenderer_RenderFieldList(t *testing.T) {
	renderer := NewHTMLRenderer()

	t.Run("Empty field list", func(t *testing.T) {
		fieldList := nodes.NewFieldListNode()
		renderer.buffer.Reset()
		renderer.renderFieldList(fieldList)

		output := renderer.buffer.String()
		if output != "" {
			t.Errorf("Expected empty output for empty field list, got: %s", output)
		}
	})

	t.Run("Single field", func(t *testing.T) {
		fieldList := nodes.NewFieldListNode()
		fieldList.AddField("Author", "John Doe")
		renderer.buffer.Reset()
		renderer.renderFieldList(fieldList)

		output := renderer.buffer.String()
		// Note: map iteration order is not guaranteed, so check for presence
		if !strings.Contains(output, `<table class="docinfo">`) {
			t.Error("Expected table with docinfo class")
		}
		if !strings.Contains(output, `<th class="docinfo-name">Author:</th>`) {
			t.Error("Expected table header with Author field")
		}
		if !strings.Contains(output, `<td>John Doe</td>`) {
			t.Error("Expected table data with author value")
		}
		if !strings.Contains(output, `</table>`) {
			t.Error("Expected closing table tag")
		}
	})

	t.Run("Multiple fields", func(t *testing.T) {
		fieldList := nodes.NewFieldListNode()
		fieldList.AddField("Author", "Jane Smith")
		fieldList.AddField("Version", "1.0.0")
		renderer.buffer.Reset()
		renderer.renderFieldList(fieldList)

		output := renderer.buffer.String()
		// Check for all expected elements
		if !strings.Contains(output, `<table class="docinfo">`) {
			t.Error("Expected table with docinfo class")
		}
		if !strings.Contains(output, "Author:") || !strings.Contains(output, "Jane Smith") {
			t.Error("Expected Author field and value")
		}
		if !strings.Contains(output, "Version:") || !strings.Contains(output, "1.0.0") {
			t.Error("Expected Version field and value")
		}
	})

	t.Run("HTML escaping in field list", func(t *testing.T) {
		fieldList := nodes.NewFieldListNode()
		fieldList.AddField("<script>", "Content with <tags> & \"quotes\"")
		renderer.buffer.Reset()
		renderer.renderFieldList(fieldList)

		output := renderer.buffer.String()
		if !strings.Contains(output, "&lt;script&gt;") {
			t.Error("Expected HTML escaping of field name")
		}
		if !strings.Contains(output, "&lt;tags&gt;") || !strings.Contains(output, "&amp;") {
			t.Error("Expected HTML escaping of field value")
		}
	})
}

func TestHTMLRenderer_RenderNode_NewTypes(t *testing.T) {
	renderer := NewHTMLRenderer()

	t.Run("renderNode handles FootnoteNode", func(t *testing.T) {
		footnote := nodes.NewFootnoteNode("1", "Test footnote", false)
		renderer.buffer.Reset()
		renderer.renderNode(footnote)

		output := renderer.buffer.String()
		if !strings.Contains(output, `class="footnote"`) {
			t.Error("Expected footnote HTML output")
		}
	})

	t.Run("renderNode handles DefinitionListNode", func(t *testing.T) {
		defList := nodes.NewDefinitionListNode()
		defList.AddDefinition("Test", "Definition")
		renderer.buffer.Reset()
		renderer.renderNode(defList)

		output := renderer.buffer.String()
		if !strings.Contains(output, "<dl>") {
			t.Error("Expected definition list HTML output")
		}
	})

	t.Run("renderNode handles FieldListNode", func(t *testing.T) {
		fieldList := nodes.NewFieldListNode()
		fieldList.AddField("Test", "Value")
		renderer.buffer.Reset()
		renderer.renderNode(fieldList)

		output := renderer.buffer.String()
		if !strings.Contains(output, `class="docinfo"`) {
			t.Error("Expected field list HTML output")
		}
	})
}

func TestHTMLRenderer_Integration(t *testing.T) {
	t.Run("Full HTML document with new node types", func(t *testing.T) {
		renderer := NewHTMLRenderer()

		// Create a mixed set of nodes including new types
		nodeList := []nodes.Node{
			nodes.NewFootnoteNode("1", "Integration test footnote", false),
		}

		// Create definition list
		defList := nodes.NewDefinitionListNode()
		defList.AddDefinition("Integration", "Testing multiple components together")
		nodeList = append(nodeList, defList)

		// Create field list
		fieldList := nodes.NewFieldListNode()
		fieldList.AddField("Test Type", "Integration")
		fieldList.AddField("Status", "Passing")
		nodeList = append(nodeList, fieldList)

		output := renderer.Render(nodeList)

		// Verify complete HTML structure
		if !strings.Contains(output, "<!DOCTYPE html>") {
			t.Error("Expected DOCTYPE declaration")
		}
		if !strings.Contains(output, "<html>") || !strings.Contains(output, "</html>") {
			t.Error("Expected HTML tags")
		}
		if !strings.Contains(output, `class="footnote"`) {
			t.Error("Expected footnote in output")
		}
		if !strings.Contains(output, "<dl>") {
			t.Error("Expected definition list in output")
		}
		if !strings.Contains(output, `class="docinfo"`) {
			t.Error("Expected field list in output")
		}
	})
}
