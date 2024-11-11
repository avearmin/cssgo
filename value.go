package gomour

import (
	"fmt"
	"io"
	"strings"
)

type ValueNode interface {
	Node
	valueNode()
}

type ColorValue interface {
	ValueNode
	colorValue()
}

type ColorValueFunc func(io.Writer) error

func (f ColorValueFunc) Render(w io.Writer) error {
	return f(w)
}

func (f ColorValueFunc) String() string {
	var b strings.Builder
	_ = f.Render(&b)
	return b.String()
}

func (f ColorValueFunc) valueNode()  {}
func (f ColorValueFunc) colorValue() {}

func Hex(hexdec int) ColorValue {
	return ColorValueFunc(func(w io.Writer) error {
		hexStr := fmt.Sprintf("#%06x", hexdec)
		_, err := w.Write([]byte(hexStr))
		return err
	})
}

func RGB(r, g, b int) ColorValue {
	return ColorValueFunc(func(w io.Writer) error {
		rgbStr := fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
		_, err := w.Write([]byte(rgbStr))
		return err
	})
}

func Color(color string) ColorValue {
	return ColorValueFunc(func(w io.Writer) error {
		_, err := w.Write([]byte(color))
		return err
	})
}

func Black() ColorValue {
	return Color("black")
}

func White() ColorValue {
	return Color("white")
}

func Red() ColorValue {
	return Color("red")
}

func Green() ColorValue {
	return Color("green")
}

func Blue() ColorValue {
	return Color("blue")
}

func Yellow() ColorValue {
	return Color("yellow")
}

func Cyan() ColorValue {
	return Color("cyan")
}

func Magenta() ColorValue {
	return Color("magenta")
}

func Gray() ColorValue {
	return Color("gray")
}

func Orange() ColorValue {
	return Color("orange")
}

func Purple() ColorValue {
	return Color("purple")
}

func Pink() ColorValue {
	return Color("pink")
}

func Brown() ColorValue {
	return Color("brown")
}
