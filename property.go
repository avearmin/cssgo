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

// Margin1 creates a "margin" property with one value applied to all sides.
// Example: Margin1(PX(10)) -> "margin: 10px;"
//
// Parameters:
// - value (SizeValue): The size value for all sides.
//
// Returns:
// - Property: A Property instance representing the "margin" property.
func Margin1(value SizeValue) Property {
	return Prop("margin", value)
}

// Margin2 creates a "margin" property with vertical (top/bottom) and horizontal (left/right) values.
// Example: Margin2(PX(10), PX(20)) -> "margin: 10px 20px;"
//
// Parameters:
// - vertical (SizeValue): The size value for top and bottom.
// - horizontal (SizeValue): The size value for left and right.
//
// Returns:
// - Property: A Property instance representing the "margin" property.
func Margin2(vertical, horizontal SizeValue) Property {
	return Prop("margin", vertical, horizontal)
}

// Margin3 creates a "margin" property with top, horizontal (left/right), and bottom values.
// Example: Margin3(PX(10), PX(20), PX(30)) -> "margin: 10px 20px 30px;"
//
// Parameters:
// - top (SizeValue): The size value for the top.
// - horizontal (SizeValue): The size value for left and right.
// - bottom (SizeValue): The size value for the bottom.
//
// Returns:
// - Property: A Property instance representing the "margin" property.
func Margin3(top, horizontal, bottom SizeValue) Property {
	return Prop("margin", top, horizontal, bottom)
}

// Margin4 creates a "margin" property with separate values for top, right, bottom, and left.
// Example: Margin4(PX(10), PX(20), PX(30), PX(40)) -> "margin: 10px 20px 30px 40px;"
//
// Parameters:
// - top (SizeValue): The size value for the top.
// - right (SizeValue): The size value for the right.
// - bottom (SizeValue): The size value for the bottom.
// - left (SizeValue): The size value for the left.
//
// Returns:
// - Property: A Property instance representing the "margin" property.
func Margin4(top, right, bottom, left SizeValue) Property {
	return Prop("margin", top, right, bottom, left)
}

// Padding1 creates a "padding" property with one value applied to all sides.
// Example: Padding1(PX(10)) -> "padding: 10px;"
//
// Parameters:
// - value (SizeValue): The size value for all sides.
//
// Returns:
// - Property: A Property instance representing the "padding" property.
func Padding1(value SizeValue) Property {
	return Prop("padding", value)
}

// Padding2 creates a "padding" property with vertical (top/bottom) and horizontal (left/right) values.
// Example: Padding2(PX(10), PX(20)) -> "padding: 10px 20px;"
//
// Parameters:
// - vertical (SizeValue): The size value for top and bottom.
// - horizontal (SizeValue): The size value for left and right.
//
// Returns:
// - Property: A Property instance representing the "padding" property.
func Padding2(vertical, horizontal SizeValue) Property {
	return Prop("padding", vertical, horizontal)
}

// Padding3 creates a "padding" property with top, horizontal (left/right), and bottom values.
// Example: Padding3(PX(10), PX(20), PX(30)) -> "padding: 10px 20px 30px;"
//
// Parameters:
// - top (SizeValue): The size value for the top.
// - horizontal (SizeValue): The size value for left and right.
// - bottom (SizeValue): The size value for the bottom.
//
// Returns:
// - Property: A Property instance representing the "padding" property.
func Padding3(top, horizontal, bottom SizeValue) Property {
	return Prop("padding", top, horizontal, bottom)
}

// Padding4 creates a "padding" property with separate values for top, right, bottom, and left.
// Example: Padding4(PX(10), PX(20), PX(30), PX(40)) -> "padding: 10px 20px 30px 40px;"
//
// Parameters:
// - top (SizeValue): The size value for the top.
// - right (SizeValue): The size value for the right.
// - bottom (SizeValue): The size value for the bottom.
// - left (SizeValue): The size value for the left.
//
// Returns:
// - Property: A Property instance representing the "padding" property.
func Padding4(top, right, bottom, left SizeValue) Property {
	return Prop("padding", top, right, bottom, left)
}

// PaddingTop creates a "padding-top" property.
// Example: PaddingTop(PX(10)) -> "padding-top: 10px;"
//
// Parameters:
// - value (SizeValue): The size value for the top padding.
//
// Returns:
// - Property: A Property instance representing the "padding-top" property.
func PaddingTop(value SizeValue) Property {
	return Prop("padding-top", value)
}

