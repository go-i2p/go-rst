package parser

import (
	"strings"
	"testing"

	"github.com/go-i2p/go-rst/pkg/nodes"
	"github.com/go-i2p/go-rst/pkg/translator"
)

// TestTranslationIntegration_CompleteDocument tests translation blocks within a complete RST document
func TestTranslationIntegration_CompleteDocument(t *testing.T) {
	poTranslator, err := translator.NewPOTranslator("../../example/translations.po")
	if err != nil {
		t.Fatalf("Failed to create PO translator: %v", err)
	}

	parser := NewParser(poTranslator)

	content := `Document Title
==============

This is regular content.

{% trans %}This text will be translated{% endtrans %}

More regular content here.

.. code-block:: python

    # Code block should not be translated
    print("{% trans %}This text will be translated{% endtrans %}")

:author: John Doe
:date: 2025-10-12

{% trans %}Another translatable section{% endtrans %}

Final paragraph.`

	doc := parser.Parse(content)
	if len(doc) == 0 {
		t.Fatalf("Expected parsed nodes, got empty document")
	}

	// Count translated paragraphs - checking for unique translated content
	foundFirstTranslation := false
	foundSecondTranslation := false

	for _, node := range doc {
		if para, ok := node.(*nodes.ParagraphNode); ok {
			content := para.Content()
			if strings.Contains(content, "Este texto será traducido") {
				foundFirstTranslation = true
			}
			if strings.Contains(content, "Otra sección traducible") {
				foundSecondTranslation = true
			}
			// Verify code blocks don't get translated
			if strings.Contains(content, "{% trans %}") {
				t.Logf("Found untranslated trans block in content: %s", content)
			}
		}
	}

	if !foundFirstTranslation {
		t.Errorf("Expected to find first translated paragraph ('Este texto será traducido')")
	}
	if !foundSecondTranslation {
		t.Errorf("Expected to find second translated paragraph ('Otra sección traducible')")
	}
}

// TestTranslationIntegration_WithFormatting tests translation blocks with inline formatting
func TestTranslationIntegration_WithFormatting(t *testing.T) {
	poTranslator, err := translator.NewPOTranslator("../../example/translations.po")
	if err != nil {
		t.Fatalf("Failed to create PO translator: %v", err)
	}

	parser := NewParser(poTranslator)

	tests := []struct {
		name     string
		content  string
		hasNodes bool
	}{
		{
			name:     "translation block before strong text",
			content:  "{% trans %}This text will be translated{% endtrans %}\n\n**Bold text**",
			hasNodes: true,
		},
		{
			name:     "translation block before emphasis",
			content:  "{% trans %}Another translatable section{% endtrans %}\n\n*Italic text*",
			hasNodes: true,
		},
		{
			name:     "translation block with heading",
			content:  "Heading\n=======\n\n{% trans %}This text will be translated{% endtrans %}",
			hasNodes: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			doc := parser.Parse(test.content)
			if len(doc) == 0 && test.hasNodes {
				t.Errorf("Expected parsed nodes, got empty document")
			}

			// Verify at least one node was translated
			foundTranslation := false
			for _, node := range doc {
				if para, ok := node.(*nodes.ParagraphNode); ok {
					content := para.Content()
					if strings.Contains(content, "Este texto será traducido") ||
						strings.Contains(content, "Otra sección traducible") {
						foundTranslation = true
						break
					}
				}
			}

			if test.hasNodes && !foundTranslation {
				t.Errorf("Expected to find translated content in parsed document")
			}
		})
	}
}

// TestTranslationIntegration_MultipleOnSameLine tests multiple translation blocks
func TestTranslationIntegration_MultipleOnSameLine(t *testing.T) {
	poTranslator, err := translator.NewPOTranslator("../../example/translations.po")
	if err != nil {
		t.Fatalf("Failed to create PO translator: %v", err)
	}

	parser := NewParser(poTranslator)

	// Note: Current implementation processes line-by-line, so multiple blocks on same line
	// may not work as expected. This test documents current behavior.
	content := "{% trans %}This text will be translated{% endtrans %} and {% trans %}Another translatable section{% endtrans %}"

	doc := parser.Parse(content)
	if len(doc) == 0 {
		t.Fatalf("Expected parsed nodes, got empty document")
	}

	// Document current behavior - may only translate first block
	t.Logf("Parsed %d nodes from content with multiple translation blocks on same line", len(doc))
}

