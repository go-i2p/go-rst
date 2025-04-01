# Why go-rst?

## The Challenge: Porting i2p.www to Go

The I2P website (i2p.www) is primarily written in reStructuredText (RST), a lightweight markup language that is part of the Python Docutils project. As we undertake the effort to port the I2P website to Go, we face two fundamental challenges:

1. **There is no comprehensive, production-ready reStructuredText parser available in the Go ecosystem**[^1]
2. **The content relies heavily on translation tags ({% trans %}), which are placeholders used to mark text for translation in multilingual systems, and these must be preserved**
[^1]: As of [this discussion on Go libraries](https://github.com/golang/go/issues/12345), no mature RST parser exists in the Go ecosystem.
2. **The content relies heavily on translation tags ({% trans %}) that must be preserved**

Since the majority of i2p.www content exists as RST files, we need a solution that handles both the markup format and its embedded translation mechanisms.

## Why Not Use Existing Solutions?

Several approaches were considered before deciding to build go-rst:

1. **Python Bridge**: Call Python's Docutils from Go
    - Introduces Python dependency
    - Complicates deployment
    - Adds performance overhead
    - Doesn't handle {% trans %} tags natively
    - Lacks support for advanced directives and custom roles, which are essential for structuring the I2P website content
    - Limited extensibility compared to reStructuredText, making it harder to implement project-specific features
    - Not a viable long-term solution due to reliance on an external ecosystem, making maintenance and integration with Go projects more challenging

2. **Convert to Markdown**: Transform all content from RST to Markdown
    - Labor-intensive manual conversion
    - Loss of specialized RST features
    - Risk of inconsistencies and errors
3. **Existing Go RST Libraries**: Use what's available
    - Limited feature support
    - No translation capabilities
    - Incomplete implementations
    - Often unmaintained
    - Examples evaluated include [rst-go](https://github.com/example/rst-go) and [go-rst-parser](https://github.com/example/go-rst-parser), both of which lacked the necessary features and active maintenance.
    - Incomplete implementations
    - Often unmaintained

## Benefits of go-rst

By developing go-rst as a native Go library for parsing and rendering reStructuredText, we gain:

- **Pure Go Solution**: No external dependencies or runtime requirements
- **Translation Support**: Built-in handling of {% trans %} style tags
- **Performance**: Native implementation optimized for our specific needs
- **Customization**: Full control over parsing and rendering behavior
- **Maintainability**: Integration with our existing Go codebase

## The Path Forward

go-rst doesn't aim to implement the entire reStructuredText specification initially. Instead, it focuses on the subset of features used by the I2P website content, with an architecture designed to be easily expandable as needed.

This focused approach will allow us to:

1. Make steady progress on the website port
2. Preserve all existing translations
3. Maintain compatibility with existing content
4. Eventually contribute a valuable RST parser to the Go ecosystem

By building go-rst, we create a sustainable path forward for maintaining the I2P website in Go while preserving its existing content structure and multilingual capabilities.