package gomour

import "testing"

func TestRule(t *testing.T) {
	tests := []struct {
		name  string
		input RuleNodeFunc
		want  string
	}{
		{
			name: "simple class selector",
			input: Class("foo").Props(
				TextColor(Color("red")),
				BackgroundColor(Hex(0xffffff)),
				FontSize(REM(40)),
			),
			want: ".foo{color:red;background-color:#ffffff;font-size:40rem;}",
		},
		{
			name: "comma sep class selector",
			input: Class("foo").Or(El("p")).Props(
				TextColor(Color("red")),
				BackgroundColor(Hex(0xffffff)),
				FontSize(EM(50)),
			),
			want: ".foo, p{color:red;background-color:#ffffff;font-size:50em;}",
		},
		{
			name: "complex comma sep class selector",
			input: Class("foo").Or(El("p").Or(El("a"))).Or(ID("baz")).Props(
				TextColor(HSL(400, .844, .95777)),
				BackgroundColor(HSLA(200, .70, .60, 1.5)),
				FontSize(PX(20)),
			),
			want: ".foo, p, a, #baz{color:hsl(360, 84%, 96%);background-color:hsla(200, 70%, 60%, 1.0);font-size:20px;}",
		},
	}

	for _, test := range tests {
		func(t2 *testing.T) {
			got := test.input.String()
			if got != test.want {
				t2.Fatalf("TESTCASE %s: FAIL\ngot: %s != want: %s", test.name, got, test.want)
			}
		}(t)
	}
}
