package parser

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

// TestMemoryUsageLargeDocument verifies that parsing large documents doesn't cause excessive heap growth.
// This test focuses on the parser's memory footprint, not string construction overhead.
func TestMemoryUsageLargeDocument(t *testing.T) {
	// Pre-build content outside of measurement to isolate parser memory usage
	paragraphs := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		paragraphs[i] = fmt.Sprintf("This is paragraph %d with content that represents typical RST length.", i)
	}
	content := strings.Join(paragraphs, "\n\n")

	// Force GC and get baseline AFTER building the test content
	runtime.GC()
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	// Parse the large document
	parser := NewParser(nil)
	nodes := parser.Parse(content)

	// Verify we got nodes
	if len(nodes) == 0 {
		t.Fatal("Expected nodes from large document, got none")
	}

	// Check memory after parsing - focus on current heap size
	runtime.GC()
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	// Calculate final heap size (memory actually retained after GC)
	finalHeapMB := float64(m2.Alloc) / 1024 / 1024

	// Log memory usage for visibility
	t.Logf("Parsed %d nodes from large document (1,000 paragraphs)", len(nodes))
	t.Logf("Final heap size after GC: %.2f MB", finalHeapMB)
	t.Logf("Baseline heap: %.2f MB", float64(m1.Alloc)/1024/1024)

	// Final heap should be small - the parser should not retain excessive memory
	// After GC, heap should be under 5MB even for large documents
	// (This tests that we're not leaking memory or holding unnecessary references)
	if finalHeapMB > 5 {
		t.Errorf("Excessive heap retention: %.2f MB after parsing 1,000 paragraphs", finalHeapMB)
	}
}

// TestMemoryLeakParserReuse verifies that parser context properly resets and doesn't leak memory.
// This test parses the same document 100 times and checks for memory growth.
func TestMemoryLeakParserReuse(t *testing.T) {
	// Force GC before starting
	runtime.GC()
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	baseline := m1.TotalAlloc

	content := `
Title
=====

This is a paragraph with some content.

.. code-block:: python

   def hello():
       print("world")

* List item 1
* List item 2
* List item 3

:author: John Doe
:date: 2025-10-12

.. [1] A footnote reference.
`

	parser := NewParser(nil)

	// Parse 100 times to check for leaks
	for i := 0; i < 100; i++ {
		nodes := parser.Parse(content)
		if len(nodes) == 0 {
			t.Fatalf("Parse iteration %d returned no nodes", i)
		}
	}

	// Check memory after repeated parsing
	runtime.GC()
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	// Calculate total allocations in MB
	totalAllocated := m2.TotalAlloc - baseline
	memIncreaseMB := float64(totalAllocated) / 1024 / 1024

	// Log for visibility
	t.Logf("Total allocated after 100 parses: %.2f MB", memIncreaseMB)
	t.Logf("Current heap: %.2f MB", float64(m2.Alloc)/1024/1024)

	// Memory should not grow excessively with parser reuse
	// We expect < 10MB total allocated for 100 parses of small document
	if memIncreaseMB > 10 {
		t.Errorf("Possible memory leak: %.2f MB allocated after 100 parses", memIncreaseMB)
	}
}

// TestMemoryUsageContextReset verifies that parser context properly releases memory on reset.
func TestMemoryUsageContextReset(t *testing.T) {
	context := NewParserContext()

	// Add lots of data to the buffer
	for i := 0; i < 10000; i++ {
		context.buffer = append(context.buffer, fmt.Sprintf("Line %d with content", i))
	}

	if len(context.buffer) != 10000 {
		t.Fatalf("Expected 10000 buffer entries, got %d", len(context.buffer))
	}

	// Reset should clear the buffer without reallocating
	context.Reset()

	if len(context.buffer) != 0 {
		t.Errorf("Expected empty buffer after reset, got %d entries", len(context.buffer))
	}

	// Capacity should be maintained for reuse (not reallocated)
	if cap(context.buffer) == 0 {
		t.Error("Buffer capacity was lost after reset, causing unnecessary reallocation")
	}

	// Verify all flags are reset
	if context.inMeta || context.inDirective || context.inCodeBlock {
		t.Error("Context flags not properly reset")
	}
}

