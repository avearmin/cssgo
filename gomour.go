package cssgo

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
