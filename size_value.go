package gomour

import (
	"io"
	"strconv"
	"strings"
)

type SizeValue interface {
	ValueNode
	sizeValue()
}

type SizeValueFunc func(w io.Writer) error

func (f SizeValueFunc) RenderCSS(w io.Writer) error {
	return f(w)
}

func (f SizeValueFunc) String() string {
	var b strings.Builder
	_ = f.RenderCSS(&b)
	return b.String()
}

func (f SizeValueFunc) valueNode() {}
func (f SizeValueFunc) sizeValue() {}

func size(value float64, unit string) SizeValueFunc {
	return SizeValueFunc(func(w io.Writer) error {
		floatStr := strconv.FormatFloat(value, 'f', -1, 64)
		if _, err := w.Write([]byte(floatStr)); err != nil {
			return err
		}
		_, err := w.Write([]byte(unit))
		return err
	})
}

func CM(value float64) SizeValueFunc {
	return size(value, "cm")
}

func MM(value float64) SizeValueFunc {
	return size(value, "mm")
}

func IN(value float64) SizeValueFunc {
	return size(value, "in")
}

func PX(value float64) SizeValueFunc {
	return size(value, "px")
}

func PT(value float64) SizeValueFunc {
	return size(value, "pt")
}

func PC(value float64) SizeValueFunc {
	return size(value, "pc")
}

func EM(value float64) SizeValueFunc {
	return size(value, "em")
}

func REM(value float64) SizeValueFunc {
	return size(value, "rem")
}

func VW(value float64) SizeValueFunc {
	return size(value, "vw")
}

func VH(value float64) SizeValueFunc {
	return size(value, "vh")
}

func PCT(value float64) SizeValueFunc {
	return size(value, "%")
}

func VMIN(value float64) SizeValueFunc {
	return size(value, "vmin")
}

func VMAX(value float64) SizeValueFunc {
	return size(value, "vmax")
}