// PaddingRight creates a "padding-right" property.
// Example: PaddingRight(PX(20)) -> "padding-right: 20px;"
//
// Parameters:
// - value (SizeValue): The size value for the right padding.
//
// Returns:
// - Property: A Property instance representing the "padding-right" property.
func PaddingRight(value SizeValue) Property {
	return Prop("padding-right", value)
}

// PaddingBottom creates a "padding-bottom" property.
// Example: PaddingBottom(PX(30)) -> "padding-bottom: 30px;"
//
// Parameters:
// - value (SizeValue): The size value for the bottom padding.
//
// Returns:
// - Property: A Property instance representing the "padding-bottom" property.
func PaddingBottom(value SizeValue) Property {
	return Prop("padding-bottom", value)
}

// PaddingLeft creates a "padding-left" property.
// Example: PaddingLeft(PX(40)) -> "padding-left: 40px;"
//
// Parameters:
// - value (SizeValue): The size value for the left padding.
//
// Returns:
// - Property: A Property instance representing the "padding-left" property.
func PaddingLeft(value SizeValue) Property {
	return Prop("padding-left", value)
}

// Border1 creates a "border" property with just the border size applied to all sides.
// Example: Border1(PX(1)) -> "border: 1px;"
//
// Parameters:
// - width (SizeValue): The width of the border for all sides.
//
// Returns:
// - Property: A Property instance representing the "border" property.
func Border1(width SizeValue) Property {
	return Prop("border", width)
}

// Border2 creates a "border" property with size and style applied to all sides.
// Example: Border2(PX(1), Solid) -> "border: 1px solid;"
//
// Parameters:
// - width (SizeValue): The width of the border for all sides.
// - style (BorderStyleValue): The style of the border (e.g., Solid, Dashed).
//
// Returns:
// - Property: A Property instance representing the "border" property.
func Border2(width SizeValue, style BorderStyleValue) Property {
	return Prop("border", width, style)
}

// Border3 creates a "border" property with size, style, and color applied to all sides.
// Example: Border3(PX(1), Solid, Black) -> "border: 1px solid black;"
//
// Parameters:
// - width (SizeValue): The width of the border for all sides.
// - style (BorderStyleValue): The style of the border (e.g., Solid, Dashed).
// - color (ColorValue): The color of the border.
//
// Returns:
// - Property: A Property instance representing the "border" property.
func Border3(width SizeValue, style BorderStyleValue, color ColorValue) Property {
	return Prop("border", width, style, color)
}

// BorderStyle1 creates a "border-style" property with one value applied to all sides.
// Example: BorderStyle1(Solid) -> "border-style: solid;"
//
// Parameters:
// - value (BorderStyleValue): The border style for all sides.
//
// Returns:
// - Property: A Property instance representing the "border-style" property.
func BorderStyle1(value BorderStyleValue) Property {
	return Prop("border-style", value)
}

// BorderStyle2 creates a "border-style" property with vertical and horizontal styles.
// Example: BorderStyle2(Solid, Dashed) -> "border-style: solid dashed;"
//
// Parameters:
// - vertical (BorderStyleValue): The border style for top and bottom.
// - horizontal (BorderStyleValue): The border style for left and right.
//
// Returns:
// - Property: A Property instance representing the "border-style" property.
func BorderStyle2(vertical, horizontal BorderStyleValue) Property {
	return Prop("border-style", vertical, horizontal)
}

// BorderStyle3 creates a "border-style" property with top, horizontal, and bottom styles.
// Example: BorderStyle3(Solid, Dashed, Dotted) -> "border-style: solid dashed dotted;"
//
// Parameters:
// - top (BorderStyleValue): The border style for the top.
// - horizontal (BorderStyleValue): The border style for left and right.
// - bottom (BorderStyleValue): The border style for the bottom.
//
// Returns:
// - Property: A Property instance representing the "border-style" property.
func BorderStyle3(top, horizontal, bottom BorderStyleValue) Property {
	return Prop("border-style", top, horizontal, bottom)
}

// BorderStyle4 creates a "border-style" property with separate styles for top, right, bottom, and left.
// Example: BorderStyle4(Solid, Dashed, Dotted, Double) -> "border-style: solid dashed dotted double;"
//
// Parameters:
// - top (BorderStyleValue): The border style for the top.
// - right (BorderStyleValue): The border style for the right.
// - bottom (BorderStyleValue): The border style for the bottom.
// - left (BorderStyleValue): The border style for the left.
//
// Returns:
// - Property: A Property instance representing the "border-style" property.
func BorderStyle4(top, right, bottom, left BorderStyleValue) Property {
	return Prop("border-style", top, right, bottom, left)
}
