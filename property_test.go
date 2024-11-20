package gomour

import (
	"strings"
	"testing"
)

type test struct {
	name  string
	input Node
	want  string
}

func RunTests(t *testing.T, tests ...test) {
	for _, te := range tests {
		func(t2 *testing.T) {
			w := strings.Builder{}
			te.input.RenderCSS(&w)
			got := w.String()
			if got != te.want {
				t2.Fatalf("TESTCASE %s: FAIL\ngot: %s != want: %s", te.name, got, te.want)
			}
		}(t)
	}
}

func TestTextColor(t *testing.T) {
	RunTests(t,
		test{"named color",
			TextColor(Color("red")),
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
			TextColor(HSL(200, .5, .5)),
			"color: hsl(200, 50%, 50%);",
		},
		test{
			"hsla color",
			TextColor(HSLA(200, .5, .5, .5)),
			"color: hsla(200, 50%, 50%, 0.5);",
		},
	)
}

func TestBackgroundColor(t *testing.T) {
	RunTests(t,
		test{"named color",
			BackgroundColor(Color("red")),
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
			BackgroundColor(HSL(200, .5, .5)),
			"background-color: hsl(200, 50%, 50%);",
		},
		test{
			"hsla color",
			BackgroundColor(HSLA(200, .5, .5, .5)),
			"background-color: hsla(200, 50%, 50%, 0.5);",
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
