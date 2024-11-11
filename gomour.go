package gomour

import (
	"io"
	"strings"
)

type Node interface {
	Render(w io.Writer) error
}

type RuleNode interface {
	Node
	ruleNode()
}

type RuleNodeFunc func(io.Writer) error

func (r RuleNodeFunc) Render(w io.Writer) error {
	return r(w)
}

func (r RuleNodeFunc) String() string {
	var b strings.Builder
	_ = r.Render(&b)
	return b.String()
}

func (r RuleNodeFunc) ruleNode() {}

func Rule(selector SelectorNode, properties ...PropertyNode) RuleNodeFunc {
	return RuleNodeFunc(func(w io.Writer) error {
		if err := selector.Render(w); err != nil {
			return err
		}

		if _, err := w.Write([]byte("{")); err != nil {
			return err
		}

		for _, prop := range properties {
			if err := prop.Render(w); err != nil {
				return err
			}
		}

		if _, err := w.Write([]byte("}")); err != nil {
			return err
		}

		return nil
	})
}
