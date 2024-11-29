package cssgo

import (
	"testing"
)

func TestBorderStyleConstants(t *testing.T) {
	RunTests(t,
		test{
			"solid constant",
			Solid,
			"solid",
		},
		test{
			"dotted constant",
			Dotted,
			"dotted",
		},
		test{
			"dashed constant",
			Dashed,
			"dashed",
		},
		test{
			"double constant",
			Double,
			"double",
		},
		test{
			"groove constant",
			Groove,
			"groove",
		},
		test{
			"ridge constant",
			Ridge,
			"ridge",
		},
		test{
			"inset constant",
			Inset,
			"inset",
		},
		test{
			"outset constant",
			Outset,
			"outset",
		},
	)
}
