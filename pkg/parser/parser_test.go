package parser

// test the functionality of the parser package using the test package
// Use example restructuredText files embedded in the test functions

import (
	"testing"

	"github.com/go-i2p/go-rst/pkg/nodes"
	"github.com/go-i2p/go-rst/pkg/translator"
)

const (
	simpleDoc  = "example/doc.rst"
	complexDoc = "example/complexDoc.rst"
)

func TestParse(t *testing.T) {
	noopTranslator := translator.NewNoopTranslator()
	parser := NewParser(noopTranslator)
	doc := parser.Parse(simpleDoc)
	if doc == nil {
		t.Errorf("Expected a document, got nil")
	}
}

func TestParseTwo(t *testing.T) {
	noopTranslator := translator.NewNoopTranslator()
	parser := NewParser(noopTranslator)
	doc := parser.Parse(complexDoc)
	if doc == nil {
		t.Errorf("Expected a document, got nil")
	}
}

func TestParseEmpty(t *testing.T) {
	noopTranslator := translator.NewNoopTranslator()
	parser := NewParser(noopTranslator)
	doc := parser.Parse("")
	if len(doc) > 0 {
		t.Errorf("Expected empty, got a document")
	}
}

func TestParseNilTranslatorEmptyInput(t *testing.T) {
	parser := NewParser(nil)
	doc := parser.Parse("")
	if len(doc) > 0 {
		t.Errorf("Expected empty, got a document")
	}
}

func TestParseNilTranslatorWithTranslationBlocks(t *testing.T) {
	parser := NewParser(nil)
	content := "{% trans %}Hello, world!{% endtrans %}"

	// This should not panic, but gracefully handle the nil translator
	doc := parser.Parse(content)
	if len(doc) == 0 {
		t.Errorf("Expected parsed nodes, got empty document")
	}
}

func TestParseStrongText(t *testing.T) {
	noopTranslator := translator.NewNoopTranslator()
	parser := NewParser(noopTranslator)
	content := "This is **bold text** in a sentence."

	doc := parser.Parse(content)
	if len(doc) == 0 {
		t.Errorf("Expected parsed nodes, got empty document")
	}

	// Check that we have a strong node
	foundStrong := false
	for _, node := range doc {
		if node.Type() == nodes.NodeStrong {
			foundStrong = true
			if node.Content() != "bold text" {
				t.Errorf("Expected strong content to be 'bold text', got '%s'", node.Content())
			}
		}
	}

	if !foundStrong {
		t.Errorf("Expected to find a strong node in parsed document")
	}
}
