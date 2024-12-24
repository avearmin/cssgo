package cssgo

import "io"

type FlexDirectionValue interface {
	ValueNode
	flexDirectionValue()
}

type FlexDirectionType string

func (f FlexDirectionType) RenderCSS(w io.Writer) error {
	_, err := w.Write([]byte(f))
	return err
}

func (f FlexDirectionType) valueNode()        {}
func (f FlexDirectionType) borderStyleValue() {}

const (
	Row           FlexDirectionType = "row"
	RowReverse    FlexDirectionType = "row-reverse"
	Column        FlexDirectionType = "column"
	ColumnReverse FlexDirectionType = "column-reverse"
)
