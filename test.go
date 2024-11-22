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
