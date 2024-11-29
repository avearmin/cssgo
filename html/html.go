package html

import (
	"github.com/avearmin/cssgo"
	"io"
	"maragu.dev/gomponents"
	"strings"
)

type ElNodeFunc func(writer io.Writer) error

func (e ElNodeFunc) Render(w io.Writer) error {
	return e(w)
}

func (e ElNodeFunc) String() string {
	var b strings.Builder
	_ = e.Render(&b)
	return b.String()
}

func (e ElNodeFunc) Type() gomponents.NodeType {
	return gomponents.ElementType
}

func StyleEl(rules ...cssgo.RuleNode) ElNodeFunc {
	return ElNodeFunc(func(w io.Writer) error {
		if _, err := w.Write([]byte("<style>")); err != nil {
			return err
		}

		for _, rule := range rules {
			rule.RenderCSS(w)
		}

		_, err := w.Write([]byte("</style>"))

		return err
	})
}

type AttrNodeFunc func(writer io.Writer) error

func (a AttrNodeFunc) Render(w io.Writer) error {
	return a(w)
}

func (a AttrNodeFunc) String() string {
	var b strings.Builder
	_ = a.Render(&b)
	return b.String()
}

func (a AttrNodeFunc) Type() gomponents.NodeType {
	return gomponents.AttributeType
}

func Style(props ...cssgo.PropertyNode) AttrNodeFunc {
	return AttrNodeFunc(func(w io.Writer) error {
		if _, err := w.Write([]byte(" style=\"")); err != nil {
			return err
		}

		for _, prop := range props {
			prop.RenderCSS(w)
		}

		_, err := w.Write([]byte("\""))

		return err
	})
}
