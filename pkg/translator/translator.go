package translator

import (
	"github.com/leonelquinteros/gotext"
)

type Translator interface {
    Translate(text string) string
}

type POTranslator struct {
    po *gotext.Po
}

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

func NewNoopTranslator() *NoopTranslator {
    return &NoopTranslator{}
}

func (t *NoopTranslator) Translate(text string) string {
    return text
}