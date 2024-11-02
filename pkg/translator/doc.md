# translator
--
    import "i2pgit.org/idk/go-rst/pkg/translator"


## Usage

#### type NoopTranslator

```go
type NoopTranslator struct{}
```

NoopTranslator implements Translator interface but doesn't translate

#### func  NewNoopTranslator

```go
func NewNoopTranslator() *NoopTranslator
```
NewNoopTranslator returns a new NoopTranslator

#### func (*NoopTranslator) Translate

```go
func (t *NoopTranslator) Translate(text string) string
```
Translate returns the same text it receives(NoopTranslator)

#### type POTranslator

```go
type POTranslator struct {
}
```

POTranslator implements Translator interface using a PO file

#### func  NewPOTranslator

```go
func NewPOTranslator(poFile string) (*POTranslator, error)
```
NewPOTranslator returns a new POTranslator

#### func (*POTranslator) Translate

```go
func (t *POTranslator) Translate(text string) string
```
Translate returns the translated text if it exists in the PO file, otherwise it
returns the original text

#### type Translator

```go
type Translator interface {
	Translate(text string) string
}
```

Translator is an interface for translating text
