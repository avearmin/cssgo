package cssgo

import "io"

// GlobalType is a concrete implementation representing the CSS global values.
// All properties accept GlobalType values.
// example: inherit, initial, unset, revert.
type GlobalType string

func (g GlobalType) RenderCSS(w io.Writer) error {
	_, err := w.Write([]byte(g))
	return err
}

func (g GlobalType) valueNode()          {}
func (g GlobalType) colorValue()         {}
func (g GlobalType) sizeValue()          {}
func (g GlobalType) displayValue()       {}
func (g GlobalType) borderStyleValue()   {}
func (g GlobalType) borderWidthValue()   {}
func (g GlobalType) flexDirectionValue() {}
func (g GlobalType) urlValue()           {}

const (
	Inherit GlobalType = "inherit"
	Initial GlobalType = "initial"
	Unset   GlobalType = "unset"
	Revert  GlobalType = "revert"
)

// AutoType represents the CSS `auto` value.
type AutoType string

func (a AutoType) RenderCSS(w io.Writer) error {
	_, err := w.Write([]byte(a))
	return err
}

func (a AutoType) valueNode() {}
func (a AutoType) sizeValue() {}

const Auto AutoType = "auto"
