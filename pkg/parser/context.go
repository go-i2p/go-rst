package parser

type ParserContext struct {
    inMeta          bool
    inDirective     bool
    currentDirective string
    inCodeBlock     bool
    codeBlockIndent int
    buffer          []string
}

func NewParserContext() *ParserContext {
    return &ParserContext{
        buffer: make([]string, 0),
    }
}

func (c *ParserContext) Reset() {
    c.inMeta = false
    c.inDirective = false
    c.currentDirective = ""
    c.inCodeBlock = false
    c.codeBlockIndent = 0
    c.buffer = c.buffer[:0]
}