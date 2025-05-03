pkg/
├── nodes/                       # Node type definitions
│   ├── blockquote.go            # Defines BlockQuoteNode for representing block quotes in RST
│   ├── code.go                  # Defines CodeNode for representing code blocks in RST
│   ├── comment.go               # Defines CommentNode for representing comments in RST
│   ├── directive.go             # Defines DirectiveNode for representing RST directives
│   ├── doc.md                   # Documentation for the nodes package
│   ├── doctest.go               # Defines DoctestNode for representing doctest blocks
│   ├── em.go                    # Defines EmphasisNode for representing emphasized (italic) text
│   ├── extra_util.go            # Utility functions for node operations
│   ├── heading.go               # Defines HeadingNode for representing section headings
│   ├── lineblock.go             # Defines LineBlockNode for representing line blocks
│   ├── link.go                  # Defines LinkNode for representing hyperlinks
│   ├── list.go                  # Defines ListNode and ListItemNode for representing lists
│   ├── meta.go                  # Defines MetaNode for representing metadata information
│   ├── paragraph.go             # Defines ParagraphNode for representing text paragraphs
│   ├── strong.go                # Defines StrongNode for representing strong (bold) text
│   ├── subtitle.go              # Defines SubtitleNode for representing document subtitles
│   ├── table.go                 # Defines TableNode for representing table structures
│   ├── title.go                 # Defines TitleNode for representing document titles
│   ├── transition.go            # Defines TransitionNode for representing transitions between sections
│   └── types.go                 # Node type enumerations and base Node interface definitions
│
├── parser/                      # RST parsing logic
│   ├── blockquote.go            # Contains logic for parsing block quotes
│   ├── code.go                  # Contains logic for parsing code blocks
│   ├── context.go               # Manages parser context and state
│   ├── directive.go             # Contains logic for parsing RST directives
│   ├── doc.md                   # Documentation for the parser package
│   ├── doctest.go               # Contains logic for parsing doctest blocks
│   ├── emphasis.go              # Contains logic for parsing emphasized text
│   ├── headiing.go              # Contains logic for parsing section headings
│   ├── lexer.go                 # Tokenizes RST input into tokens
│   ├── lineblock.go             # Contains logic for parsing line blocks
│   ├── link.go                  # Contains logic for parsing hyperlinks
│   ├── list.go                  # Contains logic for parsing lists and list items
│   ├── meta.go                  # Contains logic for parsing metadata
│   ├── paragraph.go             # Contains logic for parsing text paragraphs
│   ├── parser.go                # Main parser implementation that processes tokens into a node tree
│   ├── parser_test.go           # Tests for the parser functionality
│   ├── patterns.go              # Regex patterns for RST syntax recognition
│   ├── strong.go                # Contains logic for parsing strong (bold) text
│   ├── subtitle.go              # Contains logic for parsing document subtitles
│   ├── table.go                 # Contains logic for parsing tables
│   └── title.go                 # Contains logic for parsing document titles
│
├── renderer/                    # Output rendering components
│   ├── doc.md                   # Documentation for the renderer package
│   ├── html.go                  # HTML output renderer implementation
│   ├── markdown.go              # Markdown output renderer implementation
│   └── pdf.go                   # PDF output renderer implementation
│
└── translator/                  # Translation capabilities
    ├── doc.md                   # Documentation for the translator package
    └── translator.go            # Handles translation of text content using PO files