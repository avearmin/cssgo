package gomour

import (
	"testing"
)

func TestPX(t *testing.T) {
	RunTests(t,
		test{"0px", PX(0), "0px"},
		test{"10px", PX(10), "10px"},
		test{"100px", PX(100), "100px"},
		test{"10.555px", PX(10.555), "10.555px"},
	)
}

func TestCM(t *testing.T) {
	RunTests(t,
		test{"0cm", CM(0), "0cm"},
		test{"10cm", CM(10), "10cm"},
		test{"100cm", CM(100), "100cm"},
		test{"10.555cm", CM(10.555), "10.555cm"},
	)
}

func TestMM(t *testing.T) {
	RunTests(t,
		test{"0mm", MM(0), "0mm"},
		test{"10mm", MM(10), "10mm"},
		test{"100mm", MM(100), "100mm"},
		test{"10.555mm", MM(10.555), "10.555mm"},
	)
}

func TestIN(t *testing.T) {
	RunTests(t,
		test{"0in", IN(0), "0in"},
		test{"10in", IN(10), "10in"},
		test{"100in", IN(100), "100in"},
		test{"10.555in", IN(10.555), "10.555in"},
	)
}

func TestPT(t *testing.T) {
	RunTests(t,
		test{"0pt", PT(0), "0pt"},
		test{"10pt", PT(10), "10pt"},
		test{"100pt", PT(100), "100pt"},
		test{"10.555pt", PT(10.555), "10.555pt"},
	)
}

func TestPC(t *testing.T) {
	RunTests(t,
		test{"0pc", PC(0), "0pc"},
		test{"10pc", PC(10), "10pc"},
		test{"100pc", PC(100), "100pc"},
		test{"10.555pc", PC(10.555), "10.555pc"},
	)
}

func TestEM(t *testing.T) {
	RunTests(t,
		test{"0em", EM(0), "0em"},
		test{"10em", EM(10), "10em"},
		test{"100em", EM(100), "100em"},
		test{"10.555em", EM(10.555), "10.555em"},
	)
}

func TestREM(t *testing.T) {
	RunTests(t,
		test{"0rem", REM(0), "0rem"},
		test{"10rem", REM(10), "10rem"},
		test{"100rem", REM(100), "100rem"},
		test{"10.555rem", REM(10.555), "10.555rem"},
	)
}

func TestVW(t *testing.T) {
	RunTests(t,
		test{"0vw", VW(0), "0vw"},
		test{"10vw", VW(10), "10vw"},
		test{"100vw", VW(100), "100vw"},
		test{"10.555vw", VW(10.555), "10.555vw"},
	)
}

func TestVH(t *testing.T) {
	RunTests(t,
		test{"0vh", VH(0), "0vh"},
		test{"10vh", VH(10), "10vh"},
		test{"100vh", VH(100), "100vh"},
		test{"10.555vh", VH(10.555), "10.555vh"},
	)
}

func TestPCT(t *testing.T) {
	RunTests(t,
		test{"0%", PCT(0), "0%"},
		test{"10%", PCT(10), "10%"},
		test{"100%", PCT(100), "100%"},
		test{"10.555%", PCT(10.555), "10.555%"},
	)
}

func TestVMIN(t *testing.T) {
	RunTests(t,
		test{"0vmin", VMIN(0), "0vmin"},
		test{"10vmin", VMIN(10), "10vmin"},
		test{"100vmin", VMIN(100), "100vmin"},
		test{"10.555vmin", VMIN(10.555), "10.555vmin"},
	)
}

func TestVMAX(t *testing.T) {
	RunTests(t,
		test{"0vmax", VMAX(0), "0vmax"},
		test{"10vmax", VMAX(10), "10vmax"},
		test{"100vmax", VMAX(100), "100vmax"},
		test{"10.555vmax", VMAX(10.555), "10.555vmax"},
	)
}
