package gomour

import "testing"

func TestRule(t *testing.T) {
	tests := []struct {
		input RuleNodeFunc
		want  string
	}{
		{
			input: Rule(
				Class("foo"),
				TextColor(Red()),
				BackgroundColor(Hex(0xffffff)),
			),
			want: ".foo{color:red;background-color:#ffffff;}",
		},
	}

	for i, test := range tests {
		func(t2 *testing.T) {
			got := test.input.String()
			if got != test.want {
				t2.Fatalf("TESTCASE %d: FAIL\ngot: %s != want: %s", i, got, test.want)
			}
		}(t)
	}
}
