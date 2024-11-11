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

func (f ColorValueFunc) RenderCSS(w io.Writer) error {
	return f(w)
}

func (f ColorValueFunc) String() string {
	var b strings.Builder
	_ = f.RenderCSS(&b)
	return b.String()
}

func (f ColorValueFunc) valueNode()  {}
func (f ColorValueFunc) colorValue() {}

func Hex(hexdec int) ColorValueFunc {
	return ColorValueFunc(func(w io.Writer) error {
		hexStr := fmt.Sprintf("#%06x", hexdec)
		_, err := w.Write([]byte(hexStr))
		return err
	})
}

func RGB(r, g, b int) ColorValueFunc {
	return ColorValueFunc(func(w io.Writer) error {
		rgbStr := fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
		_, err := w.Write([]byte(rgbStr))
		return err
	})
}

func Color(color string) ColorValueFunc {
	return ColorValueFunc(func(w io.Writer) error {
		_, err := w.Write([]byte(color))
		return err
	})
}

func Black() ColorValueFunc {
	return Color("black")
}

func White() ColorValueFunc {
	return Color("white")
}

func Red() ColorValueFunc {
	return Color("red")
}

func Green() ColorValueFunc {
	return Color("green")
}

func Blue() ColorValueFunc {
	return Color("blue")
}

func Yellow() ColorValueFunc {
	return Color("yellow")
}

func Cyan() ColorValueFunc {
	return Color("cyan")
}

func Magenta() ColorValueFunc {
	return Color("magenta")
}

func Gray() ColorValueFunc {
	return Color("gray")
}

func Orange() ColorValueFunc {
	return Color("orange")
}

func Purple() ColorValueFunc {
	return Color("purple")
}

func Pink() ColorValueFunc {
	return Color("pink")
}

func Brown() ColorValueFunc {
	return Color("brown")
}
