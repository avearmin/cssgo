package gomour

import (
	"io"
	"strings"
)

type PropertyNode interface {
	Node
	propertyNode()
}

type Property func(io.Writer) error

func (p Property) RenderCSS(w io.Writer) error {
	return p(w)
}

func (p Property) String() string {
	var b strings.Builder
	_ = p.RenderCSS(&b)
	return b.String()
}

func (p Property) propertyNode() {}

func Prop(name string, values ...ValueNode) Property {
	return Property(func(w io.Writer) error {
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

func GroupProps(props ...PropertyNode) Property {
	return Property(func(w io.Writer) error {
		for _, prop := range props {
			if err := prop.RenderCSS(w); err != nil {
				return err
			}
		}
		return nil
	})
}

func TextColor(value ColorValue) Property {
	return Prop("color", value)
}

func BackgroundColor(value ColorValue) Property {
	return Prop("background-color", value)
}

func FontSize(value SizeValue) Property {
	return Prop("font-size", value)
}

func Margin(value1, value2, value3, value4 SizeValue) Property {
	return Prop("margin", value1, value2, value3, value4)
}

func Padding(value1, value2, value3, value4 SizeValue) Property {
	return Prop("padding", value1, value2, value3, value4)
}

func Height(value SizeValue) Property {
	return Prop("height", value)
}

func Width(value SizeValue) Property {
	return Prop("width", value)
}

func Border(width SizeValue, style BorderStyleValue, color ColorValue) Property {
	return Prop("border", width, style, color)
}
