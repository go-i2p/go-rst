package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"i2pgit.org/idk/go-rst/pkg/parser"
	"i2pgit.org/idk/go-rst/pkg/renderer"
	"i2pgit.org/idk/go-rst/pkg/translator"
)

func main() {
    // CLI flags
    rstFile := flag.String("rst", "", "Input RST file path")
    poFile := flag.String("po", "", "Input PO file path for translations")
    outFile := flag.String("out", "", "Output HTML file path")
    debug := flag.Bool("debug", false, "Enable debug logging")
    flag.Parse()

    if *debug {
        log.SetFlags(log.Lshortfile | log.LstdFlags)
    }

    // Validate input flags
    if *rstFile == "" {
        log.Fatal("Please provide an input RST file using -rst flag")
    }
    if *outFile == "" {
        log.Fatal("Please provide an output HTML file using -out flag")
    }

    // Read RST content
    content, err := ioutil.ReadFile(*rstFile)
    if err != nil {
        log.Fatalf("Failed to read RST file: %v", err)
    }

    if *debug {
        log.Printf("Loaded RST file: %s", *rstFile)
    }

    // Initialize translator
    trans, err := translator.NewPOTranslator(*poFile)
    if err != nil {
        log.Fatalf("Failed to initialize translator: %v", err)
    }

    if *debug && *poFile != "" {
        log.Printf("Loaded PO file: %s", *poFile)
        // Test translation
        testStr := "This text will be translated"
        translated := trans.Translate(testStr)
        log.Printf("Translation test: '%s' -> '%s'", testStr, translated)
    }

    // Initialize parser with translator
    p := parser.NewParser(trans)

    // Parse RST content
    nodes := p.Parse(string(content))

    if *debug {
        log.Printf("Parsed %d nodes", len(nodes))
    }

    // Initialize HTML renderer
    r := renderer.NewHTMLRenderer()

    // Render HTML
    html := r.Render(nodes)

    // Write output
    err = ioutil.WriteFile(*outFile, []byte(html), 0644)
    if err != nil {
        log.Fatalf("Failed to write HTML file: %v", err)
    }

    fmt.Printf("Successfully converted %s to %s\n", *rstFile, *outFile)
}