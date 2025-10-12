# go-rst

**WARNING: This repository contains significant amounts of LLM-generated code.**
**Use of LLMS has not been limited to review and documentation writing.**
**LLMs have written code here.**
**They were not allowed to proceed in unstructured ways, and we do understand the codebase.**

A Go library for parsing and rendering reStructuredText (RST) documents with translation support.
Supports only a subset of restructuredText for now, but relatively easy to expand compared to other attempts.
It is mostly unrelated to previous attempts to parse restructuredText in Go.

## Features

- RST to HTML conversion
- Translation support via PO files
- Clean and extensible API
- Pretty HTML output

## Installation

```bash
go get github.com/go-i2p/go-rst
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
    "os"
    
    "github.com/go-i2p/go-rst/pkg/parser"
    "github.com/go-i2p/go-rst/pkg/renderer"
    "github.com/go-i2p/go-rst/pkg/translator"
)

func main() {
    // Read RST content
    content, err := os.ReadFile("doc.rst")
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

For information about rst feature coverage see: [CHECKLIST.md](CHECKLIST.md)

## License

MIT License

## Credits

This project uses:
- [gotext](https://github.com/leonelquinteros/gotext) for PO file handling
- [gohtml](https://github.com/yosssi/gohtml) for HTML prettifying
- [gofpdf](github.com/jung-kurt/gofpdf) for PDF generation