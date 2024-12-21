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
func TextColor(value ColorValue) Property {
	return Prop("color", value)
}

// BackgroundColor creates a "background-color" property.
// Example: BackgroundColor(Blue) -> "background-color: blue;"
func BackgroundColor(value ColorValue) Property {
	return Prop("background-color", value)
}

// FontSize creates a "font-size" property.
// Example: FontSize(PX(16)) -> "font-size: 16px;"
func FontSize(value SizeValue) Property {
	return Prop("font-size", value)
}

// Height creates a "height" property.
// Example: Height(PX(100)) -> "height: 100px;"
func Height(value SizeValue) Property {
	return Prop("height", value)
}

// Width creates a "width" property.
// Example: Width(PX(200)) -> "width: 200px;"
func Width(value SizeValue) Property {
	return Prop("width", value)
}

// Margin1 creates a "margin" property with one value applied to all sides.
// Example: Margin1(PX(10)) -> "margin: 10px;"
func Margin1(value SizeValue) Property {
	return Prop("margin", value)
}

// Margin2 creates a "margin" property with vertical (top/bottom) and horizontal (left/right) values.
// Example: Margin2(PX(10), PX(20)) -> "margin: 10px 20px;"
func Margin2(vertical, horizontal SizeValue) Property {
	return Prop("margin", vertical, horizontal)
}

// Margin3 creates a "margin" property with top, horizontal (left/right), and bottom values.
// Example: Margin3(PX(10), PX(20), PX(30)) -> "margin: 10px 20px 30px;"
func Margin3(top, horizontal, bottom SizeValue) Property {
	return Prop("margin", top, horizontal, bottom)
}

// Margin4 creates a "margin" property with separate values for top, right, bottom, and left.
// Example: Margin4(PX(10), PX(20), PX(30), PX(40)) -> "margin: 10px 20px 30px 40px;"
func Margin4(top, right, bottom, left SizeValue) Property {
	return Prop("margin", top, right, bottom, left)
}

// Padding1 creates a "padding" property with one value applied to all sides.
// Example: Padding1(PX(10)) -> "padding: 10px;"
func Padding1(value SizeValue) Property {
	return Prop("padding", value)
}

// Padding2 creates a "padding" property with vertical (top/bottom) and horizontal (left/right) values.
// Example: Padding2(PX(10), PX(20)) -> "padding: 10px 20px;"
func Padding2(vertical, horizontal SizeValue) Property {
	return Prop("padding", vertical, horizontal)
}

// Padding3 creates a "padding" property with top, horizontal (left/right), and bottom values.
// Example: Padding3(PX(10), PX(20), PX(30)) -> "padding: 10px 20px 30px;"
func Padding3(top, horizontal, bottom SizeValue) Property {
	return Prop("padding", top, horizontal, bottom)
}

// Padding4 creates a "padding" property with separate values for top, right, bottom, and left.
// Example: Padding4(PX(10), PX(20), PX(30), PX(40)) -> "padding: 10px 20px 30px 40px;"
func Padding4(top, right, bottom, left SizeValue) Property {
	return Prop("padding", top, right, bottom, left)
}

// PaddingTop creates a "padding-top" property.
// Example: PaddingTop(PX(10)) -> "padding-top: 10px;"
func PaddingTop(value SizeValue) Property {
	return Prop("padding-top", value)
}

// PaddingRight creates a "padding-right" property.
// Example: PaddingRight(PX(20)) -> "padding-right: 20px;"
func PaddingRight(value SizeValue) Property {
	return Prop("padding-right", value)
}

// PaddingBottom creates a "padding-bottom" property.
// Example: PaddingBottom(PX(30)) -> "padding-bottom: 30px;"
func PaddingBottom(value SizeValue) Property {
	return Prop("padding-bottom", value)
}

// PaddingLeft creates a "padding-left" property.
// Example: PaddingLeft(PX(40)) -> "padding-left: 40px;"
func PaddingLeft(value SizeValue) Property {
	return Prop("padding-left", value)
}

