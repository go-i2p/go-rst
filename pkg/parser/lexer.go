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
func (l *Lexer) Tokenize(line string) Token {
	// Handle blank lines
	if strings.TrimSpace(line) == "" {
		return Token{Type: TokenBlankLine}
	}

	// Calculate indentation
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

	line = strings.TrimLeft(line, " \t")

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
			Content: matches[1],
		}
	}

	// Check for meta directive
	if l.patterns.meta.MatchString(line) {
		return Token{
			Type: TokenMeta,
		}
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
			Content: matches[1],
			Args:    args,
		}
	}

	// Check for block quote
	if matches := l.patterns.blockQuote.FindStringSubmatch(line); len(matches) > 1 {
		attribution := ""
		if len(matches) > 2 {
			attribution = matches[3]
		}
		return Token{
			Type:    TokenBlockQuote,
			Content: matches[2],
			Args:    []string{attribution},
		}
	}

	// Check for comment
	if matches := l.patterns.comment.FindStringSubmatch(line); len(matches) > 1 {
		return Token{
			Type:    TokenComment,
			Content: matches[1],
		}
	}
	// Check for bullet list
	if matches := l.patterns.bulletList.FindStringSubmatch(line); len(matches) > 1 {
		return Token{
			Type:    TokenBulletList,
			Content: matches[4],
			Args:    []string{matches[1], matches[2]}, // indent, bullet type
		}
	}

	// Check for enumerated list
	if matches := l.patterns.enumList.FindStringSubmatch(line); len(matches) > 1 {
		return Token{
			Type:    TokenEnumList,
			Content: matches[4],
			Args:    []string{matches[1], matches[2]}, // indent, marker
		}
	}

	// Check for line block (poetry-style line with | prefix)
	if matches := l.patterns.lineBlock.FindStringSubmatch(line); len(matches) > 0 {
		return Token{
			Type:    TokenLineBlock,
			Content: strings.TrimSpace(matches[1]), // The content after the | character
		}
	}

	// Regular text
	return Token{
		Type:    TokenText,
		Content: line,
	}
}

func parseDirectiveArgs(line string) []string {
	parts := strings.SplitN(line, "::", 2)
	if len(parts) != 2 {
		return nil
	}

	args := strings.Fields(strings.TrimSpace(parts[1]))
	return args
}
