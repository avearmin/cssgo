package gomour

import "testing"

func TestPadding(t *testing.T) {
	tests := []struct {
		name  string
		input PropertyNodeFunc
		want  string
	}{
		{
			"padding with 1 value",
			Padding(PX(10), nil, nil, nil),
			"padding: 10px;",
		},
		{
			"padding with 2 values",
			Padding(PX(10), PX(20), nil, nil),
			"padding: 10px 20px;",
		},
		{
			"padding with 3 values",
			Padding(PX(10), PX(20), PX(30), nil),
			"padding: 10px 20px 30px;",
		},
		{
			"padding with 3 values",
			Padding(PX(10), PX(20), PX(30), PX(40)),
			"padding: 10px 20px 30px 40px;",
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
