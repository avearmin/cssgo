package gomour

import "io"

type BorderStyleValue interface {
	ValueNode
	borderStyleValue()
}

type borderStyle string

const (
	Solid  borderStyle = "solid"
	Dotted borderStyle = "dotted"
	Dashed borderStyle = "dashed"
	Double borderStyle = "double"
	Groove borderStyle = "groove"
	Ridge  borderStyle = "ridge"
	Inset  borderStyle = "inset"
	Outset borderStyle = "outset"
)

func (b borderStyle) RenderCSS(w io.Writer) error {
	_, err := w.Write([]byte(b))
	return err
}

func (b borderStyle) valueNode()        {}
func (b borderStyle) borderStyleValue() {}
