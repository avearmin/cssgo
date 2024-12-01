package cssgo

import (
	"testing"
)

func TestGroupProps(t *testing.T) {
	RunTests(t,
		test{"border background-color color",
			GroupProps(
				Border3(PX(10), Inset, Magenta),
				BackgroundColor(Black),
				TextColor(Red),
			),
			"border: 10px inset magenta;background-color: black;color: red;",
		},
	)
}

func TestTextColor(t *testing.T) {
	RunTests(t,
		test{"named color",
			TextColor(Red),
			"color: red;",
		},
		test{
			"rgb color",
			TextColor(RGB(40, 50, 60)),
			"color: rgb(40, 50, 60);",
		},
		test{
			"hex color",
			TextColor(Hex(0xffffff)),
			"color: #ffffff;",
		},
		test{
			"hsl color",
			TextColor(HSL(200, 50, 50)),
			"color: hsl(200, 50%, 50%);",
		},
		test{
			"hsla color",
			TextColor(HSLA(200, 50, 50, .5)),
			"color: hsla(200, 50%, 50%, 0.5);",
		},
	)
}

func TestBackgroundColor(t *testing.T) {
	RunTests(t,
		test{"named color",
			BackgroundColor(Red),
			"background-color: red;",
		},
		test{
			"rgb color",
			BackgroundColor(RGB(40, 50, 60)),
			"background-color: rgb(40, 50, 60);",
		},
		test{
			"hex color",
			BackgroundColor(Hex(0xffffff)),
			"background-color: #ffffff;",
		},
		test{
			"hsl color",
			BackgroundColor(HSL(200, 50, 50)),
			"background-color: hsl(200, 50%, 50%);",
		},
		test{
			"hsla color",
			BackgroundColor(HSLA(200, 50, 50, .5)),
			"background-color: hsla(200, 50%, 50%, 0.5);",
		},
	)
}

func TestFontSize(t *testing.T) {
	RunTests(t,
		test{
			"fontsize with px value",
			FontSize(PX(10)),
			"font-size: 10px;",
		},
		test{
			"fontsize with percent value",
			FontSize(PCT(50)),
			"font-size: 50%;",
		},
		test{
			"fontsize with em value",
			FontSize(EM(1.5)),
			"font-size: 1.5em;",
		},
		test{
			"fontsize with rem value",
			FontSize(REM(2)),
			"font-size: 2rem;",
		},
		test{
			"fontsize with in value",
			FontSize(IN(1)),
			"font-size: 1in;",
		},
		test{
			"fontsize with cm value",
			FontSize(CM(2.54)),
			"font-size: 2.54cm;",
		},
		test{
			"fontsize with mm value",
			FontSize(MM(10)),
			"font-size: 10mm;",
		},
		test{
			"fontsize with pt value",
			FontSize(PT(12)),
			"font-size: 12pt;",
		},
		test{
			"fontsize with pc value",
			FontSize(PC(1)),
			"font-size: 1pc;",
		},
		test{
			"fontsize with vw value",
			FontSize(VW(25.5)),
			"font-size: 25.5vw;",
		},
		test{
			"fontsize with vh value",
			FontSize(VH(10)),
			"font-size: 10vh;",
		},
		test{
			"fontsize with vmin value",
			FontSize(VMIN(5)),
			"font-size: 5vmin;",
		},
		test{
			"fontsize with vmax value",
			FontSize(VMAX(8)),
			"font-size: 8vmax;",
		},
	)
}

func TestPadding(t *testing.T) {
	RunTests(t,
		test{
			"padding with 1 value",
			Padding1(PX(10)),
			"padding: 10px;",
		},
		test{
			"padding with 2 values",
			Padding2(PX(10), PX(20)),
			"padding: 10px 20px;",
		},
		test{
			"padding with 3 values",
			Padding3(PX(10), PX(20), PX(30)),
			"padding: 10px 20px 30px;",
		},
		test{
			"padding with 4 values",
			Padding4(PX(10), PX(20), PX(30), PX(40)),
			"padding: 10px 20px 30px 40px;",
		},
	)
}

func TestPaddingDirections(t *testing.T) {
	RunTests(t,
		test{
			"padding-top: 10px",
			PaddingTop(PX(10)),
			"padding-top: 10px;",
		},
		test{
			"padding-right: 20px",
			PaddingRight(PX(20)),
			"padding-right: 20px;",
		},
		test{
			"padding-bottom: 30px",
			PaddingBottom(PX(30)),
			"padding-bottom: 30px;",
		},
		test{
			"padding-left: 40px",
			PaddingLeft(PX(40)),
			"padding-left: 40px;",
		},
	)
}

