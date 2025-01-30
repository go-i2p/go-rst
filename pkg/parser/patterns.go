// pkg/parser/patterns.go

package parser

import (
	"regexp"
)

// Patterns holds compiled regular expressions for parsing Markdown syntax.
type Patterns struct {
	headingUnderline *regexp.Regexp
	transBlock       *regexp.Regexp
	meta             *regexp.Regexp
	directive        *regexp.Regexp
	codeBlock        *regexp.Regexp
	blockQuote       *regexp.Regexp
	doctest          *regexp.Regexp
	lineBlock        *regexp.Regexp
	comment          *regexp.Regexp
	title            *regexp.Regexp
	subtitle         *regexp.Regexp
	transition       *regexp.Regexp
}

// NewPatterns initializes and returns a new instance of Patterns with compiled regular expressions.
func NewPatterns() *Patterns {
	return &Patterns{
		headingUnderline: regexp.MustCompile(`^[=\-~]+$`),
		transBlock:       regexp.MustCompile(`{%\s*trans\s*%}(.*?){%\s*endtrans\s*%}`),
		meta:             regexp.MustCompile(`^\.\.\s+meta::`),
		directive:        regexp.MustCompile(`^\.\.\s+(\w+)::`),
		codeBlock:        regexp.MustCompile(`^\.\.\s+code::`),
		blockQuote:       regexp.MustCompile(`^(\s{4,})(.*?)(?:\s*--\s*(.*))?$`),
		doctest:          regexp.MustCompile(`^>>> (.+)\n((?:[^>].*\n)*)`),
		lineBlock:        regexp.MustCompile(`^\|(.*)$`),
		comment:          regexp.MustCompile(`^\.\.\s(.*)$`),
		title:            regexp.MustCompile(`^(={3,}|~{3,})\n(.+?)\n(?:={3,}|~{3,})$`),
		subtitle:         regexp.MustCompile(`^(-{3,})\n(.+?)\n(?:-{3,})$`),
		transition:       regexp.MustCompile(`^(\-{4,}|\={4,}|\*{4,})$`),
	}
}
