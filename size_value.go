package gomour

import (
	"io"
	"strconv"
)

type SizeValue interface {
	ValueNode
	sizeValue()
}

type Size string

func (s Size) RenderCSS(w io.Writer) error {
	_, err := w.Write([]byte(s))
	return err
}

func (s Size) valueNode() {}
func (s Size) sizeValue() {}

func size(value float64, unit string) Size {
	floatStr := strconv.FormatFloat(value, 'f', -1, 64)
	return Size(floatStr + unit)
}

func CM(value float64) Size {
	return size(value, "cm")
}

func MM(value float64) Size {
	return size(value, "mm")
}

func IN(value float64) Size {
	return size(value, "in")
}

func PX(value float64) Size {
	return size(value, "px")
}

func PT(value float64) Size {
	return size(value, "pt")
}

func PC(value float64) Size {
	return size(value, "pc")
}

func EM(value float64) Size {
	return size(value, "em")
}

func REM(value float64) Size {
	return size(value, "rem")
}

func VW(value float64) Size {
	return size(value, "vw")
}

func VH(value float64) Size {
	return size(value, "vh")
}

func PCT(value float64) Size {
	return size(value, "%")
}

func VMIN(value float64) Size {
	return size(value, "vmin")
}

func VMAX(value float64) Size {
	return size(value, "vmax")
}
