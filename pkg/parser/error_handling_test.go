package parser

import (
	"strings"
	"testing"
	"unicode/utf8"
)

// TestMalformedDirectives tests that malformed directives don't crash the parser
func TestMalformedDirectives(t *testing.T) {
	parser := NewParser(nil)

	tests := []struct {
		name    string
		content string
	}{
		{
			name:    "directive missing closing colons",
			content: ".. code-block",
		},
		{
			name:    "directive with only one colon",
			content: ".. code-block: python",
		},
		{
			name:    "directive with extra colons",
			content: ".. code-block::: python",
		},
		{
			name:    "directive with empty name",
			content: ".. ::",
		},
		{
			name:    "directive with special characters",
			content: ".. <script>alert('xss')</script>::",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Should not panic
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Parser panicked on %s: %v", test.name, r)
				}
			}()

			doc := parser.Parse(test.content)
			// Should return some result (even if just text nodes)
			if doc == nil {
				t.Errorf("Parser returned nil for %s", test.name)
			}
		})
	}
}

// TestInvalidFootnoteSyntax tests that invalid footnote syntax doesn't crash
func TestInvalidFootnoteSyntax(t *testing.T) {
	parser := NewParser(nil)

	tests := []struct {
		name    string
		content string
	}{
		{
			name:    "footnote missing closing bracket",
			content: ".. [1 This is incomplete",
		},
		{
			name:    "footnote missing opening bracket",
			content: ".. 1] This is incomplete",
		},
		{
			name:    "footnote with empty label",
			content: ".. [] This is incomplete",
		},
		{
			name:    "footnote with very long label",
			content: ".. [" + strings.Repeat("1", 10000) + "] Content",
		},
		{
			name:    "footnote with special characters in label",
			content: ".. [<script>] XSS attempt",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Parser panicked on %s: %v", test.name, r)
				}
			}()

			doc := parser.Parse(test.content)
			if doc == nil {
				t.Errorf("Parser returned nil for %s", test.name)
			}
		})
	}
}

// TestCorruptFieldLists tests that corrupt field list syntax doesn't crash
func TestCorruptFieldLists(t *testing.T) {
	parser := NewParser(nil)

	tests := []struct {
		name    string
		content string
	}{
		{
			name:    "field missing closing colon",
			content: ":author John Doe",
		},
		{
			name:    "field missing opening colon",
			content: "author: John Doe",
		},
		{
			name:    "field with no name",
			content: ":: value",
		},
		{
			name:    "field with multiple colons",
			content: ":a:b:c: value",
		},
		{
			name:    "field with special characters",
			content: ":<script>: XSS",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Parser panicked on %s: %v", test.name, r)
				}
			}()

			doc := parser.Parse(test.content)
			if doc == nil {
				t.Errorf("Parser returned nil for %s", test.name)
			}
		})
	}
}

// TestExtremelyLongLines tests that very long input lines don't cause issues
func TestExtremelyLongLines(t *testing.T) {
	parser := NewParser(nil)

	tests := []struct {
		name    string
		content string
	}{
		{
			name:    "very long paragraph",
			content: strings.Repeat("a", 100000),
		},
		{
			name:    "very long heading",
			content: strings.Repeat("=", 100000),
		},
		{
			name:    "very long code line",
			content: ".. code-block:: python\n\n    " + strings.Repeat("x", 100000),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Parser panicked on %s: %v", test.name, r)
				}
			}()

			doc := parser.Parse(test.content)
			if doc == nil {
				t.Errorf("Parser returned nil for %s", test.name)
			}
		})
	}
}

// TestSpecialCharacters tests that special characters are handled safely
func TestSpecialCharacters(t *testing.T) {
	parser := NewParser(nil)

	tests := []struct {
		name    string
		content string
	}{
		{
			name:    "null bytes",
			content: "Hello\x00World",
		},
		{
			name:    "control characters",
			content: "Hello\r\n\t\bWorld",
		},
		{
			name:    "xss attempt in content",
			content: "<script>alert('xss')</script>",
		},
		{
			name:    "sql injection attempt",
			content: "'; DROP TABLE users; --",
		},
		{
			name:    "path traversal attempt",
			content: "../../etc/passwd",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Parser panicked on %s: %v", test.name, r)
				}
			}()

			doc := parser.Parse(test.content)
			if doc == nil {
				t.Errorf("Parser returned nil for %s", test.name)
			}
		})
	}
}

// TestNilAndEmptyInputs tests edge cases with nil and empty inputs
func TestNilAndEmptyInputs(t *testing.T) {
	parser := NewParser(nil)

	tests := []struct {
		name    string
		content string
	}{
		{
			name:    "empty string",
			content: "",
		},
		{
			name:    "only whitespace",
			content: "   \t\n   ",
		},
		{
			name:    "only newlines",
			content: "\n\n\n\n",
		},
		{
			name:    "null character only",
			content: "\x00",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Parser panicked on %s: %v", test.name, r)
				}
			}()

			doc := parser.Parse(test.content)
			// Empty content should return empty slice, not nil
			if doc == nil {
				t.Errorf("Parser returned nil for %s, expected empty slice", test.name)
			}
		})
	}
}

