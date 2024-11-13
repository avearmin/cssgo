package html

import (
	"github.com/avearmin/gomour"
	"maragu.dev/gomponents"
	ghtml "maragu.dev/gomponents/html"
	"strings"
	"testing"
)

func TestStyles(t *testing.T) {
	tests := []struct {
		name  string
		input gomponents.Node
		want  string
	}{
		{
			name: "simple class selector in style element",
			input: ghtml.Div(
				StyleEl(
					gomour.Class("foo").Props(
						gomour.TextColor(gomour.Color("red")),
						gomour.BackgroundColor(gomour.Hex(0xffffff)),
					),
				),
			),
			want: "<div><style>.foo{color:red;background-color:#ffffff;}</style></div>",
		},
		{
			name: "simple class selector in style attr",
			input: Style(
				gomour.TextColor(gomour.Color("red")),
				gomour.BackgroundColor(gomour.Hex(0xffffff)),
			),
			want: "style=\"color:red;background-color:#ffffff;\"",
		},
	}

	for _, test := range tests {
		func(t2 *testing.T) {
			w := strings.Builder{}
			test.input.Render(&w)
			got := w.String()
			if got != test.want {
				t2.Fatalf("TESTCASE %s: FAIL\ngot: %s != want: %s", test.name, got, test.want)
			}
		}(t)
	}
}
