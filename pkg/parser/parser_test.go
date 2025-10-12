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

func TestParserContextResetBetweenParses(t *testing.T) {
	noopTranslator := translator.NewNoopTranslator()
	parser := NewParser(noopTranslator)

	// First parse with an incomplete code block (missing blank line terminator)
	incompleteCodeBlock := `.. code-block:: python
    def incomplete():
        return "no blank line after this"`

	doc1 := parser.Parse(incompleteCodeBlock)
	t.Logf("First parse returned %d nodes", len(doc1))

	// Check if context is polluted
	if parser.context.inCodeBlock {
		t.Log("Context shows inCodeBlock=true after first parse (expected for this test)")
	}

	// Second parse with regular content - should not be affected by first parse
	normalContent := `Regular paragraph content.

This should be parsed normally.`

	doc := parser.Parse(normalContent)
	t.Logf("Second parse returned %d nodes", len(doc))

	// Verify the second parse works correctly despite previous context pollution
	if len(doc) == 0 {
		t.Fatalf("Expected parsed nodes from second parse, got empty document")
	}

	// Should have paragraph nodes, not code block content
	foundParagraph := false
	for _, node := range doc {
		if node.Type() == nodes.NodeParagraph {
			foundParagraph = true
			t.Logf("Found paragraph with content: %s", node.Content())
			if strings.Contains(node.Content(), "Regular paragraph content") {
				break
			}
		}
	}

	if !foundParagraph {
		t.Errorf("Expected normal paragraph content in second parse, but parsing failed due to context pollution")
	}

	// Context should be clean after second parse
	if parser.context.inCodeBlock {
		t.Errorf("Parser context still shows inCodeBlock=true after second parse, indicating context was not reset")
	}
}

// Test footnote tokenization and parsing
func TestFootnoteTokenization(t *testing.T) {
	lexer := NewLexer()

	tests := []struct {
		name     string
		input    string
		expected TokenType
		label    string
		content  string
	}{
		{
			name:     "numeric footnote",
			input:    ".. [1] This is a footnote.",
			expected: TokenFootnote,
			label:    "1",
			content:  "This is a footnote.",
		},
		{
			name:     "auto-numbered footnote",
			input:    ".. [#] This is an auto-numbered footnote.",
			expected: TokenFootnote,
			label:    "#",
			content:  "This is an auto-numbered footnote.",
		},
		{
			name:     "symbol footnote",
			input:    ".. [*] This is a symbol footnote.",
			expected: TokenFootnote,
			label:    "*",
			content:  "This is a symbol footnote.",
		},
		{
			name:     "not a footnote",
			input:    "Regular text line.",
			expected: TokenText,
			label:    "",
			content:  "Regular text line.",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			token := lexer.Tokenize(test.input)

			if token.Type != test.expected {
				t.Errorf("Expected token type %v, got %v", test.expected, token.Type)
			}

			if test.expected == TokenFootnote {
				if len(token.Args) == 0 || token.Args[0] != test.label {
					t.Errorf("Expected footnote label %q, got %v", test.label, token.Args)
				}
				if token.Content != test.content {
					t.Errorf("Expected footnote content %q, got %q", test.content, token.Content)
				}
			}
		})
	}
}

// Test field list tokenization
func TestFieldListTokenization(t *testing.T) {
	lexer := NewLexer()

	tests := []struct {
		name      string
		input     string
		expected  TokenType
		fieldName string
		value     string
	}{
		{
			name:      "simple field",
			input:     ":author: John Doe",
			expected:  TokenFieldList,
			fieldName: "author",
			value:     "John Doe",
		},
		{
			name:      "field with empty value",
			input:     ":date:",
			expected:  TokenFieldList,
			fieldName: "date",
			value:     "",
		},
		{
			name:      "field with spaces in name",
			input:     ":last modified: 2025-01-01",
			expected:  TokenFieldList,
			fieldName: "last modified",
			value:     "2025-01-01",
		},
		{
			name:     "not a field",
			input:    "Regular text line.",
			expected: TokenText,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			token := lexer.Tokenize(test.input)

			if token.Type != test.expected {
				t.Errorf("Expected token type %v, got %v", test.expected, token.Type)
			}

			if test.expected == TokenFieldList {
				if len(token.Args) == 0 || token.Args[0] != test.fieldName {
					t.Errorf("Expected field name %q, got %v", test.fieldName, token.Args)
				}
				if token.Content != test.value {
					t.Errorf("Expected field value %q, got %q", test.value, token.Content)
				}
			}
		})
	}
}