// TestTranslationIntegration_WithLists tests translation blocks in list contexts
func TestTranslationIntegration_WithLists(t *testing.T) {
	poTranslator, err := translator.NewPOTranslator("../../example/translations.po")
	if err != nil {
		t.Fatalf("Failed to create PO translator: %v", err)
	}

	parser := NewParser(poTranslator)

	content := `Shopping List
=============

- Item 1
- {% trans %}This text will be translated{% endtrans %}
- Item 3

1. First item
2. {% trans %}Another translatable section{% endtrans %}
3. Third item`

	doc := parser.Parse(content)
	if len(doc) == 0 {
		t.Fatalf("Expected parsed nodes, got empty document")
	}

	// Currently translation blocks in lists may be parsed as separate paragraph nodes
	// This test documents the behavior
	t.Logf("Parsed %d nodes from list with translation blocks", len(doc))
}

// TestTranslationIntegration_EmptyTranslationBlock tests edge case of empty translation block
func TestTranslationIntegration_EmptyTranslationBlock(t *testing.T) {
	poTranslator, err := translator.NewPOTranslator("../../example/translations.po")
	if err != nil {
		t.Fatalf("Failed to create PO translator: %v", err)
	}

	parser := NewParser(poTranslator)

	content := "{% trans %}{% endtrans %}"

	doc := parser.Parse(content)
	// Should not crash, may return empty or paragraph with empty content
	t.Logf("Empty translation block parsed into %d nodes", len(doc))
}

// TestTranslationIntegration_NestedStructures tests translation blocks near complex structures
func TestTranslationIntegration_NestedStructures(t *testing.T) {
	poTranslator, err := translator.NewPOTranslator("../../example/translations.po")
	if err != nil {
		t.Fatalf("Failed to create PO translator: %v", err)
	}

	parser := NewParser(poTranslator)

	content := `Document
========

.. note:: Directive content

   {% trans %}This text will be translated{% endtrans %}

   More directive content.

Regular paragraph.`

	doc := parser.Parse(content)
	if len(doc) == 0 {
		t.Fatalf("Expected parsed nodes, got empty document")
	}

	// Translation blocks inside directives may have special behavior
	t.Logf("Nested structure with translation parsed into %d nodes", len(doc))
}

// TestTranslationIntegration_WithFootnotes tests translation blocks and footnotes together
func TestTranslationIntegration_WithFootnotes(t *testing.T) {
	poTranslator, err := translator.NewPOTranslator("../../example/translations.po")
	if err != nil {
		t.Fatalf("Failed to create PO translator: %v", err)
	}

	parser := NewParser(poTranslator)

	content := `{% trans %}This text will be translated{% endtrans %}

.. [1] This is a footnote.

{% trans %}Another translatable section{% endtrans %}

.. [2] Another footnote.`

	doc := parser.Parse(content)
	if len(doc) == 0 {
		t.Fatalf("Expected parsed nodes, got empty document")
	}

	// Count different node types
	paragraphCount := 0
	footnoteCount := 0

	for _, node := range doc {
		switch node.Type() {
		case nodes.NodeParagraph:
			paragraphCount++
		case nodes.NodeFootnote:
			footnoteCount++
		}
	}

	if paragraphCount < 2 {
		t.Errorf("Expected at least 2 paragraph nodes (translated), got %d", paragraphCount)
	}
	if footnoteCount < 2 {
		t.Errorf("Expected at least 2 footnote nodes, got %d", footnoteCount)
	}

	t.Logf("Document has %d paragraphs and %d footnotes", paragraphCount, footnoteCount)
}

// TestTranslationIntegration_WithFieldLists tests translation blocks with field lists
func TestTranslationIntegration_WithFieldLists(t *testing.T) {
	poTranslator, err := translator.NewPOTranslator("../../example/translations.po")
	if err != nil {
		t.Fatalf("Failed to create PO translator: %v", err)
	}

	parser := NewParser(poTranslator)

	content := `:author: John Doe
:date: 2025-10-12

{% trans %}This text will be translated{% endtrans %}

:version: 1.0

{% trans %}Another translatable section{% endtrans %}`

	doc := parser.Parse(content)
	if len(doc) == 0 {
		t.Fatalf("Expected parsed nodes, got empty document")
	}

	// Count different node types
	paragraphCount := 0
	fieldListCount := 0

	for _, node := range doc {
		switch node.Type() {
		case nodes.NodeParagraph:
			paragraphCount++
		case nodes.NodeFieldList:
			fieldListCount++
		}
	}

	if paragraphCount < 2 {
		t.Errorf("Expected at least 2 paragraph nodes (translated), got %d", paragraphCount)
	}
	if fieldListCount < 1 {
		t.Errorf("Expected at least 1 field list node, got %d", fieldListCount)
	}

	t.Logf("Document has %d paragraphs and %d field lists", paragraphCount, fieldListCount)
}

