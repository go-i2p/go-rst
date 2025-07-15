package parser

import (
	"strings"

	"github.com/go-i2p/go-rst/pkg/nodes"
)

func (p *Parser) processCodeBlock(line string, currentNode nodes.Node) nodes.Node {
	if currentNode == nil || currentNode.Type() != nodes.NodeCode {
		return currentNode
	}

	// Check if we're at the end of the code block
	// RST code blocks end when we encounter a line that is not indented
	// or when we hit a blank line followed by unindented content
	trimmedLine := strings.TrimSpace(line)

	// If this is a blank line, add it to buffer but continue
	if trimmedLine == "" {
		p.context.buffer = append(p.context.buffer, line)
		return currentNode
	}

	// Check if line is properly indented (at least 4 spaces for code content)
	if len(line) >= 4 && line[:4] == "    " {
		// This is indented content, add to buffer
		// Remove the base indentation (4 spaces) to preserve relative indentation
		p.context.buffer = append(p.context.buffer, line[4:])
		return currentNode
	}
	// Line is not indented, this means end of code block
	// Finalize the code block
	codeNode := currentNode.(*nodes.CodeNode)
	content := strings.Join(p.context.buffer, "\n")
	codeNode.SetContent(strings.TrimSpace(content))

	// Add the completed code block to nodes
	p.nodes = append(p.nodes, codeNode)

	// Reset context and return nil to signal completion
	p.context.Reset()
	return nil
}
