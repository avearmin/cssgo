package gomour

import (
	"testing"
)

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
			Padding(PX(10), nil, nil, nil),
			"padding: 10px;",
		},
		test{
			"padding with 2 values",
			Padding(PX(10), PX(20), nil, nil),
			"padding: 10px 20px;",
		},
		test{
			"padding with 3 values",
			Padding(PX(10), PX(20), PX(30), nil),
			"padding: 10px 20px 30px;",
		},
		test{
			"padding with 4 values",
			Padding(PX(10), PX(20), PX(30), PX(40)),
			"padding: 10px 20px 30px 40px;",
		},
	)
}

func TestMargin(t *testing.T) {
	RunTests(t,
		test{
			"margin with 1 value",
			Margin(PX(10), nil, nil, nil),
			"margin: 10px;",
		},
		test{
			"margin with 2 values",
			Margin(PX(10), PX(20), nil, nil),
			"margin: 10px 20px;",
		},
		test{
			"margin with 3 values",
			Margin(PX(10), PX(20), PX(30), nil),
			"margin: 10px 20px 30px;",
		},
		test{
			"margin with 4 values",
			Margin(PX(10), PX(20), PX(30), PX(40)),
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
			"border: 5px solid green;",
			Border(PX(5), Solid, Green),
			"border: 5px solid green;",
		},
		test{
			"border: dotted #ffffff;",
			Border(nil, Dotted, Hex(0xffffff)),
			"border: dotted #ffffff;",
		},
		test{
			"border: dashed;",
			Border(nil, Dashed, nil),
			"border: dashed;",
		},
	)
}
