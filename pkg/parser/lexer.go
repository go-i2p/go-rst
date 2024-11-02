// pkg/parser/lexer.go

package parser

import (
	"strings"
)

type TokenType int

const (
	TokenText TokenType = iota
	TokenHeadingUnderline
	TokenTransBlock
	TokenMeta
	TokenDirective
	TokenCodeBlock
	TokenBlankLine
	TokenIndent
)

type Token struct {
	Type    TokenType
	Content string
	Args    []string
}

type Lexer struct {
	patterns *Patterns
}

func NewLexer() *Lexer {
	return &Lexer{
		patterns: NewPatterns(),
	}
}

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