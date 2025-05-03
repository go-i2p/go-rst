package parser

import (
	"strings"

	"github.com/go-i2p/go-rst/pkg/nodes"
)

// IsTransition checks if a line is a transition
func (p *Patterns) IsTransition(line string) bool {
	// A transition is a line with 4+ repeated punctuation characters
	return len(strings.TrimSpace(line)) >= 4 && p.transition.MatchString(strings.TrimSpace(line))
}

// TransitionChar extracts the character used in the transition
func (p *Patterns) TransitionChar(line string) rune {
	trimmed := strings.TrimSpace(line)
	if len(trimmed) > 0 {
		return rune(trimmed[0])
	}
	return '-' // Default to hyphen if empty (shouldn't happen)
}

// processTransition handles the parsing of transition sections
// A transition is a horizontal line separator typically used between sections
func (p *Parser) processTransition(content string) *nodes.TransitionNode {
	// Extract the character used in the transition
	var transChar rune = '-' // Default to hyphen
	if len(content) > 0 {
		transChar = rune(content[0])
	}

	// Create a new transition node with the character
	return nodes.NewTransitionNode(transChar)
}
