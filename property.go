package gomour

import (
	"io"
	"strings"
)

type PropertyNode interface {
	Node
	propertyNode()
}

type PropertyNodeFunc func(io.Writer) error

func (p PropertyNodeFunc) Render(w io.Writer) error {
	return p(w)
}

func (p PropertyNodeFunc) String() string {
	var b strings.Builder
	_ = p.Render(&b)
	return b.String()
}

func (p PropertyNodeFunc) propertyNode() {}

func Property(name string, value ValueNode) PropertyNode {
	return PropertyNodeFunc(func(w io.Writer) error {
		if _, err := w.Write([]byte(name + ":")); err != nil {
			return err
		}

		if err := value.Render(w); err != nil {
			return err
		}

		if _, err := w.Write([]byte(";")); err != nil {
			return err
		}

		return nil
	})
}

func TextColor(value ColorValue) PropertyNode {
	return Property("color", value)
}

func BackgroundColor(value ColorValue) PropertyNode {
	return Property("background-color", value)
}
