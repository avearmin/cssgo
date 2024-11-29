package cssgo

import (
	"io"
	"strings"
)

type SelectorNode interface {
	Node
	selectorNode()
}

type ClassIDSelectorNode interface {
	SelectorNode
	classIDNode()
}

type SelectorNodeFunc func(io.Writer) error

func (s SelectorNodeFunc) RenderCSS(w io.Writer) error {
	return s(w)
}

func (s SelectorNodeFunc) String() string {
	var b strings.Builder
	_ = s.RenderCSS(&b)
	return b.String()
}

func (s SelectorNodeFunc) selectorNode() {}

func (s SelectorNodeFunc) Or(other SelectorNodeFunc) SelectorNodeFunc {
	return SelectorNodeFunc(func(w io.Writer) error {
		if err := s.RenderCSS(w); err != nil {
			return err
		}

		if _, err := w.Write([]byte(", ")); err != nil {
			return err
		}

		err := other.RenderCSS(w)

		return err
	})
}

func (s SelectorNodeFunc) Props(properties ...PropertyNode) RuleNodeFunc {
	return RuleNodeFunc(func(w io.Writer) error {
		if err := s.RenderCSS(w); err != nil {
			return err
		}

		if _, err := w.Write([]byte("{")); err != nil {
			return err
		}

		for _, prop := range properties {
			if err := prop.RenderCSS(w); err != nil {
				return err
			}
		}

		_, err := w.Write([]byte("}"))

		return err
	})
}

func selector(selectorLiteral, name string) SelectorNodeFunc {
	return SelectorNodeFunc(func(w io.Writer) error {
		if selectorLiteral != "" {
			if _, err := w.Write([]byte(selectorLiteral)); err != nil {
				return err
			}
		}

		_, err := w.Write([]byte(name))
		return err
	})
}

func Class(name string) SelectorNodeFunc {
	return selector(".", name)
}

func ID(name string) SelectorNodeFunc {
	return selector("#", name)
}

func El(name string) SelectorNodeFunc {
	return selector("", name)
}
