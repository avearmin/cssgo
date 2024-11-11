package gomour

import (
	"io"
	"strings"
)

type SelectorNode interface {
	Node
	selectorNode()
}

type selectorType string

const (
	classSelectorType selectorType = "."
	idSelectorType    selectorType = "#"
)

type SelectorNodeFunc func(io.Writer) error

func (s SelectorNodeFunc) Render(w io.Writer) error {
	return s(w)
}

func (s SelectorNodeFunc) String() string {
	var b strings.Builder
	_ = s.Render(&b)
	return b.String()
}

func (s SelectorNodeFunc) selectorNode() {}

func selector(selectorType selectorType, name string) SelectorNode {
	return SelectorNodeFunc(func(w io.Writer) error {
		_, err := w.Write([]byte(string(selectorType) + name))
		return err
	})
}

func Class(name string) SelectorNode {
	return selector(classSelectorType, name)
}

func ID(name string) SelectorNode {
	return selector(idSelectorType, name)
}

func El(tag string) SelectorNode {
	return SelectorNodeFunc(func(w io.Writer) error {
		_, err := w.Write([]byte(tag))
		return err
	})
}