// Border1 creates a "border" property with just the border size applied to all sides.
// Example: Border1(PX(1)) -> "border: 1px;"
func Border1(width BorderWidthValue) Property {
	return Prop("border", width)
}

// Border2 creates a "border" property with size and style applied to all sides.
// Example: Border2(PX(1), Solid) -> "border: 1px solid;"
func Border2(width BorderWidthValue, style BorderStyleValue) Property {
	return Prop("border", width, style)
}

// Border3 creates a "border" property with size, style, and color applied to all sides.
// Example: Border3(PX(1), Solid, Black) -> "border: 1px solid black;"
func Border3(width BorderWidthValue, style BorderStyleValue, color ColorValue) Property {
	return Prop("border", width, style, color)
}

// BorderStyle1 creates a "border-style" property with one value applied to all sides.
// Example: BorderStyle1(Solid) -> "border-style: solid;"
func BorderStyle1(value BorderStyleValue) Property {
	return Prop("border-style", value)
}

// BorderStyle2 creates a "border-style" property with vertical and horizontal styles.
// Example: BorderStyle2(Solid, Dashed) -> "border-style: solid dashed;"
func BorderStyle2(vertical, horizontal BorderStyleValue) Property {
	return Prop("border-style", vertical, horizontal)
}

// BorderStyle3 creates a "border-style" property with top, horizontal, and bottom styles.
// Example: BorderStyle3(Solid, Dashed, Dotted) -> "border-style: solid dashed dotted;"
func BorderStyle3(top, horizontal, bottom BorderStyleValue) Property {
	return Prop("border-style", top, horizontal, bottom)
}

// BorderStyle4 creates a "border-style" property with separate styles for top, right, bottom, and left.
// Example: BorderStyle4(Solid, Dashed, Dotted, Double) -> "border-style: solid dashed dotted double;"
func BorderStyle4(top, right, bottom, left BorderStyleValue) Property {
	return Prop("border-style", top, right, bottom, left)
}

// BorderWidth1 creates a "border-width" property with a single value applied to all sides.
// Example: BorderWidth1(PX(10)) -> "border-width: 10px;"
func BorderWidth1(value BorderWidthValue) Property {
	return Prop("border-width", value)
}

// BorderWidth2 creates a "border-width" property with vertical (top/bottom) and horizontal (left/right) values.
// Example: BorderWidth2(PX(10), PX(20)) -> "border-width: 10px 20px;"
func BorderWidth2(vertical, horizontal BorderWidthValue) Property {
	return Prop("border-width", vertical, horizontal)
}

// BorderWidth3 creates a "border-width" property with top, horizontal (left/right), and bottom values.
// Example: BorderWidth3(PX(10), PX(20), PX(30)) -> "border-width: 10px 20px 30px;"
func BorderWidth3(top, horizontal, bottom BorderWidthValue) Property {
	return Prop("border-width", top, horizontal, bottom)
}

// BorderWidth4 creates a "border-width" property with separate values for top, right, bottom, and left.
// Example: BorderWidth4(PX(10), PX(20), PX(30), PX(40)) -> "border-width: 10px 20px 30px 40px;"
func BorderWidth4(top, right, bottom, left BorderWidthValue) Property {
	return Prop("border-width", top, right, bottom, left)
}

// BorderTopWidth creates a "border-top-width" property.
// Example: BorderTopWidth(Thick) -> "border-top-width: thick;"
func BorderTopWidth(top BorderWidthValue) Property {
	return Prop("border-top-width", top)
}

// BorderRightWidth creates a "border-right-width" property.
// Example: BorderRightWidth(Medium) -> "border-right-width: medium;"
func BorderRightWidth(right BorderWidthValue) Property {
	return Prop("border-right-width", right)
}

// BorderBottomWidth creates a "border-bottom-width" property.
// Example: BorderBottomWidth(PX(10)) -> "border-bottom-width: 10px;"
func BorderBottomWidth(bottom BorderWidthValue) Property {
	return Prop("border-bottom-width", bottom)
}

