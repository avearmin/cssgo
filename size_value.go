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

func size(n int, unit string) SizeValueFunc {
	return SizeValueFunc(func(w io.Writer) error {

		if _, err := w.Write([]byte(strconv.Itoa(n))); err != nil {
			return err
		}
		_, err := w.Write([]byte(unit))
		return err
	})
}

func CM(n int) SizeValueFunc {
	return size(n, "cm")
}

func MM(n int) SizeValueFunc {
	return size(n, "mm")
}

func IN(n int) SizeValueFunc {
	return size(n, "in")
}

func PX(n int) SizeValueFunc {
	return size(n, "px")
}

func PT(n int) SizeValueFunc {
	return size(n, "pt")
}

func PC(n int) SizeValueFunc {
	return size(n, "pc")
}

func EM(n int) SizeValueFunc {
	return size(n, "em")
}

func REM(n int) SizeValueFunc {
	return size(n, "rem")
}

func VW(n int) SizeValueFunc {
	return size(n, "vw")
}

func VH(n int) SizeValueFunc {
	return size(n, "vh")
}

func Per(n int) SizeValueFunc {
	return size(n, "%")
}

func VMIN(n int) SizeValueFunc {
	return size(n, "vmin")
}

func VMAX(n int) SizeValueFunc {
	return size(n, "vmax")
}
