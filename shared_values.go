package cssgo

import "io"

// AutoType represents the CSS `auto` value.
type AutoType string

func (a AutoType) RenderCSS(w io.Writer) error {
	_, err := w.Write([]byte(a))
	return err
}

func (a AutoType) valueNode() {}

const Auto AutoType = "auto"