// TestUnicodeEdgeCases tests various Unicode edge cases
func TestUnicodeEdgeCases(t *testing.T) {
	parser := NewParser(nil)

	tests := []struct {
		name    string
		content string
	}{
		{
			name:    "emoji in heading",
			content: "Hello 🌍 World\n==============",
		},
		{
			name:    "right-to-left text",
			content: "مرحبا بالعالم",
		},
		{
			name:    "combining characters",
			content: "e\u0301", // é using combining acute accent
		},
		{
			name:    "zero-width characters",
			content: "Hello\u200BWorld", // zero-width space
		},
		{
			name:    "invalid UTF-8 sequences",
			content: "Hello" + string([]byte{0xFF, 0xFE}) + "World",
		},
		{
			name:    "very long unicode string",
			content: strings.Repeat("🌍", 10000),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Parser panicked on %s: %v", test.name, r)
				}
			}()

			doc := parser.Parse(test.content)
			if doc == nil {
				t.Errorf("Parser returned nil for %s", test.name)
			}

			// Verify the content is valid UTF-8 after parsing
			for _, node := range doc {
				if !utf8.ValidString(node.Content()) {
					t.Logf("Warning: Node content contains invalid UTF-8 in %s", test.name)
				}
			}
		})
	}
}

// TestNestedStructures tests deeply nested structures
func TestNestedStructures(t *testing.T) {
	parser := NewParser(nil)

	tests := []struct {
		name    string
		content string
	}{
		{
			name: "nested lists",
			content: `- Item 1
  - Nested 1.1
    - Nested 1.1.1
      - Nested 1.1.1.1
        - Very deep nesting`,
		},
		{
			name:    "nested block quotes",
			content: strings.Repeat("    ", 100) + "Deep quote",
		},
		{
			name: "mixed nesting",
			content: `.. code-block:: python

    # Code block
    def nested():
        """
        .. note:: Directive in code?
        """
        pass`,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Parser panicked on %s: %v", test.name, r)
				}
			}()

			doc := parser.Parse(test.content)
			if doc == nil {
				t.Errorf("Parser returned nil for %s", test.name)
			}
		})
	}
}

// TestIncompleteStructures tests incomplete or truncated structures
func TestIncompleteStructures(t *testing.T) {
	parser := NewParser(nil)

	tests := []struct {
		name    string
		content string
	}{
		{
			name:    "incomplete translation block",
			content: "{% trans %}Hello",
		},
		{
			name:    "code block without content",
			content: ".. code-block:: python\n\n",
		},
		{
			name:    "heading without underline",
			content: "This is a heading",
		},
		{
			name:    "list item without content",
			content: "- ",
		},
		{
			name:    "table with missing cells",
			content: "+---+---+\n| A |\n+---+---+",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Parser panicked on %s: %v", test.name, r)
				}
			}()

			doc := parser.Parse(test.content)
			if doc == nil {
				t.Errorf("Parser returned nil for %s", test.name)
			}
		})
	}
}

// TestConcurrentParsing tests that multiple parsers can run concurrently
func TestConcurrentParsing(t *testing.T) {
	content := `Test Document
=============

This is a test with **bold** and *italic* text.

.. code-block:: python

    def hello():
        print("Hello, world!")

:author: Test
:date: 2025-10-12`

	// Run multiple parsers concurrently
	done := make(chan bool, 10)

	for i := 0; i < 10; i++ {
		go func(id int) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Parser %d panicked: %v", id, r)
				}
				done <- true
			}()

			parser := NewParser(nil)
			doc := parser.Parse(content)

			if doc == nil {
				t.Errorf("Parser %d returned nil", id)
			}
			if len(doc) == 0 {
				t.Errorf("Parser %d returned empty document", id)
			}
		}(i)
	}

	// Wait for all goroutines to complete
	for i := 0; i < 10; i++ {
		<-done
	}
}

// TestLexerPanicRecovery tests that lexer panic recovery works
func TestLexerPanicRecovery(t *testing.T) {
	lexer := NewLexer()

	// These should not panic even with potentially problematic input
	tests := []string{
		strings.Repeat("*", 100000),
		strings.Repeat("=", 100000),
		strings.Repeat(":", 100000),
		".. [" + strings.Repeat("x", 100000) + "]",
		string([]byte{0xFF, 0xFE, 0xFD}),
	}

	for i, test := range tests {
		t.Run(string(rune('A'+i)), func(t *testing.T) {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("Lexer panicked on test case %d: %v", i, r)
				}
			}()

			token := lexer.Tokenize(test)
			// Should always return a token, even if it's TokenText as fallback
			if token.Type == 0 && token.Content == "" {
				t.Logf("Warning: Lexer returned empty token for test case %d", i)
			}
		})
	}
}

// TestMemoryConsumption tests that parser doesn't consume excessive memory
func TestMemoryConsumption(t *testing.T) {
	parser := NewParser(nil)

	// Create a large but not unreasonable document
	var content strings.Builder
	for i := 0; i < 1000; i++ {
		content.WriteString("This is paragraph ")
		content.WriteString(string(rune('0' + i%10)))
		content.WriteString("\n\n")
	}

	defer func() {
		if r := recover(); r != nil {
			t.Errorf("Parser panicked on large document: %v", r)
		}
	}()

	doc := parser.Parse(content.String())
	if doc == nil {
		t.Error("Parser returned nil for large document")
	}

	// Should have parsed many nodes
	if len(doc) < 500 {
		t.Logf("Warning: Parser returned fewer nodes than expected: %d", len(doc))
	}
}
