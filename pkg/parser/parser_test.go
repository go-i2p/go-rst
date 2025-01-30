package parser

// test the functionality of the parser package using the test package
// Use example restructuredText files embedded in the test functions

import (
	"testing"

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
