package gomour

import "io"

type BorderStyleValue interface {
	ValueNode
	borderStyleValue()
}

type BorderStyle string

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