func TestMargin(t *testing.T) {
	RunTests(t,
		test{
			"margin with 1 value",
			Margin1(PX(10)),
			"margin: 10px;",
		},
		test{
			"margin with 2 values",
			Margin2(PX(10), PX(20)),
			"margin: 10px 20px;",
		},
		test{
			"margin with 3 values",
			Margin3(PX(10), PX(20), PX(30)),
			"margin: 10px 20px 30px;",
		},
		test{
			"margin with 4 values",
			Margin4(PX(10), PX(20), PX(30), PX(40)),
			"margin: 10px 20px 30px 40px;",
		},
	)
}

func TestHeight(t *testing.T) {
	RunTests(t,
		test{
			"height: 100px",
			Height(PX(100)),
			"height: 100px;",
		},
		test{
			"height: 50%",
			Height(PCT(50)),
			"height: 50%;",
		},
		test{
			"height: 10em",
			Height(EM(10)),
			"height: 10em;",
		},
	)
}

func TestWidth(t *testing.T) {
	RunTests(t,
		test{
			"width: 200px",
			Width(PX(200)),
			"width: 200px;",
		},
		test{
			"width: 75%",
			Width(PCT(75)),
			"width: 75%;",
		},
		test{
			"width: 5rem",
			Width(REM(5)),
			"width: 5rem;",
		},
	)
}

func TestBorder(t *testing.T) {
	RunTests(t,
		test{
			"border with only width",
			Border1(PX(5)),
			"border: 5px;",
		},
		test{
			"border with width and style",
			Border2(PX(5), Solid),
			"border: 5px solid;",
		},
		test{
			"border with width, style, and color",
			Border3(PX(5), Solid, Green),
			"border: 5px solid green;",
		},
	)
}

func TestBorderStyle(t *testing.T) {
	RunTests(t,
		test{
			"border-style with 1 value",
			BorderStyle1(Dotted),
			"border-style: dotted;",
		},
		test{
			"border-style with 2 values",
			BorderStyle2(Dotted, Dashed),
			"border-style: dotted dashed;",
		},
		test{
			"border-style with 3 values",
			BorderStyle3(Dotted, Dashed, Double),
			"border-style: dotted dashed double;",
		},
		test{
			"border-style with 4 values",
			BorderStyle4(Dotted, Dashed, Double, Solid),
			"border-style: dotted dashed double solid;",
		},
	)
}

func TestBorderWidth(t *testing.T) {
	RunTests(t,
		test{
			"border-width: 10px;",
			BorderWidth1(PX(10)),
			"border-width: 10px;",
		},
		test{
			"border-width: thick;",
			BorderWidth1(Thick),
			"border-width: thick;",
		},
		test{
			"border-width: 10px 20px;",
			BorderWidth2(PX(10), PX(20)),
			"border-width: 10px 20px;",
		},
		test{
			"border-width: thick thin;",
			BorderWidth2(Thick, Thin),
			"border-width: thick thin;",
		},
		test{
			"border-width: 10px 20px 30px;",
			BorderWidth3(PX(10), PX(20), PX(30)),
			"border-width: 10px 20px 30px;",
		},
		test{
			"border-width: thick thin medium;",
			BorderWidth3(Thick, Thin, Medium),
			"border-width: thick thin medium;",
		},
		test{
			"border-width: 10px 20px 30px 40px;",
			BorderWidth4(PX(10), PX(20), PX(30), PX(40)),
			"border-width: 10px 20px 30px 40px;",
		},
		test{
			"border-width: thick thin medium thin;",
			BorderWidth4(Thick, Thin, Medium, Thin),
			"border-width: thick thin medium thin;",
		},
	)
}

func TestBorderDirectionWidths(t *testing.T) {
	RunTests(t,
		test{
			"border-left-width: 10px;",
			BorderLeftWidth(PX(10)),
			"border-left-width: 10px;",
		},
		test{
			"border-left-width: thick;",
			BorderLeftWidth(Thick),
			"border-left-width: thick;",
		},
		test{
			"border-right-width: medium;",
			BorderRightWidth(Medium),
			"border-right-width: medium;",
		},
		test{
			"border-right-width: 20px;",
			BorderRightWidth(PX(20)),
			"border-right-width: 20px;",
		},
		test{
			"border-top-width: 15px;",
			BorderTopWidth(PX(15)),
			"border-top-width: 15px;",
		},
		test{
			"border-top-width: thin;",
			BorderTopWidth(Thin),
			"border-top-width: thin;",
		},
		test{
			"border-bottom-width: 5px;",
			BorderBottomWidth(PX(5)),
			"border-bottom-width: 5px;",
		},
		test{
			"border-bottom-width: thick;",
			BorderBottomWidth(Thick),
			"border-bottom-width: thick;",
		},
	)
}
