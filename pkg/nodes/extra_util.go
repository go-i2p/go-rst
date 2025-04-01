package nodes

import "strings"

// GetIndentedContent Utility function to get node content with proper indentation
func GetIndentedContent(node Node) string {
	content := node.Content()
	if node.Level() > 0 {
		indent := strings.Repeat("    ", node.Level())
		lines := strings.Split(content, "\n")
		for i, line := range lines {
			lines[i] = indent + line
		}
		content = strings.Join(lines, "\n")
	}
	return content
}
