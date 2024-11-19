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
