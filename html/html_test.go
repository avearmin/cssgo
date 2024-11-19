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
			input: ghtml.Head(
				StyleEl(
					gomour.Class("foo").Props(
						gomour.TextColor(gomour.Color("red")),
						gomour.BackgroundColor(gomour.Hex(0xffffff)),
					),
				),
			),
			want: "<head><style>.foo{color: red;background-color: #ffffff;}</style></head>",
		},
		{
			name: "simple class selector in style attr",
			input: ghtml.Div(
				Style(
					gomour.TextColor(gomour.Color("red")),
					gomour.BackgroundColor(gomour.Hex(0xffffff)),
				),
				gomponents.Text("hello world"),
			),
			want: "<div style=\"color: red;background-color: #ffffff;\">hello world</div>",
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
