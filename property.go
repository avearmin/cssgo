package cssgo

import (
	"io"
	"strings"
)

// PropertyNode represents a CSS property node.
// Any type implementing this interface can be rendered as a valid CSS property.
type PropertyNode interface {
	Node
	propertyNode()
}

// Property is a concrete type that represents a CSS property as a function.
// It implements the PropertyNode interface, allowing it to be rendered as CSS.
type Property func(io.Writer) error

// RenderCSS writes the CSS representation of the property to the provided io.Writer.
// Example: For a property like "color: red;", this method writes "color: red;".
//
// Parameters:
// - w (io.Writer): The destination where the CSS representation is written.
//
// Returns:
// - error: If an error occurs during the write operation, it is returned; otherwise, nil.
func (p Property) RenderCSS(w io.Writer) error {
	return p(w)
}

func (p Property) String() string {
	var b strings.Builder
	_ = p.RenderCSS(&b)
	return b.String()
}

func (p Property) propertyNode() {}

// Prop creates a CSS property with a name and one or more values.
// Example: Prop("color", Red) -> "color: red;"
//
// Parameters:
// - name (string): The name of the CSS property (e.g., "color").
// - values (...ValueNode): One or more values for the property (e.g., "red", "10px").
//
// Returns:
// - Property: A Property instance representing the CSS property.
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

// GroupProps groups multiple properties into a single property.
// This is useful for rendering multiple properties together.
//
// Parameters:
// - props (...PropertyNode): One or more PropertyNode instances.
//
// Returns:
// - Property: A Property instance that renders all the provided properties.
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

// TextColor creates a "color" property for text color.
// Example: TextColor(Red) -> "color: red;"
//
// Parameters:
// - value (ColorValue): The color value for the text.
//
// Returns:
// - Property: A Property instance representing the "color" property.
func TextColor(value ColorValue) Property {
	return Prop("color", value)
}

// BackgroundColor creates a "background-color" property.
// Example: BackgroundColor(Blue) -> "background-color: blue;"
//
// Parameters:
// - value (ColorValue): The color value for the background.
//
// Returns:
// - Property: A Property instance representing the "background-color" property.
func BackgroundColor(value ColorValue) Property {
	return Prop("background-color", value)
}

// FontSize creates a "font-size" property.
// Example: FontSize(PX(16)) -> "font-size: 16px;"
//
// Parameters:
// - value (SizeValue): The size value for the font.
//
// Returns:
// - Property: A Property instance representing the "font-size" property.
func FontSize(value SizeValue) Property {
	return Prop("font-size", value)
}

// Margin creates a "margin" property with up to four size values.
// Example: Margin(PX(10), PX(20), nil, nil) -> "margin: 10px 20px;"
//
// Parameters:
// - value1, value2, value3, value4 (SizeValue): The size values for top, right, bottom, and left.
//
// Returns:
// - Property: A Property instance representing the "margin" property.
func Margin(value1, value2, value3, value4 SizeValue) Property {
	return Prop("margin", value1, value2, value3, value4)
}

// Padding creates a "padding" property with up to four size values.
// Example: Padding(PX(10), PX(20), PX(30), PX(40)) -> "padding: 10px 20px 30px 40px;"
//
// Parameters:
// - value1, value2, value3, value4 (SizeValue): The size values for top, right, bottom, and left.
//
// Returns:
// - Property: A Property instance representing the "padding" property.
func Padding(value1, value2, value3, value4 SizeValue) Property {
	return Prop("padding", value1, value2, value3, value4)
}

// Height creates a "height" property.
// Example: Height(PX(100)) -> "height: 100px;"
//
// Parameters:
// - value (SizeValue): The size value for the height.
//
// Returns:
// - Property: A Property instance representing the "height" property.
func Height(value SizeValue) Property {
	return Prop("height", value)
}

// Width creates a "width" property.
// Example: Width(PX(200)) -> "width: 200px;"
//
// Parameters:
// - value (SizeValue): The size value for the width.
//
// Returns:
// - Property: A Property instance representing the "width" property.
func Width(value SizeValue) Property {
	return Prop("width", value)
}

// Border creates a "border" property with width, style, and color.
// Example: Border(PX(1), Solid, Black) -> "border: 1px solid black;"
//
// Parameters:
// - width (SizeValue): The width of the border.
// - style (BorderStyleValue): The style of the border (e.g., "solid", "dotted").
// - color (ColorValue): The color of the border.
//
// Returns:
// - Property: A Property instance representing the "border" property.
func Border(width SizeValue, style BorderStyleValue, color ColorValue) Property {
	return Prop("border", width, style, color)
}
