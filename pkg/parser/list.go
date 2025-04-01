package parser

import "github.com/go-i2p/go-rst/pkg/nodes"

func (p *Parser) processListItem(content string, args []string, isOrdered bool, currentNode nodes.Node) nodes.Node {
	indent := 0
	if len(args) > 0 {
		// Calculate indent from the whitespace prefix
		indent = len(args[0])
	}

	/*marker := ""
	if len(args) > 1 {
		marker = args[1]
	}*/

	// Create a new list item node
	listItem := nodes.NewListItemNode(content)

	// If we don't have a current node or it's not a list, create a new list
	if currentNode == nil || currentNode.Type() != nodes.NodeList {
		listNode := nodes.NewListNode(isOrdered)
		listNode.AppendChild(listItem)
		return listNode
	}

	// If we have a list but with different type or indent, create a new list
	listNode := currentNode.(*nodes.ListNode)
	if listNode.IsOrdered() != isOrdered || listNode.Indent() != indent {
		newList := nodes.NewListNode(isOrdered)
		newList.SetIndent(indent)
		newList.AppendChild(listItem)
		return newList
	}

	// Add item to existing list
	listNode.AppendChild(listItem)
	return listNode
}
