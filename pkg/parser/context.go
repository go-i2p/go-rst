package parser

// ParserContext represents the current state of the parser.
type ParserContext struct {
	inMeta           bool
	inDirective      bool
	currentDirective string
	inCodeBlock      bool
	codeBlockIndent  int
	buffer           []string
}

// NewParserContext creates a new ParserContext instance.
func NewParserContext() *ParserContext {
	return &ParserContext{
		buffer: make([]string, 0),
	}
}

// Reset resets the parser context to its initial state.
func (c *ParserContext) Reset() {
	c.inMeta = false
	c.inDirective = false
	c.currentDirective = ""
	c.inCodeBlock = false
	c.codeBlockIndent = 0
	c.buffer = c.buffer[:0]
}
