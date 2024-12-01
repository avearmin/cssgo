package cssgo

import "io"

// BorderStyleValue defines an interface for CSS-compatible border style values.
// This ensures that any type implementing this interface can be rendered as a valid CSS border style.
type BorderStyleValue interface {
	ValueNode
	borderStyleValue()
}

// BorderDisplay represents a CSS border style, such as "solid", "dotted", or "dashed".
// It is a concrete type that implements the BorderStyleValue interface.
type BorderDisplay string

// Predefined border style values as per the CSS specification.
const (
	Solid  BorderDisplay = "solid"
	Dotted BorderDisplay = "dotted"
	Dashed BorderDisplay = "dashed"
	Double BorderDisplay = "double"
	Groove BorderDisplay = "groove"
	Ridge  BorderDisplay = "ridge"
	Inset  BorderDisplay = "inset"
	Outset BorderDisplay = "outset"
)

func (b BorderDisplay) RenderCSS(w io.Writer) error {
	_, err := w.Write([]byte(b))
	return err
}

func (b BorderDisplay) valueNode()        {}
func (b BorderDisplay) borderStyleValue() {}

// BorderSizeValue defines an interface for CSS-compatible border width values.
// This ensures that any type implementing this interface can be rendered as a valid CSS border style.
type BorderWidthValue interface {
	ValueNode
	borderWidthValue()
}

// BorderSize represents a CSS border width size, such as "thin", "medium", "thick", or "10px".
// It is a concrete type that implements the BorderWidthValue interface.
type BorderSize string

// Predefined border width values as per the CSS specification.
const (
	Thin   BorderSize = "thin"
	Medium BorderSize = "medium"
	Thick  BorderSize = "thick"
)

func (b BorderSize) RenderCSS(w io.Writer) error {
	_, err := w.Write([]byte(b))
	return err
}

func (b BorderSize) valueNode()        {}
func (b BorderSize) borderWidthValue() {}
