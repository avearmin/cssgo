package gomour

import "io"

// BorderStyleValue defines an interface for CSS-compatible border style values.
// This ensures that any type implementing this interface can be rendered as a valid CSS border style.
type BorderStyleValue interface {
	ValueNode
	borderStyleValue()
}

// BorderStyle represents a CSS border style, such as "solid", "dotted", or "dashed".
// It is a concrete type that implements the BorderStyleValue interface.
type BorderStyle string

// Predefined border style values as per the CSS specification.
const (
	Solid  BorderStyle = "solid"
	Dotted BorderStyle = "dotted"
	Dashed BorderStyle = "dashed"
	Double BorderStyle = "double"
	Groove BorderStyle = "groove"
	Ridge  BorderStyle = "ridge"
	Inset  BorderStyle = "inset"
	Outset BorderStyle = "outset"
)

func (b BorderStyle) RenderCSS(w io.Writer) error {
	_, err := w.Write([]byte(b))
	return err
}

func (b BorderStyle) valueNode()        {}
func (b BorderStyle) borderStyleValue() {}
