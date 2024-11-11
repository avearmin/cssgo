package gomour

import (
	"io"
	"strings"
)

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

func Rule(selector SelectorNode, properties ...PropertyNode) RuleNode {
	return RuleNodeFunc(func(w io.Writer) error {
		return renderRule(w, selector, properties...)
	})
}

func renderRule(w io.Writer, selector SelectorNode, properties ...PropertyNode) error {
	sw := &statefulWriter{w: w}

	renderSelector(sw, selector)

	sw.Write([]byte("{"))

	for _, prop := range properties {
		renderProperty(sw, prop)
	}

	sw.Write([]byte("}"))

	return sw.err
}
