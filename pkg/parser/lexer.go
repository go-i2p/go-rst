// pkg/parser/lexer.go

package parser

import (
	"strings"
)

// TokenType represents the type of a token.
type TokenType int

const (
	TokenText             TokenType = iota // TokenText represents a regular text token.
	TokenHeadingUnderline                  // TokenHeadingUnderline represents a heading underline token.
	TokenTransBlock                        // TokenTransBlock represents a transition block token.
	TokenMeta                              // TokenMeta represents a metadata token.
	TokenDirective                         // TokenDirective represents a directive token.
	TokenCodeBlock                         // TokenCodeBlock represents a code block token.
	TokenBlankLine                         // TokenBlankLine represents a blank line token.
	TokenIndent                            // TokenIndent represents an indent token.
	TokenBlockQuote                        // TokenBlockQuote represents a block quote token.
	TokenComment                           // TokenComment represents a comment token.
	TokenBulletList                        // TokenBulletList represents a bullet list item token.
	TokenEnumList                          // TokenEnumList represents an enumerated list item token.
	TokenDoctest                           // TokenDoctest represents a doctest token.
	TokenLineBlock                         // TokenLineBlock represents a line block token.
	TokenTransition                        // TokenTransition represents a transition token.
	TokenEmphasis                          // TokenEmphasis represents emphasized (italic) text
	TokenStrong                            // TokenStrong represents strong (bold) text
	TokenFootnote                          // TokenFootnote represents a footnote reference token.
	TokenDefinitionList                    // TokenDefinitionList represents a definition list token.
	TokenFieldList                         // TokenFieldList represents a field list token.
)

// Token represents a single token in the input text.
type Token struct {
	Type    TokenType
	Content string
	Args    []string
}

// Lexer represents a lexer for the input text.
type Lexer struct {
	patterns *Patterns
}

// NewLexer creates a new Lexer instance.
func NewLexer() *Lexer {
	return &Lexer{
		patterns: NewPatterns(),
	}
}

// Tokenize tokenizes a single line of input text.
// Returns TokenText as a safe fallback if any panic occurs during tokenization.
func (l *Lexer) Tokenize(line string) Token {
	// Add panic recovery to prevent crashes from malformed input
	defer func() {
		if r := recover(); r != nil {
			// Log the panic but don't crash - return safe fallback
			// In production, this could log to a proper logger
		}
	}()

	// Handle blank lines first
	if strings.TrimSpace(line) == "" {
		return Token{Type: TokenBlankLine}
	}

	// Normalize line by removing leading whitespace
	normalizedLine := l.normalizeLineIndentation(line)

	// Try directive-based tokens first
	if token := l.checkDirectiveTokens(normalizedLine); token.Type != TokenText {
		return token
	}

	// Try structural tokens
	if token := l.checkStructuralTokens(normalizedLine); token.Type != TokenText {
		return token
	}

	// Try list tokens
	if token := l.checkListTokens(normalizedLine); token.Type != TokenText {
		return token
	}

	// Try special element tokens (footnotes, definition lists, field lists)
	if token := l.checkSpecialElementTokens(normalizedLine); token.Type != TokenText {
		return token
	}

	// Try formatting tokens
	if token := l.checkFormattingTokens(normalizedLine); token.Type != TokenText {
		return token
	}

	// Default to regular text
	return Token{
		Type:    TokenText,
		Content: normalizedLine,
	}
}

// normalizeLineIndentation calculates indentation and returns the trimmed line.
func (l *Lexer) normalizeLineIndentation(line string) string {
	// Calculate indentation (though not currently used in token creation)
	indent := 0
	for _, r := range line {
		if r == ' ' {
			indent++
		} else if r == '\t' {
			indent += 4
		} else {
			break
		}
	}

	return strings.TrimLeft(line, " \t")
}

// checkDirectiveTokens checks for directive-related tokens including meta, code blocks, and custom directives.
func (l *Lexer) checkDirectiveTokens(line string) Token {
	// Check for meta directive
	if l.patterns.meta.MatchString(line) {
		return Token{Type: TokenMeta}
	}

	// Check for code block
	if l.patterns.codeBlock.MatchString(line) {
		args := parseDirectiveArgs(line)
		return Token{
			Type: TokenCodeBlock,
			Args: args,
		}
	}

	// Check for other directives
	if matches := l.patterns.directive.FindStringSubmatch(line); len(matches) > 1 {
		args := parseDirectiveArgs(line)
		return Token{
			Type:    TokenDirective,
			Content: safeExtractMatch(matches, 1),
			Args:    args,
		}
	}

	return Token{Type: TokenText}
}

// checkStructuralTokens checks for structural elements like headings, transitions, and translation blocks.
func (l *Lexer) checkStructuralTokens(line string) Token {
	// Check for basic structural elements first
	if token := l.checkBasicStructural(line); token.Type != TokenText {
		return token
	}

	// Check for content blocks
	if token := l.checkContentBlocks(line); token.Type != TokenText {
		return token
	}

	return Token{Type: TokenText}
}

