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

func (p PropertyNodeFunc) RenderCSS(w io.Writer) error {
	return p(w)
}

func (p PropertyNodeFunc) String() string {
	var b strings.Builder
	_ = p.RenderCSS(&b)
	return b.String()
}

func (p PropertyNodeFunc) propertyNode() {}

func Property(name string, value ValueNode) PropertyNodeFunc {
	return PropertyNodeFunc(func(w io.Writer) error {
		if _, err := w.Write([]byte(name + ":")); err != nil {
			return err
		}

		if err := value.RenderCSS(w); err != nil {
			return err
		}

		if _, err := w.Write([]byte(";")); err != nil {
			return err
		}

		return nil
	})
}

func GroupProps(properties ...PropertyNode) PropertyNodeFunc {
	return PropertyNodeFunc(func(w io.Writer) error {
		for _, prop := range properties {
			if err := prop.RenderCSS(w); err != nil {
				return err
			}
		}
		return nil
	})
}

func TextColor(value ColorValue) PropertyNodeFunc {
	return Property("color", value)
}

func BackgroundColor(value ColorValue) PropertyNodeFunc {
	return Property("background-color", value)
}
