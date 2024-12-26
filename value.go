package cssgo

import (
	"fmt"
	"io"
)

type ValueNode interface {
	Node
	valueNode()
}

// goType is a generic adapter for Go Types to be expressed as css values.
type goType string

func (g goType) RenderCSS(w io.Writer) error {
	_, err := w.Write([]byte(g))
	return err

}

func (g goType) valueNode() {}

// toCSS adapts Go values to be written as values for CSS properties.
func toCSS(val any) goType {
	return goType(fmt.Sprintf("%v", val))
}