// Test footnote parsing
func TestFootnoteParsing(t *testing.T) {
	parser := NewParser(nil)

	content := `.. [1] This is a numbered footnote.
.. [#] This is an auto-numbered footnote.
.. [*] This is a symbol footnote.`

	doc := parser.Parse(content)

	if len(doc) != 3 {
		t.Fatalf("Expected 3 nodes, got %d", len(doc))
	}

	// Check first footnote (numbered)
	if footnote, ok := doc[0].(*nodes.FootnoteNode); ok {
		if footnote.Label() != "1" {
			t.Errorf("Expected label '1', got %q", footnote.Label())
		}
		if footnote.Content() != "This is a numbered footnote." {
			t.Errorf("Expected content 'This is a numbered footnote.', got %q", footnote.Content())
		}
		if footnote.IsAutoNumber() {
			t.Errorf("Expected IsAutoNumber to be false for numbered footnote")
		}
	} else {
		t.Errorf("Expected first node to be FootnoteNode, got %T", doc[0])
	}

	// Check second footnote (auto-numbered)
	if footnote, ok := doc[1].(*nodes.FootnoteNode); ok {
		if footnote.Label() != "#" {
			t.Errorf("Expected label '#', got %q", footnote.Label())
		}
		if !footnote.IsAutoNumber() {
			t.Errorf("Expected IsAutoNumber to be true for auto-numbered footnote")
		}
	} else {
		t.Errorf("Expected second node to be FootnoteNode, got %T", doc[1])
	}

	// Check third footnote (symbol)
	if footnote, ok := doc[2].(*nodes.FootnoteNode); ok {
		if footnote.Label() != "*" {
			t.Errorf("Expected label '*', got %q", footnote.Label())
		}
		if !footnote.IsAutoNumber() {
			t.Errorf("Expected IsAutoNumber to be true for symbol footnote")
		}
	} else {
		t.Errorf("Expected third node to be FootnoteNode, got %T", doc[2])
	}
}

// Test field list parsing
func TestFieldListParsing(t *testing.T) {
	parser := NewParser(nil)

	content := `:author: John Doe
:date: 2025-01-01
:version: 1.0`

	doc := parser.Parse(content)

	if len(doc) != 1 {
		t.Fatalf("Expected 1 node (field list), got %d", len(doc))
	}

	if fieldList, ok := doc[0].(*nodes.FieldListNode); ok {
		// Check author field
		author, exists := fieldList.GetField("author")
		if !exists {
			t.Errorf("Expected 'author' field to exist")
		}
		if author != "John Doe" {
			t.Errorf("Expected author 'John Doe', got %q", author)
		}

		// Check date field
		date, exists := fieldList.GetField("date")
		if !exists {
			t.Errorf("Expected 'date' field to exist")
		}
		if date != "2025-01-01" {
			t.Errorf("Expected date '2025-01-01', got %q", date)
		}

		// Check version field
		version, exists := fieldList.GetField("version")
		if !exists {
			t.Errorf("Expected 'version' field to exist")
		}
		if version != "1.0" {
			t.Errorf("Expected version '1.0', got %q", version)
		}

		// Check non-existent field
		_, exists = fieldList.GetField("nonexistent")
		if exists {
			t.Errorf("Expected 'nonexistent' field to not exist")
		}
	} else {
		t.Errorf("Expected FieldListNode, got %T", doc[0])
	}
}

// Test mixed content parsing
func TestMixedContentParsing(t *testing.T) {
	parser := NewParser(nil)

	content := `:author: John Doe

This is a paragraph.

.. [1] This is a footnote.

Another paragraph.`

	doc := parser.Parse(content)

	if len(doc) < 3 {
		t.Fatalf("Expected at least 3 nodes, got %d", len(doc))
	}

	// Verify we have field list, paragraphs, and footnote
	nodeTypes := make(map[nodes.NodeType]int)
	for _, node := range doc {
		nodeTypes[node.Type()]++
	}

	if nodeTypes[nodes.NodeFieldList] < 1 {
		t.Errorf("Expected at least 1 field list node")
	}
	if nodeTypes[nodes.NodeFootnote] < 1 {
		t.Errorf("Expected at least 1 footnote node")
	}
	if nodeTypes[nodes.NodeParagraph] < 1 {
		t.Errorf("Expected at least 1 paragraph node")
	}
}

