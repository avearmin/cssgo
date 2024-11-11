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

func A() SelectorNode {
	return El("a")
}

func Abbr() SelectorNode {
	return El("abbr")
}

func Address() SelectorNode {
	return El("address")
}

func Area() SelectorNode {
	return El("area")
}

func Article() SelectorNode {
	return El("article")
}

func Aside() SelectorNode {
	return El("aside")
}

func Audio() SelectorNode {
	return El("audio")
}

func B() SelectorNode {
	return El("b")
}

func Base() SelectorNode {
	return El("base")
}

func Bdi() SelectorNode {
	return El("bdi")
}

func Bdo() SelectorNode {
	return El("bdo")
}

func Blockquote() SelectorNode {
	return El("blockquote")
}

func Body() SelectorNode {
	return El("body")
}

func Br() SelectorNode {
	return El("br")
}

func Button() SelectorNode {
	return El("button")
}

func Canvas() SelectorNode {
	return El("canvas")
}

func Caption() SelectorNode {
	return El("caption")
}

func Cite() SelectorNode {
	return El("cite")
}

func Code() SelectorNode {
	return El("code")
}

func Col() SelectorNode {
	return El("col")
}

func Colgroup() SelectorNode {
	return El("colgroup")
}

func Data() SelectorNode {
	return El("data")
}

func Datalist() SelectorNode {
	return El("datalist")
}

func Dd() SelectorNode {
	return El("dd")
}

func Del() SelectorNode {
	return El("del")
}

func Details() SelectorNode {
	return El("details")
}

func Dfn() SelectorNode {
	return El("dfn")
}

func Dialog() SelectorNode {
	return El("dialog")
}

func Div() SelectorNode {
	return El("div")
}

func Dl() SelectorNode {
	return El("dl")
}

func Dt() SelectorNode {
	return El("dt")
}

func Em() SelectorNode {
	return El("em")
}

func Embed() SelectorNode {
	return El("embed")
}

func Fieldset() SelectorNode {
	return El("fieldset")
}

func Figcaption() SelectorNode {
	return El("figcaption")
}

func Figure() SelectorNode {
	return El("figure")
}

func Footer() SelectorNode {
	return El("footer")
}

func Form() SelectorNode {
	return El("form")
}

func H1() SelectorNode {
	return El("h1")
}

func H2() SelectorNode {
	return El("h2")
}

func H3() SelectorNode {
	return El("h3")
}

func H4() SelectorNode {
	return El("h4")
}

func H5() SelectorNode {
	return El("h5")
}

func H6() SelectorNode {
	return El("h6")
}

func Head() SelectorNode {
	return El("head")
}

func Header() SelectorNode {
	return El("header")
}

func Hr() SelectorNode {
	return El("hr")
}

func Html() SelectorNode {
	return El("html")
}

func I() SelectorNode {
	return El("i")
}

func Iframe() SelectorNode {
	return El("iframe")
}

func Img() SelectorNode {
	return El("img")
}

func Input() SelectorNode {
	return El("input")
}

func Ins() SelectorNode {
	return El("ins")
}

func Kbd() SelectorNode {
	return El("kbd")
}

func Label() SelectorNode {
	return El("label")
}

func Legend() SelectorNode {
	return El("legend")
}

func Li() SelectorNode {
	return El("li")
}

func Link() SelectorNode {
	return El("link")
}

func Main() SelectorNode {
	return El("main")
}

func Map() SelectorNode {
	return El("map")
}

func Mark() SelectorNode {
	return El("mark")
}

func Meta() SelectorNode {
	return El("meta")
}

func Meter() SelectorNode {
	return El("meter")
}

func Nav() SelectorNode {
	return El("nav")
}

func Noscript() SelectorNode {
	return El("noscript")
}

func Object() SelectorNode {
	return El("object")
}

func Ol() SelectorNode {
	return El("ol")
}

func Optgroup() SelectorNode {
	return El("optgroup")
}

func Option() SelectorNode {
	return El("option")
}

func Output() SelectorNode {
	return El("output")
}

func P() SelectorNode {
	return El("p")
}

func Param() SelectorNode {
	return El("param")
}

func Picture() SelectorNode {
	return El("picture")
}

func Pre() SelectorNode {
	return El("pre")
}

func Progress() SelectorNode {
	return El("progress")
}

func Q() SelectorNode {
	return El("q")
}

func Rp() SelectorNode {
	return El("rp")
}

func Rt() SelectorNode {
	return El("rt")
}

func Ruby() SelectorNode {
	return El("ruby")
}

func S() SelectorNode {
	return El("s")
}

func Samp() SelectorNode {
	return El("samp")
}

func Script() SelectorNode {
	return El("script")
}

func Section() SelectorNode {
	return El("section")
}

func Select() SelectorNode {
	return El("select")
}

func Small() SelectorNode {
	return El("small")
}

func Source() SelectorNode {
	return El("source")
}

func Span() SelectorNode {
	return El("span")
}

func Strong() SelectorNode {
	return El("strong")
}

func Style() SelectorNode {
	return El("style")
}

func Sub() SelectorNode {
	return El("sub")
}

func Summary() SelectorNode {
	return El("summary")
}

func Sup() SelectorNode {
	return El("sup")
}

func Table() SelectorNode {
	return El("table")
}

func Tbody() SelectorNode {
	return El("tbody")
}

func Td() SelectorNode {
	return El("td")
}

func Template() SelectorNode {
	return El("template")
}

func Textarea() SelectorNode {
	return El("textarea")
}

func Tfoot() SelectorNode {
	return El("tfoot")
}

func Th() SelectorNode {
	return El("th")
}

func Thead() SelectorNode {
	return El("thead")
}

func Time() SelectorNode {
	return El("time")
}

func Title() SelectorNode {
	return El("title")
}

func Tr() SelectorNode {
	return El("tr")
}

func Track() SelectorNode {
	return El("track")
}

func U() SelectorNode {
	return El("u")
}

func Ul() SelectorNode {
	return El("ul")
}

func Var() SelectorNode {
	return El("var")
}

func Video() SelectorNode {
	return El("video")
}

func Wbr() SelectorNode {
	return El("wbr")
}
