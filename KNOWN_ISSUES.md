# Known Issues

## Inline Formatting Not Supported in Mixed Content

**Status:** Known Limitation  
**Severity:** Moderate  
**Affects:** Parser architecture, all renderers (HTML, Markdown, PDF)

### Description

The parser does not support inline formatting (bold, italic) mixed with regular text on the same line. When a line contains both formatted and unformatted text, only the formatted portion is captured.

### Examples

**Input RST:**
```rst
A paragraph with **bold** and *italic* text.
```

**Current Output (HTML):**
```html
<strong>bold</strong>
```

**Expected Output (HTML):**
```html
<p>A paragraph with <strong>bold</strong> and <em>italic</em> text.</p>
```

### Root Cause

The lexer processes each line and returns the first token match (heading, strong, emphasis, etc.). When it finds `**bold**`, it:
1. Creates a Token with type=TokenStrong and content="bold"
2. Discards the surrounding text ("A paragraph with " and " and *italic* text.")
3. Parser creates a standalone StrongNode with no context

The architecture is line-oriented with single-token-per-line processing, not designed for inline mixed content.

### Impact

- **HTML Output:** Only formatted portions render, surrounding text is lost
- **Markdown Output:** Same issue, incomplete content
- **PDF Output:** Same issue
- **Use Cases Affected:** 
  - Documentation with inline emphasis
  - Mixed bold/italic in sentences
  - Links with surrounding text

### Workaround

Use formatting for entire lines/paragraphs rather than inline:

```rst
**This entire line is bold.**

This is a normal paragraph.

*This entire line is italic.*
```

### Proper Fix Requirements

A complete fix would require:

1. **Lexer Refactoring:**
   - Implement inline token scanning that finds ALL formatting in a line
   - Return array of tokens with positions for mixed content
   - Support nested/overlapping formatting

2. **Parser Enhancement:**
   - Process multiple tokens per line
   - Build paragraph nodes with inline children
   - Maintain proper text node structure

3. **Node Model Update:**
   - Support inline children in ParagraphNode
   - Create TextNode for plain text segments
   - Properly nest formatting nodes

4. **Renderer Updates:**
   - HTML: Walk inline children, integrate into paragraphs
   - Markdown: Reconstruct inline formatting in order
   - PDF: Handle inline formatting runs

5. **Testing:**
   - Edge cases: nested formatting, overlapping, malformed
   - Round-trip testing for all renderers
   - Performance with complex documents

### Estimated Effort

**High complexity:** 3-5 days of development + testing
- Architectural refactoring of core parser
- Risk of regressions in existing functionality
- Requires comprehensive test coverage

### Related

- CHECKLIST.md documents ~45% RST feature completion
- Project prioritizes practical web publishing over full RST compliance
- Other inline features (inline links, substitutions) likely have similar limitations

### Decision

**Deferred:** This represents a known limitation of the current architecture rather than a simple bug. The library was designed for practical web publishing with a subset of RST features. Full inline parsing support should be considered as a future enhancement requiring major version bump due to architectural changes.

**Date:** October 12, 2025  
**Documented by:** Copilot Agent during AUDIT.md gap analysis
