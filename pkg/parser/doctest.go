package parser

import "strings"

func processDoctestBlock(line string) Token {
	// Handle doctest blocks
	if strings.HasPrefix(line, ">>> ") {
		return Token{
			Type:    TokenDoctest,
			Content: line[4:], // Remove the leading '>>> '
		}
	}
	return Token{}
}
