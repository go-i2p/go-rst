package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/go-i2p/go-rst/pkg/nodes"
	"github.com/go-i2p/go-rst/pkg/parser"
	"github.com/go-i2p/go-rst/pkg/renderer"
	"github.com/go-i2p/go-rst/pkg/translator"
)

// Configuration holds command line configuration
type Configuration struct {
	rstFile       string
	poFile        string
	outFileFormat string
	outFile       string
	debug         bool
}

func main() {
	config := parseCommandLineFlags()
	configureLogging(config.debug)
	validateInputFlags(config)

	content := readRSTFile(config.rstFile, config.debug)
	translator := initializeTranslator(config.poFile, config.debug)
	nodes := parseRSTContent(content, translator, config.debug)
	renderOutput(nodes, config)

	fmt.Printf("Successfully converted %s to %s\n", config.rstFile, config.outFile)
}

// parseCommandLineFlags extracts and validates command line flags
func parseCommandLineFlags() Configuration {
	rstFile := flag.String("rst", "", "Input RST file path")
	poFile := flag.String("po", "", "Input PO file path for translations")
	outFileFormat := flag.String("out-format", "html", "Output file format (html, pdf, markdown)")
	outFile := flag.String("out", "", "Output file path")
	debug := flag.Bool("debug", false, "Enable debug logging")
	flag.Parse()

	return Configuration{
		rstFile:       *rstFile,
		poFile:        *poFile,
		outFileFormat: *outFileFormat,
		outFile:       *outFile,
		debug:         *debug,
	}
}

// configureLogging sets up logging configuration based on debug flag
func configureLogging(debug bool) {
	if debug {
		log.SetFlags(log.Lshortfile | log.LstdFlags)
	}
}

// validateInputFlags checks required command line arguments
func validateInputFlags(config Configuration) {
	if config.rstFile == "" {
		log.Fatal("Please provide an input RST file using -rst flag")
	}
	if config.outFile == "" {
		log.Fatal("Please provide an output HTML file using -out flag")
	}
}

// readRSTFile reads and returns RST file content
func readRSTFile(rstFile string, debug bool) []byte {
	content, err := ioutil.ReadFile(rstFile)
	if err != nil {
		log.Fatalf("Failed to read RST file: %v", err)
	}

	if debug {
		log.Printf("Loaded RST file: %s", rstFile)
	}

	return content
}

// initializeTranslator creates and configures a translator instance
func initializeTranslator(poFile string, debug bool) translator.Translator {
	trans, err := translator.NewPOTranslator(poFile)
	if err != nil {
		log.Fatalf("Failed to initialize translator: %v", err)
	}

	if debug && poFile != "" {
		log.Printf("Loaded PO file: %s", poFile)
		testTranslation(trans, debug)
	}

	return trans
}

// testTranslation performs a test translation for debugging
func testTranslation(trans translator.Translator, debug bool) {
	testStr := "This text will be translated"
	translated := trans.Translate(testStr)
	log.Printf("Translation test: '%s' -> '%s'", testStr, translated)
}

// parseRSTContent parses RST content into nodes
func parseRSTContent(content []byte, trans translator.Translator, debug bool) []nodes.Node {
	p := parser.NewParser(trans)
	nodes := p.Parse(string(content))

	if debug {
		log.Printf("Parsed %d nodes", len(nodes))
	}

	return nodes
}

// renderOutput renders nodes to the specified output format
func renderOutput(nodes []nodes.Node, config Configuration) {
	switch config.outFileFormat {
	case "html":
		renderHTML(nodes, config.outFile)
	case "pdf":
		renderPDF(nodes, config.outFile)
	case "markdown":
		renderMarkdown(nodes, config.outFile)
	}
}

// renderHTML renders nodes to HTML format
func renderHTML(nodes []nodes.Node, outFile string) {
	r := renderer.NewHTMLRenderer()
	html := r.RenderPretty(nodes)
	WriteRendered(outFile, []byte(html))
}

// renderPDF renders nodes to PDF format
func renderPDF(nodes []nodes.Node, outFile string) {
	r := renderer.NewPDFRenderer()
	err := r.Render(nodes)
	if err != nil {
		log.Fatalf("Failed to render PDF: %v", err)
	}
	r.SaveToFile(outFile)
}

// renderMarkdown renders nodes to Markdown format
func renderMarkdown(nodes []nodes.Node, outFile string) {
	r := renderer.NewMarkdownRenderer()
	err := r.Render(nodes)
	if err != nil {
		log.Fatalf("Failed to render Markdown: %v", err)
	}
	WriteRendered(outFile, []byte(r.String()))
}

func WriteRendered(outFile string, doc []byte) {
	// Write output
	err := ioutil.WriteFile(outFile, []byte(doc), 0o644)
	if err != nil {
		log.Fatalf("Failed to write HTML file: %v", err)
	}
}
