# Contributing to go-rst

## Adding New Node Types

### 1. Define the Node Type
First, add your new node type in `pkg/nodes/types.go`:

```go
const (
    NodeText NodeType = iota
    NodeHeading
    NodeParagraph
    NodeMeta
    NodeDirective
    NodeCode
    YourNewNodeType    // Add your new node type here
)
```

### 2. Create Node Structure
In `pkg/nodes/nodes.go`, define your new node structure:

```go
type YourNewNode struct {
    BaseNode
    // Add specific fields for your node type
    SpecificField string
}

func NewYourNewNode(specific string) Node {
    return &YourNewNode{
        BaseNode: BaseNode{
            nodeType: YourNewNodeType,
        },
        SpecificField: specific,
    }
}
```

### 3. Update Parser Components

#### a. Add Pattern (if needed)
In `pkg/parser/patterns.go`, add a new regex pattern:

```go
type Patterns struct {
    // Existing patterns...
    yourNewPattern *regexp.Regexp
}

func NewPatterns() *Patterns {
    return &Patterns{
        // Existing patterns...
        yourNewPattern: regexp.MustCompile(`your-pattern-here`),
    }
}
```

#### b. Add Token Type
In `pkg/parser/lexer.go`, add a new token type:

```go
const (
    // Existing token types...
    TokenYourNew TokenType = iota
)
```

#### c. Update Lexer
Modify the `Tokenize` method in `pkg/parser/lexer.go` to recognize your new pattern:

```go
func (l *Lexer) Tokenize(line string) Token {
    // Existing checks...
    
    if l.patterns.yourNewPattern.MatchString(line) {
        return Token{
            Type: TokenYourNew,
            Content: // Extract relevant content
        }
    }
    
    // Default case...
}
```

#### d. Update Parser
In `pkg/parser/parser.go`, add handling for your new token type:

```go
func (p *Parser) processToken(token, prevToken Token, currentNode nodes.Node) nodes.Node {
    switch token.Type {
    // Existing cases...
    case TokenYourNew:
        return nodes.NewYourNewNode(token.Content)
    }
    return currentNode
}
```

### 4. Add Renderer Support
In `pkg/renderer/html.go`, implement HTML rendering for your new node type:

```go
func (r *HTMLRenderer) renderNode(node nodes.Node) string {
    switch node.Type() {
    // Existing cases...
    case nodes.YourNewNodeType:
        return r.renderYourNewNode(node.(*nodes.YourNewNode))
    }
    return ""
}

func (r *HTMLRenderer) renderYourNewNode(node *nodes.YourNewNode) string {
    return fmt.Sprintf("<your-tag>%s</your-tag>", node.SpecificField)
}
```

### 5. Testing

1. Add test cases in appropriate test files
2. Create example RST documents showcasing your new node type
3. Verify translation handling if applicable

### Best Practices

1. **Type Safety**: Use Go's type system effectively
2. **Error Handling**: Return meaningful errors
3. **Documentation**: Add godoc comments to all public types and functions
4. **Testing**: Write comprehensive tests for new functionality
5. **Performance**: Consider memory allocation and processing efficiency

### Translation Support

If your node type contains translatable content:

1. Ensure the content is passed through the translator in the parser
2. Update example PO files to demonstrate translation capabilities
3. Add translation tests

### Pull Request Process

1. Fork the repository
2. Create a feature branch
3. Implement your changes
4. Add tests and documentation
5. Submit a pull request with a clear description

### Questions?

Feel free to open an issue for discussion before implementing major changes.