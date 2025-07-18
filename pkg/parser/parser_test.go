package parser

// test the functionality of the parser package using the test package
// Use example restructuredText files embedded in the test functions

import (
	"strings"
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

func TestParseCodeBlock(t *testing.T) {
	noopTranslator := translator.NewNoopTranslator()
	parser := NewParser(noopTranslator)
	content := `.. code-block:: python

    def hello():
        print("Hello, world!")
        return True

End of test.`

	doc := parser.Parse(content)
	if len(doc) == 0 {
		t.Errorf("Expected parsed nodes, got empty document")
	}

	// Check that we have a code node
	foundCode := false
	for _, node := range doc {
		if node.Type() == nodes.NodeCode {
			foundCode = true
			codeNode := node.(*nodes.CodeNode)
			if codeNode.Language() != "python" {
				t.Errorf("Expected code language to be 'python', got '%s'", codeNode.Language())
			}
			if !strings.Contains(codeNode.Content(), "def hello():") {
				t.Errorf("Expected code content to contain 'def hello():', got '%s'", codeNode.Content())
			}
			if !strings.Contains(codeNode.Content(), "print(\"Hello, world!\")") {
				t.Errorf("Expected code content to contain print statement, got '%s'", codeNode.Content())
			}
		}
	}

	if !foundCode {
		t.Errorf("Expected to find a code node in parsed document")
	}
}
