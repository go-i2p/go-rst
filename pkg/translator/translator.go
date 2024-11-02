package translator

import (
	"github.com/leonelquinteros/gotext"
)

// Translator is an interface for translating text
type Translator interface {
	Translate(text string) string
}

// POTranslator implements Translator interface using a PO file
type POTranslator struct {
	po *gotext.Po
}

// NewPOTranslator returns a new POTranslator
func NewPOTranslator(poFile string) (*POTranslator, error) {
	translator := &POTranslator{
		po: gotext.NewPo(),
	}

	// If no PO file is provided, return a pass-through translator
	if poFile == "" {
		return translator, nil
	}

	// Parse PO file
	translator.po.ParseFile(poFile)

	return translator, nil
}

// Translate returns the translated text if it exists in the PO file, otherwise it returns the original text
func (t *POTranslator) Translate(text string) string {
	if t.po == nil {
		return text
	}
	translated := t.po.Get(text)
	if translated == "" {
		return text
	}
	return translated
}

// NoopTranslator implements Translator interface but doesn't translate
type NoopTranslator struct{}

// NewNoopTranslator returns a new NoopTranslator
func NewNoopTranslator() *NoopTranslator {
	return &NoopTranslator{}
}

// Translate returns the same text it receives(NoopTranslator)
func (t *NoopTranslator) Translate(text string) string {
	return text
}
