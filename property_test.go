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
		test{
			"intial color",
			TextColor(Initial),
			"color: initial;",
		},
		test{
			"inherit color",
			TextColor(Inherit),
			"color: inherit;",
		},
		test{
			"revert color",
			TextColor(Revert),
			"color: revert;",
		},
		test{
			"unset color",
			TextColor(Unset),
			"color: unset;",
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
		test{
			"intial color",
			BackgroundColor(Initial),
			"background-color: initial;",
		},
		test{
			"inherit color",
			BackgroundColor(Inherit),
			"background-color: inherit;",
		},
		test{
			"revert color",
			BackgroundColor(Revert),
			"background-color: revert;",
		},
		test{
			"unset color",
			BackgroundColor(Unset),
			"background-color: unset;",
		},
	)
}

func TestAccentColor(t *testing.T) {
	RunTests(t,
		test{
			"named color",
			AccentColor(Red),
			"accent-color: red;",
		},
		test{
			"rgb color",
			AccentColor(RGB(40, 50, 60)),
			"accent-color: rgb(40, 50, 60);",
		},
		test{
			"hex color",
			AccentColor(Hex(0xffffff)),
			"accent-color: #ffffff;",
		},
		test{
			"hsl color",
			AccentColor(HSL(200, 50, 50)),
			"accent-color: hsl(200, 50%, 50%);",
		},
		test{
			"hsla color",
			AccentColor(HSLA(200, 50, 50, 0.5)),
			"accent-color: hsla(200, 50%, 50%, 0.5);",
		},
		test{
			"intial color",
			AccentColor(Initial),
			"accent-color: initial;",
		},
		test{
			"inherit color",
			AccentColor(Inherit),
			"accent-color: inherit;",
		},
		test{
			"revert color",
			AccentColor(Revert),
			"accent-color: revert;",
		},
		test{
			"unset color",
			AccentColor(Unset),
			"accent-color: unset;",
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
		test{
			"intial fontize",
			FontSize(Initial),
			"font-size: initial;",
		},
		test{
			"inherit fontsize",
			FontSize(Inherit),
			"font-size: inherit;",
		},
		test{
			"revert fontsize",
			FontSize(Revert),
			"font-size: revert;",
		},
		test{
			"unset fontsize",
			FontSize(Unset),
			"font-size: unset;",
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

func TestBorderColor(t *testing.T) {
	RunTests(t,
		test{
			"border-color with 1 value",
			BorderColor1(Red),
			"border-color: red;",
		},
		test{
			"border-color with 2 values",
			BorderColor2(Red, Blue),
			"border-color: red blue;",
		},
		test{
			"border-color with 3 values",
			BorderColor3(Red, Blue, Green),
			"border-color: red blue green;",
		},
		test{
			"border-color with 4 values",
			BorderColor4(Red, Blue, Green, Purple),
			"border-color: red blue green purple;",
		},
	)
}

func TestBorderRadius(t *testing.T) {
	RunTests(t,
		test{
			"border-radius with 1 value",
			BorderRadius(PX(5)),
			"border-radius: 5px;",
		},
		test{
			"border-radius with 2 values",
			BorderRadius2(PX(5), PX(10)),
			"border-radius: 5px 10px;",
		},
		test{
			"border-radius with 3 values",
			BorderRadius3(PX(5), PX(10), PX(15)),
			"border-radius: 5px 10px 15px;",
		},
		test{
			"border-radius with 4 values",
			BorderRadius4(PX(5), PX(10), PX(15), PX(20)),
			"border-radius: 5px 10px 15px 20px;",
		},
	)
}

func TestBorderRadiusCorners(t *testing.T) {
	RunTests(t,
		test{
			"border-top-left-radius: 5px;",
			BorderTopLeftRadius(PX(5)),
			"border-top-left-radius: 5px;",
		},
		test{
			"border-top-right-radius: 10px;",
			BorderTopRightRadius(PX(10)),
			"border-top-right-radius: 10px;",
		},
		test{
			"border-bottom-right-radius: 15px;",
			BorderBottomRightRadius(PX(15)),
			"border-bottom-right-radius: 15px;",
		},
		test{
			"border-bottom-left-radius: 20px;",
			BorderBottomLeftRadius(PX(20)),
			"border-bottom-left-radius: 20px;",
		},
	)
}

func TestLineHeight(t *testing.T) {
	RunTests(t,
		test{
			"line-height in px",
			LineHeight(PX(16)),
			"line-height: 16px;",
		},
		test{
			"line-height in em",
			LineHeight(EM(1.5)),
			"line-height: 1.5em;",
		},
		test{
			"line-height in percent",
			LineHeight(PCT(120)),
			"line-height: 120%;",
		},
	)
}

func TestLetterSpacing(t *testing.T) {
	RunTests(t,
		test{
			"letter-spacing in px",
			LetterSpacing(PX(2)),
			"letter-spacing: 2px;",
		},
		test{
			"letter-spacing in em",
			LetterSpacing(EM(0.1)),
			"letter-spacing: 0.1em;",
		},
		test{
			"letter-spacing in percent",
			LetterSpacing(PCT(5)),
			"letter-spacing: 5%;",
		},
	)
}

func TestDisplay(t *testing.T) {
	RunTests(t,
		test{
			"Display block",
			Display(Block),
			"display: block;",
		},
		test{
			"Display inline",
			Display(Inline),
			"display: inline;",
		},
		test{
			"Display inline-block",
			Display(InlineBlock),
			"display: inline-block;",
		},
		test{
			"Display flex",
			Display(Flex),
			"display: flex;",
		},
		test{
			"Display grid",
			Display(Grid),
			"display: grid;",
		},
		test{
			"Display inline-flex",
			Display(InlineFlex),
			"display: inline-flex;",
		},
		test{
			"Display inline-grid",
			Display(InlineGrid),
			"display: inline-grid;",
		},
		test{
			"Display table",
			Display(Table),
			"display: table;",
		},
		test{
			"Display table-row",
			Display(TableRow),
			"display: table-row;",
		},
		test{
			"Display table-cell",
			Display(TableCell),
			"display: table-cell;",
		},
		test{
			"Display list-item",
			Display(ListItem),
			"display: list-item;",
		},
		test{
			"Display run-in",
			Display(RunIn),
			"display: run-in;",
		},
		test{
			"Display flow-root",
			Display(FlowRoot),
			"display: flow-root;",
		},
		test{
			"Display contents",
			Display(Contents),
			"display: contents;",
		},
		test{
			"Display table-column",
			Display(TableColumn),
			"display: table-column;",
		},
		test{
			"Display table-column-group",
			Display(TableColumnGroup),
			"display: table-column-group;",
		},
		test{
			"Display table-header-group",
			Display(TableHeaderGroup),
			"display: table-header-group;",
		},
		test{
			"Display table-footer-group",
			Display(TableFooterGroup),
			"display: table-footer-group;",
		},
		test{
			"Display table-row-group",
			Display(TableRowGroup),
			"display: table-row-group;",
		},
	)
}

func TestFlexDirection(t *testing.T) {
	RunTests(t,
		test{
			"FlexDirection row",
			FlexDirection(Row),
			"flex-direction: row;",
		},
		test{
			"FlexDirection row-reverse",
			FlexDirection(RowReverse),
			"flex-direction: row-reverse;",
		},
		test{
			"FlexDirection column",
			FlexDirection(Column),
			"flex-direction: column;",
		},
		test{
			"FlexDirection column-reverse",
			FlexDirection(ColumnReverse),
			"flex-direction: column-reverse;",
		},
	)
}

func TestZIndex(t *testing.T) {
	RunTests(t,
		test{
			"ZIndex -7",
			ZIndex(-7),
			"z-index: -7;",
		},
		test{
			"ZIndex 7",
			ZIndex(7),
			"z-index: 7;",
		},
		test{
			"ZIndex auto",
			ZIndex(Auto),
			"z-index: auto;",
		},
	)
}

func TestBackgroundImage(t *testing.T) {
	RunTests(t,
		test{
			"background-image test/example.png",
			BackgroundImage(Url("test/example.png")),
			"background-image: url('test/example.png');",
		},
		test{
			"intial background-image",
			BackgroundImage(Initial),
			"background-image: initial;",
		},
		test{
			"inherit background-image",
			BackgroundImage(Inherit),
			"background-image: inherit;",
		},
		test{
			"revert background-image",
			BackgroundImage(Revert),
			"background-image: revert;",
		},
		test{
			"unset background-image",
			BackgroundImage(Unset),
			"background-image: unset;",
		},
	)
}
