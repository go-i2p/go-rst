# go-rst

A Go library for parsing and rendering reStructuredText (RST) documents with translation support.
Supports only a subset of restructuredText for now, but relatively easy to expand compared to other attempts.

## Features

- RST to HTML conversion
- Translation support via PO files
- Clean and extensible API
- Pretty HTML output

## Installation

```bash
go get i2pgit.org/idk/go-rst
```

## Quick Start

### Command Line Usage

```bash
go-rst -rst example/doc.rst -po example/translations.po -out output.html
```

### Library Usage

```go
package main

import (
    "fmt"
    "io/ioutil"
    
    "i2pgit.org/idk/go-rst/pkg/parser"
    "i2pgit.org/idk/go-rst/pkg/renderer"
    "i2pgit.org/idk/go-rst/pkg/translator"
)

func main() {
    // Read RST content
    content, err := ioutil.ReadFile("doc.rst")
    if err != nil {
        panic(err)
    }

    // Initialize translator with PO file (optional)
    trans, err := translator.NewPOTranslator("translations.po")
    if err != nil {
        panic(err)
    }

    // Create parser with translator
    p := parser.NewParser(trans)

    // Parse RST content
    nodes := p.Parse(string(content))

    // Create HTML renderer
    r := renderer.NewHTMLRenderer()

    // Render to HTML
    html := r.RenderPretty(nodes)

    // Save or use the HTML
    fmt.Println(html)
}
```

## Documentation

For more detailed information about adding new node types or contributing to the project, see [CONTRIBUTING.md](CONTRIBUTING.md).

## License

MIT License

## Credits

This project uses:
- [gotext](https://github.com/leonelquinteros/gotext) for PO file handling
- [gohtml](https://github.com/yosssi/gohtml) for HTML prettifying