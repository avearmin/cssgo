package gomour

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

type selectorType string

const (
	classSelectorType selectorType = "."
	idSelectorType    selectorType = "#"
)

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

func selector(selectorType selectorType, name string) SelectorNodeFunc {
	return SelectorNodeFunc(func(w io.Writer) error {
		_, err := w.Write([]byte(string(selectorType) + name))
		return err
	})
}

func Class(name string) SelectorNodeFunc {
	return selector(classSelectorType, name)
}

func ID(name string) SelectorNodeFunc {
	return selector(idSelectorType, name)
}

func El(tag string) SelectorNodeFunc {
	return SelectorNodeFunc(func(w io.Writer) error {
		_, err := w.Write([]byte(tag))
		return err
	})
}