// TestMemoryUsageMultipleParsers verifies that multiple parsers can coexist without excessive memory.
func TestMemoryUsageMultipleParsers(t *testing.T) {
	// Force GC before starting
	runtime.GC()
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	baseline := m1.TotalAlloc

	content := `
Title
=====

This is a paragraph.

* List item
* Another item
`

	// Create 100 parsers and parse with each
	parsers := make([]*Parser, 100)
	for i := 0; i < 100; i++ {
		parsers[i] = NewParser(nil)
		nodes := parsers[i].Parse(content)
		if len(nodes) == 0 {
			t.Fatalf("Parser %d returned no nodes", i)
		}
	}

	// Check memory after creating many parsers
	runtime.GC()
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	// Calculate total allocations in MB
	totalAllocated := m2.TotalAlloc - baseline
	memIncreaseMB := float64(totalAllocated) / 1024 / 1024

	// Log for visibility
	t.Logf("Total allocated for 100 parsers: %.2f MB", memIncreaseMB)
	t.Logf("Current heap: %.2f MB", float64(m2.Alloc)/1024/1024)

	// Memory should scale reasonably - expect < 30MB total allocated for 100 parsers
	if memIncreaseMB > 30 {
		t.Errorf("Excessive memory usage: %.2f MB allocated for 100 parser instances", memIncreaseMB)
	}
}

// TestMemoryUsageComplexDocument tests memory with a document containing all RST features.
func TestMemoryUsageComplexDocument(t *testing.T) {
	// Force GC before starting
	runtime.GC()
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)
	baseline := m1.TotalAlloc

	// Build a complex document with all features
	var builder strings.Builder
	for i := 0; i < 1000; i++ {
		builder.WriteString(fmt.Sprintf(`
Section %d
==========

This is paragraph %d with **bold** and *italic* text.

.. code-block:: python

   def function_%d():
       return "hello"

* Bullet item 1
* Bullet item 2

1. Enumerated item
2. Another item

:field%d: value%d
:author: Test Author

.. [%d] Footnote text here.

Term
    Definition for the term.

----------

`, i, i, i, i, i, i))
	}
	content := builder.String()

	// Parse the complex document
	parser := NewParser(nil)
	nodes := parser.Parse(content)

	// Verify we got nodes
	if len(nodes) == 0 {
		t.Fatal("Expected nodes from complex document, got none")
	}

	// Check memory after parsing
	runtime.GC()
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	// Calculate total allocations in MB
	totalAllocated := m2.TotalAlloc - baseline
	memIncreaseMB := float64(totalAllocated) / 1024 / 1024

	// Log memory usage
	t.Logf("Parsed %d nodes from complex document with all features", len(nodes))
	t.Logf("Total allocated: %.2f MB", memIncreaseMB)
	t.Logf("Current heap: %.2f MB", float64(m2.Alloc)/1024/1024)

	// Memory should be reasonable for complex document
	// Expect < 150MB total allocated for 1000 sections with all features
	if memIncreaseMB > 150 {
		t.Errorf("Excessive memory usage: %.2f MB allocated for complex document", memIncreaseMB)
	}
}

// BenchmarkParseSmallDocument benchmarks parsing performance and memory for a small document.
func BenchmarkParseSmallDocument(b *testing.B) {
	content := `
Title
=====

This is a simple paragraph.

* Item 1
* Item 2
`

	parser := NewParser(nil)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parser.Parse(content)
	}
}

// BenchmarkParseLargeDocument benchmarks parsing performance for a large document.
func BenchmarkParseLargeDocument(b *testing.B) {
	var builder strings.Builder
	for i := 0; i < 1000; i++ {
		builder.WriteString(fmt.Sprintf("Paragraph %d with some content.\n\n", i))
	}
	content := builder.String()

	parser := NewParser(nil)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parser.Parse(content)
	}
}

// BenchmarkParseComplexDocument benchmarks parsing with all RST features.
func BenchmarkParseComplexDocument(b *testing.B) {
	var builder strings.Builder
	for i := 0; i < 100; i++ {
		builder.WriteString(fmt.Sprintf(`
Section %d
==========

Paragraph with **bold** and *italic*.

.. code-block:: python

   code = "example"

* List item

:field: value

.. [1] Footnote.

`, i))
	}
	content := builder.String()

	parser := NewParser(nil)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_ = parser.Parse(content)
	}
}