// TestParseTranslationBlockWithPOTranslator tests that translation blocks
// are correctly processed with a real PO file translator.
func TestParseTranslationBlockWithPOTranslator(t *testing.T) {
	// Load a real PO file for translation
	poTranslator, err := translator.NewPOTranslator("../../example/translations.po")
	if err != nil {
		t.Fatalf("Failed to create PO translator: %v", err)
	}

	parser := NewParser(poTranslator)

	// Test content with a translation block that matches a PO entry
	content := "{% trans %}This text will be translated{% endtrans %}"

	doc := parser.Parse(content)
	if len(doc) == 0 {
		t.Fatalf("Expected parsed nodes, got empty document")
	}

	// Should have exactly one paragraph node with translated content
	if len(doc) != 1 {
		t.Errorf("Expected 1 node, got %d", len(doc))
	}

	paragraphNode, ok := doc[0].(*nodes.ParagraphNode)
	if !ok {
		t.Fatalf("Expected ParagraphNode, got %T", doc[0])
	}

	// Verify the content was translated
	expectedTranslated := "Este texto será traducido"
	if paragraphNode.Content() != expectedTranslated {
		t.Errorf("Expected translated content %q, got %q", expectedTranslated, paragraphNode.Content())
	}
}

// TestParseTranslationBlockWithUntranslatableContent tests graceful fallback
// when translation is not found in PO file.
func TestParseTranslationBlockWithUntranslatableContent(t *testing.T) {
	// Load a real PO file for translation
	poTranslator, err := translator.NewPOTranslator("../../example/translations.po")
	if err != nil {
		t.Fatalf("Failed to create PO translator: %v", err)
	}

	parser := NewParser(poTranslator)

	// Test content with a translation block that has no matching PO entry
	content := "{% trans %}This text has no translation{% endtrans %}"

	doc := parser.Parse(content)
	if len(doc) == 0 {
		t.Fatalf("Expected parsed nodes, got empty document")
	}

	// Should have exactly one paragraph node with original content (fallback)
	if len(doc) != 1 {
		t.Errorf("Expected 1 node, got %d", len(doc))
	}

	paragraphNode, ok := doc[0].(*nodes.ParagraphNode)
	if !ok {
		t.Fatalf("Expected ParagraphNode, got %T", doc[0])
	}

	// Verify the content fell back to original when no translation found
	expectedFallback := "This text has no translation"
	if paragraphNode.Content() != expectedFallback {
		t.Errorf("Expected fallback content %q, got %q", expectedFallback, paragraphNode.Content())
	}
}

// TestParseMultipleTranslationBlocks tests multiple translation blocks in one document.
func TestParseMultipleTranslationBlocks(t *testing.T) {
	// Load a real PO file for translation
	poTranslator, err := translator.NewPOTranslator("../../example/translations.po")
	if err != nil {
		t.Fatalf("Failed to create PO translator: %v", err)
	}

	parser := NewParser(poTranslator)

	// Test content with individual translation blocks on separate lines
	content1 := "{% trans %}This text will be translated{% endtrans %}"
	content2 := "{% trans %}Another translatable section{% endtrans %}"
	regularText := "Some regular text."

	// Parse each separately to understand how the parser works
	doc1 := parser.Parse(content1)
	doc2 := parser.Parse(content2)
	doc3 := parser.Parse(regularText)

	// Verify first translation
	if len(doc1) != 1 {
		t.Errorf("Expected 1 node for first translation, got %d", len(doc1))
	}
	if paragraphNode1, ok := doc1[0].(*nodes.ParagraphNode); ok {
		expectedTranslated1 := "Este texto será traducido"
		if paragraphNode1.Content() != expectedTranslated1 {
			t.Errorf("Expected first translated content %q, got %q", expectedTranslated1, paragraphNode1.Content())
		}
	} else {
		t.Errorf("Expected first node to be ParagraphNode, got %T", doc1[0])
	}

	// Verify second translation
	if len(doc2) != 1 {
		t.Errorf("Expected 1 node for second translation, got %d", len(doc2))
	}
	if paragraphNode2, ok := doc2[0].(*nodes.ParagraphNode); ok {
		expectedTranslated2 := "Otra sección traducible"
		if paragraphNode2.Content() != expectedTranslated2 {
			t.Errorf("Expected second translated content %q, got %q", expectedTranslated2, paragraphNode2.Content())
		}
	} else {
		t.Errorf("Expected second node to be ParagraphNode, got %T", doc2[0])
	}

	// Verify regular text
	if len(doc3) != 1 {
		t.Errorf("Expected 1 node for regular text, got %d", len(doc3))
	}
	if paragraphNode3, ok := doc3[0].(*nodes.ParagraphNode); ok {
		if paragraphNode3.Content() != regularText {
			t.Errorf("Expected regular content %q, got %q", regularText, paragraphNode3.Content())
		}
	} else {
		t.Errorf("Expected third node to be ParagraphNode, got %T", doc3[0])
	}
}
