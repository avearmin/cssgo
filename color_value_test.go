package gomour

import "testing"

func TestHex(t *testing.T) {
	RunTests(t,
		test{"Hex: #000000", Hex(0), "#000000"},
		test{"Hex: #ffffff", Hex(0xFFFFFF), "#ffffff"},
		test{"Hex: #123abc", Hex(0x123ABC), "#123abc"},
	)
}

func TestRGB(t *testing.T) {
	RunTests(t,
		test{"RGB: rgb(0, 0, 0)", RGB(0, 0, 0), "rgb(0, 0, 0)"},
		test{"RGB: rgb(255, 255, 255)", RGB(255, 255, 255), "rgb(255, 255, 255)"},
		test{"RGB: rgb(123, 45, 67)", RGB(123, 45, 67), "rgb(123, 45, 67)"},
	)
}

func TestRGBA(t *testing.T) {
	RunTests(t,
		test{"RGBA: rgba(0, 0, 0, 0.000000)", RGBA(0, 0, 0, 0), "rgba(0, 0, 0, 0.000000)"},
		test{"RGBA: rgba(255, 255, 255, 1.000000)", RGBA(255, 255, 255, 1), "rgba(255, 255, 255, 1.000000)"},
		test{"RGBA: rgba(123, 45, 67, 0.500000)", RGBA(123, 45, 67, 0.5), "rgba(123, 45, 67, 0.500000)"},
	)
}

func TestColor(t *testing.T) {
	RunTests(t,
		test{"Color: red", Color("red"), "red"},
		test{"Color: blue", Color("blue"), "blue"},
		test{"Color: green", Color("green"), "green"},
		test{"Color: transparent", Color("transparent"), "transparent"},
		test{"Color: invalid", Color("invalid"), "invalid"}, // Outputs as-is, no validation
	)
}

func TestHSL(t *testing.T) {
	RunTests(t,
		test{"HSL: hsl(0, 0%, 0%)", HSL(0, 0, 0), "hsl(0, 0%, 0%)"},
		test{"HSL: hsl(360, 100%, 50%)", HSL(360, 1, 0.5), "hsl(360, 100%, 50%)"},
		test{"HSL: hsl(180, 50%, 25%)", HSL(180, 0.5, 0.25), "hsl(180, 50%, 25%)"},
	)
}

func TestHSLA(t *testing.T) {
	RunTests(t,
		test{"HSLA: hsla(0, 0%, 0%, 0.0)", HSLA(0, 0, 0, 0), "hsla(0, 0%, 0%, 0.0)"},
		test{"HSLA: hsla(360, 100%, 50%, 1.0)", HSLA(360, 1, 0.5, 1), "hsla(360, 100%, 50%, 1.0)"},
		test{"HSLA: hsla(180, 50%, 25%, 0.5)", HSLA(180, 0.5, 0.25, 0.5), "hsla(180, 50%, 25%, 0.5)"},
	)
}
