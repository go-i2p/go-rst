package nodes

import (
	"testing"
)

func TestFootnoteNode(t *testing.T) {
	t.Run("NewFootnoteNode creates valid node", func(t *testing.T) {
		label := "1"
		content := "This is a footnote"
		autoNumber := false

		node := NewFootnoteNode(label, content, autoNumber)

		if node.Type() != NodeFootnote {
			t.Errorf("Expected NodeFootnote, got %v", node.Type())
		}
		if node.Content() != content {
			t.Errorf("Expected content %q, got %q", content, node.Content())
		}
		if node.Label() != label {
			t.Errorf("Expected label %q, got %q", label, node.Label())
		}
		if node.IsAutoNumber() != autoNumber {
			t.Errorf("Expected autoNumber %v, got %v", autoNumber, node.IsAutoNumber())
		}
	})

	t.Run("Auto-numbered footnote", func(t *testing.T) {
		node := NewFootnoteNode("#", "Auto footnote", true)

		if !node.IsAutoNumber() {
			t.Error("Expected auto-numbered footnote")
		}
		if node.Label() != "#" {
			t.Errorf("Expected label #, got %q", node.Label())
		}
	})

	t.Run("Symbol footnote", func(t *testing.T) {
		node := NewFootnoteNode("*", "Symbol footnote", false)

		if node.IsAutoNumber() {
			t.Error("Expected non-auto-numbered footnote")
		}
		if node.Label() != "*" {
			t.Errorf("Expected label *, got %q", node.Label())
		}
	})
}

func TestDefinitionListNode(t *testing.T) {
	t.Run("NewDefinitionListNode creates valid node", func(t *testing.T) {
		node := NewDefinitionListNode()

		if node.Type() != NodeDefinitionList {
			t.Errorf("Expected NodeDefinitionList, got %v", node.Type())
		}
		if node.Count() != 0 {
			t.Errorf("Expected count 0, got %d", node.Count())
		}
	})

	t.Run("AddDefinition adds term and definition", func(t *testing.T) {
		node := NewDefinitionListNode()
		term := "Term 1"
		definition := "Definition of term 1"

		node.AddDefinition(term, definition)

		if node.Count() != 1 {
			t.Errorf("Expected count 1, got %d", node.Count())
		}

		terms := node.Terms()
		definitions := node.Definitions()

		if len(terms) != 1 || terms[0] != term {
			t.Errorf("Expected terms [%q], got %v", term, terms)
		}
		if len(definitions) != 1 || definitions[0] != definition {
			t.Errorf("Expected definitions [%q], got %v", definition, definitions)
		}
	})

	t.Run("Multiple definitions", func(t *testing.T) {
		node := NewDefinitionListNode()

		node.AddDefinition("Term 1", "Definition 1")
		node.AddDefinition("Term 2", "Definition 2")

		if node.Count() != 2 {
			t.Errorf("Expected count 2, got %d", node.Count())
		}

		terms := node.Terms()
		definitions := node.Definitions()

		if len(terms) != 2 || terms[1] != "Term 2" {
			t.Errorf("Expected second term 'Term 2', got %v", terms)
		}
		if len(definitions) != 2 || definitions[1] != "Definition 2" {
			t.Errorf("Expected second definition 'Definition 2', got %v", definitions)
		}
	})

	t.Run("Terms and Definitions return copies", func(t *testing.T) {
		node := NewDefinitionListNode()
		node.AddDefinition("Original", "Original def")

		terms := node.Terms()
		definitions := node.Definitions()

		// Modify returned slices
		terms[0] = "Modified"
		definitions[0] = "Modified def"

		// Original should be unchanged
		originalTerms := node.Terms()
		originalDefs := node.Definitions()

		if originalTerms[0] != "Original" {
			t.Error("Terms() should return a copy")
		}
		if originalDefs[0] != "Original def" {
			t.Error("Definitions() should return a copy")
		}
	})
}

func TestFieldListNode(t *testing.T) {
	t.Run("NewFieldListNode creates valid node", func(t *testing.T) {
		node := NewFieldListNode()

		if node.Type() != NodeFieldList {
			t.Errorf("Expected NodeFieldList, got %v", node.Type())
		}
		if node.HasFields() {
			t.Error("Expected no fields initially")
		}
	})

	t.Run("AddField and GetField work correctly", func(t *testing.T) {
		node := NewFieldListNode()
		key := "Author"
		value := "John Doe"

		node.AddField(key, value)

		if !node.HasFields() {
			t.Error("Expected to have fields after adding")
		}

		retrievedValue, exists := node.GetField(key)
		if !exists {
			t.Error("Expected field to exist")
		}
		if retrievedValue != value {
			t.Errorf("Expected value %q, got %q", value, retrievedValue)
		}
	})

	t.Run("GetField returns false for non-existent field", func(t *testing.T) {
		node := NewFieldListNode()

		_, exists := node.GetField("NonExistent")
		if exists {
			t.Error("Expected non-existent field to return false")
		}
	})

	t.Run("Fields returns copy of map", func(t *testing.T) {
		node := NewFieldListNode()
		node.AddField("Key1", "Value1")
		node.AddField("Key2", "Value2")

		fields := node.Fields()

		// Modify returned map
		fields["Key1"] = "Modified"
		fields["NewKey"] = "NewValue"

		// Original should be unchanged
		originalValue, _ := node.GetField("Key1")
		if originalValue != "Value1" {
			t.Error("Fields() should return a copy")
		}

		_, hasNewKey := node.GetField("NewKey")
		if hasNewKey {
			t.Error("Modifications to returned map should not affect original")
		}
	})

	t.Run("Multiple fields", func(t *testing.T) {
		node := NewFieldListNode()

		node.AddField("Author", "John Doe")
		node.AddField("Date", "2025-01-01")
		node.AddField("Version", "1.0")

		fields := node.Fields()
		if len(fields) != 3 {
			t.Errorf("Expected 3 fields, got %d", len(fields))
		}

		expectedFields := map[string]string{
			"Author":  "John Doe",
			"Date":    "2025-01-01",
			"Version": "1.0",
		}

		for key, expectedValue := range expectedFields {
			if value, exists := fields[key]; !exists || value != expectedValue {
				t.Errorf("Expected field %q = %q, got %q (exists: %v)", key, expectedValue, value, exists)
			}
		}
	})
}