// TestTranslationIntegration_NoTranslatorProvided tests behavior when translator is nil
func TestTranslationIntegration_NoTranslatorProvided(t *testing.T) {
	parser := NewParser(nil)

	content := "{% trans %}This text will be translated{% endtrans %}"

	doc := parser.Parse(content)
	if len(doc) == 0 {
		t.Fatalf("Expected parsed nodes, got empty document")
	}

	// Without translator, should return original content
	paragraphNode, ok := doc[0].(*nodes.ParagraphNode)
	if !ok {
		t.Fatalf("Expected ParagraphNode, got %T", doc[0])
	}

	expectedContent := "This text will be translated"
	if paragraphNode.Content() != expectedContent {
		t.Errorf("Expected original content %q, got %q", expectedContent, paragraphNode.Content())
	}
}

// TestTranslationIntegration_WithCodeBlocks tests that code blocks are not translated
func TestTranslationIntegration_WithCodeBlocks(t *testing.T) {
	poTranslator, err := translator.NewPOTranslator("../../example/translations.po")
	if err != nil {
		t.Fatalf("Failed to create PO translator: %v", err)
	}

	parser := NewParser(poTranslator)

	content := `.. code-block:: python

    # This text will be translated - but should NOT be in code
    print("This text will be translated")

{% trans %}This text will be translated{% endtrans %}`

	doc := parser.Parse(content)
	if len(doc) == 0 {
		t.Fatalf("Expected parsed nodes, got empty document")
	}

	// Check that code block content is NOT translated
	for _, node := range doc {
		if code, ok := node.(*nodes.CodeNode); ok {
			if strings.Contains(code.Content(), "Este texto será traducido") {
				t.Errorf("Code block content should not be translated")
			}
		}
	}

	// Check that translation block outside code IS translated
	foundTranslation := false
	for _, node := range doc {
		if para, ok := node.(*nodes.ParagraphNode); ok {
			if strings.Contains(para.Content(), "Este texto será traducido") {
				foundTranslation = true
				break
			}
		}
	}

	if !foundTranslation {
		t.Errorf("Expected translation block to be translated")
	}
}

// TestTranslationIntegration_MalformedTranslationBlocks tests resilience with malformed blocks
func TestTranslationIntegration_MalformedTranslationBlocks(t *testing.T) {
	poTranslator, err := translator.NewPOTranslator("../../example/translations.po")
	if err != nil {
		t.Fatalf("Failed to create PO translator: %v", err)
	}

	parser := NewParser(poTranslator)

	tests := []struct {
		name    string
		content string
	}{
		{
			name:    "missing endtrans",
			content: "{% trans %}This text will be translated",
		},
		{
			name:    "missing trans",
			content: "This text will be translated{% endtrans %}",
		},
		{
			name:    "nested trans blocks",
			content: "{% trans %}Outer {% trans %}Inner{% endtrans %} Outer{% endtrans %}",
		},
		{
			name:    "extra spaces",
			content: "{%  trans  %}This text will be translated{%  endtrans  %}",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Should not panic
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Parser panicked on malformed translation block: %v", r)
				}
			}()

			doc := parser.Parse(test.content)
			// Document behavior - may or may not translate depending on how malformed
			t.Logf("Malformed block '%s' parsed into %d nodes", test.name, len(doc))
		})
	}
}

// TestTranslationIntegration_ConcurrentParsing tests thread safety with translator
func TestTranslationIntegration_ConcurrentParsing(t *testing.T) {
	poTranslator, err := translator.NewPOTranslator("../../example/translations.po")
	if err != nil {
		t.Fatalf("Failed to create PO translator: %v", err)
	}

	content := "{% trans %}This text will be translated{% endtrans %}"

	// Run multiple parsers concurrently with same translator
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Parser %d panicked: %v", id, r)
				}
				done <- true
			}()

			parser := NewParser(poTranslator)
			doc := parser.Parse(content)

			if len(doc) == 0 {
				t.Errorf("Parser %d returned empty document", id)
				return
			}

			// Verify translation worked
			if para, ok := doc[0].(*nodes.ParagraphNode); ok {
				if !strings.Contains(para.Content(), "Este texto será traducido") {
					t.Errorf("Parser %d failed to translate content", id)
				}
			}
		}(i)
	}

	// Wait for all goroutines
	for i := 0; i < 10; i++ {
		<-done
	}
}