// checkBasicStructural checks for basic structural elements like headings and transitions.
func (l *Lexer) checkBasicStructural(line string) Token {
	// Check for heading underline
	if l.patterns.headingUnderline.MatchString(line) {
		return Token{
			Type:    TokenHeadingUnderline,
			Content: line,
		}
	}

	// Check for translation blocks
	if matches := l.patterns.transBlock.FindStringSubmatch(line); len(matches) > 1 {
		return Token{
			Type:    TokenTransBlock,
			Content: safeExtractMatch(matches, 1),
		}
	}

	// Check for transitions
	if l.patterns.IsTransition(line) {
		transChar := l.patterns.TransitionChar(line)
		return Token{
			Type:    TokenTransition,
			Content: string(transChar),
		}
	}

	return Token{Type: TokenText}
}

// checkContentBlocks checks for content-containing blocks like quotes, comments, and line blocks.
func (l *Lexer) checkContentBlocks(line string) Token {
	// Check for footnote first (before comment check, since footnotes start with "..")
	if matches := l.patterns.footnote.FindStringSubmatch(line); len(matches) > 2 {
		return Token{
			Type:    TokenFootnote,
			Content: safeExtractMatch(matches, 2),           // footnote content
			Args:    []string{safeExtractMatch(matches, 1)}, // footnote label (number, #, or *)
		}
	}

	// Check for block quote
	if matches := l.patterns.blockQuote.FindStringSubmatch(line); len(matches) > 1 {
		attribution := safeExtractMatch(matches, 3)
		return Token{
			Type:    TokenBlockQuote,
			Content: safeExtractMatch(matches, 2),
			Args:    []string{attribution},
		}
	}

	// Check for comment
	if matches := l.patterns.comment.FindStringSubmatch(line); len(matches) > 1 {
		return Token{
			Type:    TokenComment,
			Content: safeExtractMatch(matches, 1),
		}
	}

	// Check for line block (poetry-style line with | prefix)
	if matches := l.patterns.lineBlock.FindStringSubmatch(line); len(matches) > 0 {
		return Token{
			Type:    TokenLineBlock,
			Content: strings.TrimSpace(safeExtractMatch(matches, 1)),
		}
	}

	return Token{Type: TokenText}
}

// checkListTokens checks for bullet and enumerated list items.
func (l *Lexer) checkListTokens(line string) Token {
	// Check for bullet list
	if matches := l.patterns.bulletList.FindStringSubmatch(line); len(matches) > 1 {
		return Token{
			Type:    TokenBulletList,
			Content: safeExtractMatch(matches, 4),
			Args:    safeExtractMatches(matches, 1, 2), // indent, bullet type
		}
	}

	// Check for enumerated list
	if matches := l.patterns.enumList.FindStringSubmatch(line); len(matches) > 1 {
		return Token{
			Type:    TokenEnumList,
			Content: safeExtractMatch(matches, 4),
			Args:    safeExtractMatches(matches, 1, 2), // indent, marker
		}
	}

	return Token{Type: TokenText}
}

// checkFormattingTokens checks for inline formatting like emphasis and strong text.
func (l *Lexer) checkFormattingTokens(line string) Token {
	// Check for strong (bold text) - must come before emphasis to avoid conflict
	if matches := l.patterns.strong.FindStringSubmatch(line); len(matches) > 1 {
		return Token{
			Type:    TokenStrong,
			Content: safeExtractMatch(matches, 1), // The text between double asterisks
		}
	}

	// Check for emphasis (italic text)
	if matches := l.patterns.emphasis.FindStringSubmatch(line); len(matches) > 1 {
		return Token{
			Type:    TokenEmphasis,
			Content: safeExtractMatch(matches, 1), // The text between asterisks
		}
	}

	return Token{Type: TokenText}
}

// checkSpecialElementTokens checks for definition lists and field lists.
func (l *Lexer) checkSpecialElementTokens(line string) Token {
	// Check for field list (metadata-style)
	if matches := l.patterns.fieldList.FindStringSubmatch(line); len(matches) > 2 {
		return Token{
			Type:    TokenFieldList,
			Content: safeExtractMatch(matches, 2),           // field value
			Args:    []string{safeExtractMatch(matches, 1)}, // field name
		}
	}

	// Note: Definition lists are complex multiline structures and need special handling
	// in the parser rather than simple line-by-line tokenization
	// For now, we'll handle them as text and let the parser deal with multiline patterns

	return Token{Type: TokenText}
}

func parseDirectiveArgs(line string) []string {
	parts := strings.SplitN(line, "::", 2)
	if len(parts) != 2 {
		return nil
	}

	args := strings.Fields(strings.TrimSpace(parts[1]))
	return args
}

// safeExtractMatch safely extracts a regex match at the given index with bounds checking.
// Returns empty string if index is out of bounds or matches is nil.
func safeExtractMatch(matches []string, index int) string {
	if matches == nil || index < 0 || index >= len(matches) {
		return ""
	}
	return matches[index]
}

// safeExtractMatches safely extracts multiple regex matches with bounds checking.
// Returns a slice with extracted values or empty strings for out-of-bounds indices.
func safeExtractMatches(matches []string, indices ...int) []string {
	result := make([]string, len(indices))
	for i, idx := range indices {
		result[i] = safeExtractMatch(matches, idx)
	}
	return result
}
