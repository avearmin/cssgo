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

func (s SelectorNodeFunc) RenderCSS(w io.Writer) error {
	return s(w)
}

func (s SelectorNodeFunc) String() string {
	var b strings.Builder
	_ = s.RenderCSS(&b)
	return b.String()
}

func (s SelectorNodeFunc) selectorNode() {}

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

func A() SelectorNodeFunc {
	return El("a")
}

func Abbr() SelectorNodeFunc {
	return El("abbr")
}

func Address() SelectorNodeFunc {
	return El("address")
}

func Area() SelectorNodeFunc {
	return El("area")
}

func Article() SelectorNodeFunc {
	return El("article")
}

func Aside() SelectorNodeFunc {
	return El("aside")
}

func Audio() SelectorNodeFunc {
	return El("audio")
}

func B() SelectorNodeFunc {
	return El("b")
}

func Base() SelectorNodeFunc {
	return El("base")
}

func Bdi() SelectorNodeFunc {
	return El("bdi")
}

func Bdo() SelectorNodeFunc {
	return El("bdo")
}

func Blockquote() SelectorNodeFunc {
	return El("blockquote")
}

func Body() SelectorNodeFunc {
	return El("body")
}

func Br() SelectorNodeFunc {
	return El("br")
}

func Button() SelectorNodeFunc {
	return El("button")
}

func Canvas() SelectorNodeFunc {
	return El("canvas")
}

func Caption() SelectorNodeFunc {
	return El("caption")
}

func Cite() SelectorNodeFunc {
	return El("cite")
}

func Code() SelectorNodeFunc {
	return El("code")
}

func Col() SelectorNodeFunc {
	return El("col")
}

func Colgroup() SelectorNodeFunc {
	return El("colgroup")
}

func Data() SelectorNodeFunc {
	return El("data")
}

func Datalist() SelectorNodeFunc {
	return El("datalist")
}

func Dd() SelectorNodeFunc {
	return El("dd")
}

func Del() SelectorNodeFunc {
	return El("del")
}

func Details() SelectorNodeFunc {
	return El("details")
}

func Dfn() SelectorNodeFunc {
	return El("dfn")
}

func Dialog() SelectorNodeFunc {
	return El("dialog")
}

func Div() SelectorNodeFunc {
	return El("div")
}

func Dl() SelectorNodeFunc {
	return El("dl")
}

func Dt() SelectorNodeFunc {
	return El("dt")
}

func Em() SelectorNodeFunc {
	return El("em")
}

func Embed() SelectorNodeFunc {
	return El("embed")
}

func Fieldset() SelectorNodeFunc {
	return El("fieldset")
}

func Figcaption() SelectorNodeFunc {
	return El("figcaption")
}

func Figure() SelectorNodeFunc {
	return El("figure")
}

func Footer() SelectorNodeFunc {
	return El("footer")
}

func Form() SelectorNodeFunc {
	return El("form")
}

func H1() SelectorNodeFunc {
	return El("h1")
}

func H2() SelectorNodeFunc {
	return El("h2")
}

func H3() SelectorNodeFunc {
	return El("h3")
}

func H4() SelectorNodeFunc {
	return El("h4")
}

func H5() SelectorNodeFunc {
	return El("h5")
}

func H6() SelectorNodeFunc {
	return El("h6")
}

func Head() SelectorNodeFunc {
	return El("head")
}

func Header() SelectorNodeFunc {
	return El("header")
}

func Hr() SelectorNodeFunc {
	return El("hr")
}

func Html() SelectorNodeFunc {
	return El("html")
}

func I() SelectorNodeFunc {
	return El("i")
}

func Iframe() SelectorNodeFunc {
	return El("iframe")
}

func Img() SelectorNodeFunc {
	return El("img")
}

func Input() SelectorNodeFunc {
	return El("input")
}

func Ins() SelectorNodeFunc {
	return El("ins")
}

func Kbd() SelectorNodeFunc {
	return El("kbd")
}

func Label() SelectorNodeFunc {
	return El("label")
}

func Legend() SelectorNodeFunc {
	return El("legend")
}

func Li() SelectorNodeFunc {
	return El("li")
}

func Link() SelectorNodeFunc {
	return El("link")
}

func Main() SelectorNodeFunc {
	return El("main")
}

func Map() SelectorNodeFunc {
	return El("map")
}

func Mark() SelectorNodeFunc {
	return El("mark")
}

func Meta() SelectorNodeFunc {
	return El("meta")
}

func Meter() SelectorNodeFunc {
	return El("meter")
}

func Nav() SelectorNodeFunc {
	return El("nav")
}

func Noscript() SelectorNodeFunc {
	return El("noscript")
}

func Object() SelectorNodeFunc {
	return El("object")
}

func Ol() SelectorNodeFunc {
	return El("ol")
}

func Optgroup() SelectorNodeFunc {
	return El("optgroup")
}

func Option() SelectorNodeFunc {
	return El("option")
}

func Output() SelectorNodeFunc {
	return El("output")
}

func P() SelectorNodeFunc {
	return El("p")
}

func Param() SelectorNodeFunc {
	return El("param")
}

func Picture() SelectorNodeFunc {
	return El("picture")
}

func Pre() SelectorNodeFunc {
	return El("pre")
}

func Progress() SelectorNodeFunc {
	return El("progress")
}

func Q() SelectorNodeFunc {
	return El("q")
}

func Rp() SelectorNodeFunc {
	return El("rp")
}

func Rt() SelectorNodeFunc {
	return El("rt")
}

func Ruby() SelectorNodeFunc {
	return El("ruby")
}

func S() SelectorNodeFunc {
	return El("s")
}

func Samp() SelectorNodeFunc {
	return El("samp")
}

func Script() SelectorNodeFunc {
	return El("script")
}

func Section() SelectorNodeFunc {
	return El("section")
}

func Select() SelectorNodeFunc {
	return El("select")
}

func Small() SelectorNodeFunc {
	return El("small")
}

func Source() SelectorNodeFunc {
	return El("source")
}

func Span() SelectorNodeFunc {
	return El("span")
}

func Strong() SelectorNodeFunc {
	return El("strong")
}

func Style() SelectorNodeFunc {
	return El("style")
}

func Sub() SelectorNodeFunc {
	return El("sub")
}

func Summary() SelectorNodeFunc {
	return El("summary")
}

func Sup() SelectorNodeFunc {
	return El("sup")
}

func Table() SelectorNodeFunc {
	return El("table")
}

func Tbody() SelectorNodeFunc {
	return El("tbody")
}

func Td() SelectorNodeFunc {
	return El("td")
}

func Template() SelectorNodeFunc {
	return El("template")
}

func Textarea() SelectorNodeFunc {
	return El("textarea")
}

func Tfoot() SelectorNodeFunc {
	return El("tfoot")
}

func Th() SelectorNodeFunc {
	return El("th")
}

func Thead() SelectorNodeFunc {
	return El("thead")
}

func Time() SelectorNodeFunc {
	return El("time")
}

func Title() SelectorNodeFunc {
	return El("title")
}

func Tr() SelectorNodeFunc {
	return El("tr")
}

func Track() SelectorNodeFunc {
	return El("track")
}

func U() SelectorNodeFunc {
	return El("u")
}

func Ul() SelectorNodeFunc {
	return El("ul")
}

func Var() SelectorNodeFunc {
	return El("var")
}

func Video() SelectorNodeFunc {
	return El("video")
}

func Wbr() SelectorNodeFunc {
	return El("wbr")
}
