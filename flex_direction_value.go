package cssgo

import "io"

// FlexDirectionValue defines an interface for CSS-compatible flex direction values.
// This ensures that any type implementing this interface can be rendered as a valid CSS flex direction.
type FlexDirectionValue interface {
	ValueNode
	flexDirectionValue()
}

// FlexDirectionType represents a CSS flex direction, such as "row", "row-reverse", or "column".
// It is a concrete type that implements the FlexDirectionValue interface.
type FlexDirectionType string

func (f FlexDirectionType) RenderCSS(w io.Writer) error {
	_, err := w.Write([]byte(f))
	return err
}

func (f FlexDirectionType) valueNode()          {}
func (f FlexDirectionType) flexDirectionValue() {}

// Predefined flex direction values as per the CSS specification.
const (
	Row           FlexDirectionType = "row"
	RowReverse    FlexDirectionType = "row-reverse"
	Column        FlexDirectionType = "column"
	ColumnReverse FlexDirectionType = "column-reverse"
)
