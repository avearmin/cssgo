package gomour

import (
	"testing"
)

func TestPX(t *testing.T) {
	RunTests(t,
		test{
			"0px",
			PX(0),
			"0px",
		},
		test{
			"10px",
			PX(10),
			"10px",
		},
		test{
			"100px",
			PX(100),
			"100px",
		},
		test{
			"10.555px",
			PX(10.555),
			"10.555px",
		},
	)
}
