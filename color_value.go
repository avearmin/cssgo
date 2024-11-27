package gomour

import (
	"fmt"
	"io"
	"strings"
)

type ColorValue interface {
	ValueNode
	colorValue()
}

type namedColor string

const (
	AliceBlue            namedColor = "aliceblue"
	AntiqueWhite         namedColor = "antiquewhite"
	Aqua                 namedColor = "aqua"
	Aquamarine           namedColor = "aquamarine"
	Azure                namedColor = "azure"
	Beige                namedColor = "beige"
	Bisque               namedColor = "bisque"
	Black                namedColor = "black"
	BlanchedAlmond       namedColor = "blanchedalmond"
	Blue                 namedColor = "blue"
	BlueViolet           namedColor = "blueviolet"
	Brown                namedColor = "brown"
	BurlyWood            namedColor = "burlywood"
	CadetBlue            namedColor = "cadetblue"
	Chartreuse           namedColor = "chartreuse"
	Chocolate            namedColor = "chocolate"
	Coral                namedColor = "coral"
	CornflowerBlue       namedColor = "cornflowerblue"
	Cornsilk             namedColor = "cornsilk"
	Crimson              namedColor = "crimson"
	Cyan                 namedColor = "cyan"
	DarkBlue             namedColor = "darkblue"
	DarkCyan             namedColor = "darkcyan"
	DarkGoldenRod        namedColor = "darkgoldenrod"
	DarkGray             namedColor = "darkgray"
	DarkGreen            namedColor = "darkgreen"
	DarkKhaki            namedColor = "darkkhaki"
	DarkMagenta          namedColor = "darkmagenta"
	DarkOliveGreen       namedColor = "darkolivegreen"
	DarkOrange           namedColor = "darkorange"
	DarkOrchid           namedColor = "darkorchid"
	DarkRed              namedColor = "darkred"
	DarkSalmon           namedColor = "darksalmon"
	DarkSeaGreen         namedColor = "darkseagreen"
	DarkSlateBlue        namedColor = "darkslateblue"
	DarkSlateGray        namedColor = "darkslategray"
	DarkTurquoise        namedColor = "darkturquoise"
	DarkViolet           namedColor = "darkviolet"
	DeepPink             namedColor = "deeppink"
	DeepSkyBlue          namedColor = "deepskyblue"
	DimGray              namedColor = "dimgray"
	DodgerBlue           namedColor = "dodgerblue"
	FireBrick            namedColor = "firebrick"
	FloralWhite          namedColor = "floralwhite"
	ForestGreen          namedColor = "forestgreen"
	Fuchsia              namedColor = "fuchsia"
	Gainsboro            namedColor = "gainsboro"
	GhostWhite           namedColor = "ghostwhite"
	Gold                 namedColor = "gold"
	GoldenRod            namedColor = "goldenrod"
	Gray                 namedColor = "gray"
	Green                namedColor = "green"
	GreenYellow          namedColor = "greenyellow"
	HoneyDew             namedColor = "honeydew"
	HotPink              namedColor = "hotpink"
	IndianRed            namedColor = "indianred"
	Indigo               namedColor = "indigo"
	Ivory                namedColor = "ivory"
	Khaki                namedColor = "khaki"
	Lavender             namedColor = "lavender"
	LavenderBlush        namedColor = "lavenderblush"
	LawnGreen            namedColor = "lawngreen"
	LemonChiffon         namedColor = "lemonchiffon"
	LightBlue            namedColor = "lightblue"
	LightCoral           namedColor = "lightcoral"
	LightCyan            namedColor = "lightcyan"
	LightGoldenRodYellow namedColor = "lightgoldenrodyellow"
	LightGray            namedColor = "lightgray"
	LightGreen           namedColor = "lightgreen"
	LightPink            namedColor = "lightpink"
	LightSalmon          namedColor = "lightsalmon"
	LightSeaGreen        namedColor = "lightseagreen"
	LightSkyBlue         namedColor = "lightskyblue"
	LightSlateGray       namedColor = "lightslategray"
	LightSteelBlue       namedColor = "lightsteelblue"
	LightYellow          namedColor = "lightyellow"
	Lime                 namedColor = "lime"
	LimeGreen            namedColor = "limegreen"
	Linen                namedColor = "linen"
	Magenta              namedColor = "magenta"
	Maroon               namedColor = "maroon"
	MediumAquaMarine     namedColor = "mediumaquamarine"
	MediumBlue           namedColor = "mediumblue"
	MediumOrchid         namedColor = "mediumorchid"
	MediumPurple         namedColor = "mediumpurple"
	MediumSeaGreen       namedColor = "mediumseagreen"
	MediumSlateBlue      namedColor = "mediumslateblue"
	MediumSpringGreen    namedColor = "mediumspringgreen"
	MediumTurquoise      namedColor = "mediumturquoise"
	MediumVioletRed      namedColor = "mediumvioletred"
	MidnightBlue         namedColor = "midnightblue"
	MintCream            namedColor = "mintcream"
	MistyRose            namedColor = "mistyrose"
	Moccasin             namedColor = "moccasin"
	NavajoWhite          namedColor = "navajowhite"
	Navy                 namedColor = "navy"
	OldLace              namedColor = "oldlace"
	Olive                namedColor = "olive"
	OliveDrab            namedColor = "olivedrab"
	Orange               namedColor = "orange"
	OrangeRed            namedColor = "orangered"
	Orchid               namedColor = "orchid"
	PaleGoldenRod        namedColor = "palegoldenrod"
	PaleGreen            namedColor = "palegreen"
	PaleTurquoise        namedColor = "paleturquoise"
	PaleVioletRed        namedColor = "palevioletred"
	PapayaWhip           namedColor = "papayawhip"
	PeachPuff            namedColor = "peachpuff"
	Peru                 namedColor = "peru"
	Pink                 namedColor = "pink"
	Plum                 namedColor = "plum"
	PowderBlue           namedColor = "powderblue"
	Purple               namedColor = "purple"
	RebeccaPurple        namedColor = "rebeccapurple"
	Red                  namedColor = "red"
	RosyBrown            namedColor = "rosybrown"
	RoyalBlue            namedColor = "royalblue"
	SaddleBrown          namedColor = "saddlebrown"
	Salmon               namedColor = "salmon"
	SandyBrown           namedColor = "sandybrown"
	SeaGreen             namedColor = "seagreen"
	SeaShell             namedColor = "seashell"
	Sienna               namedColor = "sienna"
	Silver               namedColor = "silver"
	SkyBlue              namedColor = "skyblue"
	SlateBlue            namedColor = "slateblue"
	SlateGray            namedColor = "slategray"
	Snow                 namedColor = "snow"
	SpringGreen          namedColor = "springgreen"
	SteelBlue            namedColor = "steelblue"
	Tan                  namedColor = "tan"
	Teal                 namedColor = "teal"
	Thistle              namedColor = "thistle"
	Tomato               namedColor = "tomato"
	Turquoise            namedColor = "turquoise"
	Violet               namedColor = "violet"
	Wheat                namedColor = "wheat"
	White                namedColor = "white"
	WhiteSmoke           namedColor = "whitesmoke"
	Yellow               namedColor = "yellow"
	YellowGreen          namedColor = "yellowgreen"
)