// BorderLeftWidth creates a "border-left-width" property.
// Example: BorderLeftWidth(Thin) -> "border-left-width: thin;"
func BorderLeftWidth(left BorderWidthValue) Property {
	return Prop("border-left-width", left)
}

// BorderColor1 creates a "border-color" property with one value applied to all sides.
// Example: BorderColor1(Red) -> "border-color: red;"
func BorderColor1(value ColorValue) Property {
	return Prop("border-color", value)
}

// BorderColor2 creates a "border-color" property with vertical (top/bottom) and horizontal (left/right) values.
// Example: BorderColor2(Red, Blue) -> "border-color: red blue;"
func BorderColor2(vertical, horizontal ColorValue) Property {
	return Prop("border-color", vertical, horizontal)
}

// BorderColor3 creates a "border-color" property with top, horizontal (left/right), and bottom values.
// Example: BorderColor3(Red, Blue, Green) -> "border-color: red blue green;"
func BorderColor3(top, horizontal, bottom ColorValue) Property {
	return Prop("border-color", top, horizontal, bottom)
}

// BorderColor4 creates a "border-color" property with separate values for top, right, bottom, and left.
// Example: BorderColor4(Red, Blue, Green, Purple) -> "border-color: red blue green purple;"
func BorderColor4(top, right, bottom, left ColorValue) Property {
	return Prop("border-color", top, right, bottom, left)
}

// BorderRadius sets the border-radius on all corners.
// Example: BorderRadius(PX(5)) -> "border-radius: 5px;"
func BorderRadius(value SizeValue) Property {
	return Prop("border-radius", value)
}

// BorderRadius2 sets the top-left and top-right radii, which also apply
// to bottom-right and bottom-left respectively if not otherwise specified.
// Example: BorderRadius2(PX(5), PX(10)) -> "border-radius: 5px 10px;"
func BorderRadius2(topLeft, topRight SizeValue) Property {
	return Prop("border-radius", topLeft, topRight)
}

// BorderRadius3 sets top-left, top-right, and bottom-right radii,
// with bottom-left defaulting if omitted.
// Example: BorderRadius3(PX(5), PX(10), PX(15)) -> "border-radius: 5px 10px 15px;"
func BorderRadius3(topLeft, topRight, bottomRight SizeValue) Property {
	return Prop("border-radius", topLeft, topRight, bottomRight)
}

// BorderRadius4 sets all four corners (top-left, top-right, bottom-right, bottom-left).
// Example: BorderRadius4(PX(5), PX(10), PX(15), PX(20)) -> "border-radius: 5px 10px 15px 20px;"
func BorderRadius4(topLeft, topRight, bottomRight, bottomLeft SizeValue) Property {
	return Prop("border-radius", topLeft, topRight, bottomRight, bottomLeft)
}

// BorderTopLeftRadius sets only the top-left corner radius.
// Example: BorderTopLeftRadius(PX(5)) -> "border-top-left-radius: 5px;"
func BorderTopLeftRadius(value SizeValue) Property {
	return Prop("border-top-left-radius", value)
}

// BorderTopRightRadius sets only the top-right corner radius.
// Example: BorderTopRightRadius(PX(5)) -> "border-top-right-radius: 5px;"
func BorderTopRightRadius(value SizeValue) Property {
	return Prop("border-top-right-radius", value)
}

// BorderBottomRightRadius sets only the bottom-right corner radius.
// Example: BorderBottomRightRadius(PX(5)) -> "border-bottom-right-radius: 5px;"
func BorderBottomRightRadius(value SizeValue) Property {
	return Prop("border-bottom-right-radius", value)
}

// BorderBottomLeftRadius sets only the bottom-left corner radius.
// Example: BorderBottomLeftRadius(PX(5)) -> "border-bottom-left-radius: 5px;"
func BorderBottomLeftRadius(value SizeValue) Property {
	return Prop("border-bottom-left-radius", value)
}

// LineHeight sets the line-height property using a SizeValue.
// Example: LineHeight(EM(1.5)) -> "line-height: 1.5em;"
func LineHeight(value SizeValue) Property {
	return Prop("line-height", value)
}
