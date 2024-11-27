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

func Property(name string, values ...ValueNode) PropertyNodeFunc {
	return PropertyNodeFunc(func(w io.Writer) error {
		if _, err := w.Write([]byte(name + ":")); err != nil {
			return err
		}

		for _, value := range values {
			if value == nil {
				continue
			}
			if _, err := w.Write([]byte(" ")); err != nil {
				return err
			}
			if err := value.RenderCSS(w); err != nil {
				return err
			}
		}

		if _, err := w.Write([]byte(";")); err != nil {
			return err
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

func FontSize(value SizeValue) PropertyNodeFunc {
	return Property("font-size", value)
}

func Margin(value1, value2, value3, value4 SizeValue) PropertyNodeFunc {
	return Property("margin", value1, value2, value3, value4)
}

func Padding(value1, value2, value3, value4 SizeValue) PropertyNodeFunc {
	return Property("padding", value1, value2, value3, value4)
}

func Height(value SizeValue) PropertyNodeFunc {
	return Property("height", value)
}

func Width(value SizeValue) PropertyNodeFunc {
	return Property("width", value)
}

func Border(width SizeValue, style BorderStyleValue, color ColorValue) PropertyNodeFunc {
	return Property("border", width, style, color)
}