func (n namedColor) RenderCSS(w io.Writer) error {
	_, err := w.Write([]byte(n))
	return err
}

func (n namedColor) valueNode()  {}
func (n namedColor) colorValue() {}

type ColorValueFunc func(io.Writer) error

func (f ColorValueFunc) RenderCSS(w io.Writer) error {
	return f(w)
}

func (f ColorValueFunc) String() string {
	var b strings.Builder
	_ = f.RenderCSS(&b)
	return b.String()
}

func (f ColorValueFunc) valueNode()  {}
func (f ColorValueFunc) colorValue() {}

func Hex(hexdec int) ColorValueFunc {
	return ColorValueFunc(func(w io.Writer) error {
		hexStr := fmt.Sprintf("#%06x", hexdec)
		_, err := w.Write([]byte(hexStr))
		return err
	})
}

func RGB(r, g, b int) ColorValueFunc {
	return ColorValueFunc(func(w io.Writer) error {
		rgbStr := fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
		_, err := w.Write([]byte(rgbStr))
		return err
	})
}

func RGBA(r, g, b int, alpha float64) ColorValueFunc {
	return ColorValueFunc(func(w io.Writer) error {
		rgbaStr := fmt.Sprintf("rgba(%d, %d, %d, %f)", r, g, b, alpha)
		_, err := w.Write([]byte(rgbaStr))
		return err
	})
}

func HSL(hue uint, saturation, lightness float64) ColorValueFunc {
	if hue > 360 {
		hue = 360
	}
	if saturation < 0.0 {
		saturation = 0.0
	}
	if saturation > 1.0 {
		saturation = 1.0
	}
	if lightness < 0.0 {
		lightness = 0.0
	}
	if lightness > 1.0 {
		lightness = 1.0
	}

	return ColorValueFunc(func(w io.Writer) error {
		hslStr := fmt.Sprintf("hsl(%d, %.0f%%, %.0f%%)", hue, saturation*100, lightness*100)
		_, err := w.Write([]byte(hslStr))
		return err
	})
}

func HSLA(hue uint, saturation, lightness, alpha float64) ColorValueFunc {
	if hue > 360 {
		hue = 360
	}
	if saturation < 0.0 {
		saturation = 0.0
	}
	if saturation > 1.0 {
		saturation = 1.0
	}
	if lightness < 0.0 {
		lightness = 0.0
	}
	if lightness > 1.0 {
		lightness = 1.0
	}
	if alpha < 0.0 {
		alpha = 0.0
	}
	if alpha > 1.0 {
		alpha = 1.0
	}

	return ColorValueFunc(func(w io.Writer) error {
		hslaStr := fmt.Sprintf("hsla(%d, %.0f%%, %.0f%%, %.1f)", hue, saturation*100, lightness*100, alpha)
		_, err := w.Write([]byte(hslaStr))
		return err
	})
}
