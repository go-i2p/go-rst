// pkg/parser/patterns.go

package parser

import "regexp"

// Patterns holds compiled regular expressions for parsing Markdown syntax.
type Patterns struct {
	headingUnderline *regexp.Regexp
	transBlock       *regexp.Regexp
	meta             *regexp.Regexp
	directive        *regexp.Regexp
	codeBlock        *regexp.Regexp
}

// NewPatterns initializes and returns a new instance of Patterns with compiled regular expressions.
func NewPatterns() *Patterns {
	return &Patterns{
		headingUnderline: regexp.MustCompile(`^[=\-~]+$`),
		transBlock:       regexp.MustCompile(`{%\s*trans\s*%}(.*?){%\s*endtrans\s*%}`),
		meta:             regexp.MustCompile(`^\.\.\s+meta::`),
		directive:        regexp.MustCompile(`^\.\.\s+(\w+)::`),
		codeBlock:        regexp.MustCompile(`^\.\.\s+code::`),
	}
}
