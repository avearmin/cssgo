package cssgo

import "io"

// DisplayValue defines an interface for types representing CSS-compatible DisplayType values.
type DisplayValue interface {
	ValueNode
	displayValue()
}

// DisplayType represents a CSS display value (e.g., "block", "inline", "flex").
type DisplayType string

// RenderCSS writes the DisplayType value to the writer.
func (d DisplayType) RenderCSS(w io.Writer) error {
	_, err := w.Write([]byte(d))
	return err
}

func (d DisplayType) valueNode()    {}
func (d DisplayType) displayValue() {}

// DisplayType value constants.
const (
	Block            DisplayType = "block"
	Inline           DisplayType = "inline"
	InlineBlock      DisplayType = "inline-block"
	Flex             DisplayType = "flex"
	Grid             DisplayType = "grid"
	InlineFlex       DisplayType = "inline-flex"
	InlineGrid       DisplayType = "inline-grid"
	Table            DisplayType = "table"
	TableRow         DisplayType = "table-row"
	TableCell        DisplayType = "table-cell"
	ListItem         DisplayType = "list-item"
	RunIn            DisplayType = "run-in"
	FlowRoot         DisplayType = "flow-root"
	Contents         DisplayType = "contents"
	TableColumn      DisplayType = "table-column"
	TableColumnGroup DisplayType = "table-column-group"
	TableHeaderGroup DisplayType = "table-header-group"
	TableFooterGroup DisplayType = "table-footer-group"
	TableRowGroup    DisplayType = "table-row-group"
)
