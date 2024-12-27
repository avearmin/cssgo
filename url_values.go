package cssgo

import (
	"fmt"
	"io"
)

// UrlyValue defines an interface for types representing CSS-compatible URL values.
type UrlValue interface {
	ValueNode
	urlValue()
}

// UrlType represents a CSS URL value (e.g., "url('images/background.png')").
type UrlType string

// RenderCSS writes the DisplayType value to the writer.
func (u UrlType) RenderCSS(w io.Writer) error {
	_, err := w.Write([]byte(u))
	return err
}

func (u UrlType) valueNode() {}
func (u UrlType) urlValue()  {}

func Url(path string) UrlType {
	return UrlType(fmt.Sprintf("url('%s')", path))
}
