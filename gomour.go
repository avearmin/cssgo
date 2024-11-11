package gomour

import (
	"io"
	"strings"
)

type Node interface {
	RenderCSS(w io.Writer) error
}

type RuleNode interface {
	Node
	ruleNode()
}

type RuleNodeFunc func(io.Writer) error

func (r RuleNodeFunc) RenderCSS(w io.Writer) error {
	return r(w)
}

func (r RuleNodeFunc) String() string {
	var b strings.Builder
	_ = r.RenderCSS(&b)
	return b.String()
}

func (r RuleNodeFunc) ruleNode() {}

func Rule(selector SelectorNode, properties ...PropertyNode) RuleNodeFunc {
	return RuleNodeFunc(func(w io.Writer) error {
		if err := selector.RenderCSS(w); err != nil {
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

		if _, err := w.Write([]byte("}")); err != nil {
			return err
		}

		return nil
	})
}

type HTMLNodeFunc func(w io.Writer) error

func (h HTMLNodeFunc) Render(w io.Writer) error {
	return h(w)
}

func (h HTMLNodeFunc) String() string {
	var b strings.Builder
	_ = h.Render(&b)
	return b.String()
}
