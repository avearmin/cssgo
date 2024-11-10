package gomour

type SelectorNode interface {
	CSSNode
	selectorType()
}

type classNode struct {
	name string
}

func (c classNode) RenderCSS() string {
	return "." + c.name
}

func (c classNode) selectorType()

func Class(name string) SelectorNode {
	return classNode{name: name}
}

type idNode struct {
	name string	
}

func (i idNode) RenderCSS() string {
	return "#" + i.name
}

func (i idNode) selectorType()

func ID(name string) SelectorNode {
	return idNode{name: name}
}

type tagNode struct {
	name string
}

func (t tagNode) RenderCSS() string {
	return t.name
}

func (t tagNode) selectorType()

func Tag(name string) SelectorNode {
	return tagNode{name: name}
}

func Div() SelectorNode {
	return Tag("div")
}

func P() SelectorNode {
	return Tag("p")
}
