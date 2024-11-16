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

func RGBA(r, g, b int, a float64) ColorValueFunc {
	return ColorValueFunc(func(w io.Writer) error {
		rgbaStr := fmt.Sprintf("rgba(%d, %d, %d, %f)", r, g, b, a)
		_, err := w.Write([]byte(rgbaStr))
		return err
	})
}

func Color(color string) ColorValueFunc {
	return ColorValueFunc(func(w io.Writer) error {
		_, err := w.Write([]byte(color))
		return err
	})
}

func HSL(hue uint, saturation, lightness float64) ColorValueFunc {
	if hue > 360 {
		hue = 360
	}
	if saturation < 0.0 {
		saturation = 0.0
	}
	if saturation > 1.0 {
		saturation = 1.0
	}
	if lightness < 0.0 {
		lightness = 0.0
	}
	if lightness > 1.0 {
		lightness = 1.0
	}

	return ColorValueFunc(func(w io.Writer) error {
		hslStr := fmt.Sprintf("hsl(%d, %.0f%%, %.0f%%)", hue, saturation*100, lightness*100)
		_, err := w.Write([]byte(hslStr))
		return err
	})
}
