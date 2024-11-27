package gomour

import (
	"fmt"
	"io"
	"strings"
)

// ColorValue defines an interface for types representing CSS-compatible color values.
// It ensures that any implementing type can be rendered as valid CSS.
type ColorValue interface {
	ValueNode
	colorValue()
}

// namedColor represents a predefined CSS named color (e.g., "blue", "red").
// These named colors follow the CSS standard.
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

// ColorValueFunc is a function type used to define custom color values
// (e.g., RGB, RGBA, Hex). It implements the ColorValue interface.
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

// Hex generates a hexadecimal color value (e.g., #ffffff).
// Parameter:
// - `hexdec` (int): An integer representing the hexadecimal color value (e.g., 0xffffff).
func Hex(hexdec int) ColorValueFunc {
	return ColorValueFunc(func(w io.Writer) error {
		hexStr := fmt.Sprintf("#%06x", hexdec)
		_, err := w.Write([]byte(hexStr))
		return err
	})
}

// RGB generates an RGB color value (e.g., rgb(255, 255, 255)).
// Parameters:
// - `r` (int): Red component (0–255).
// - `g` (int): Green component (0–255).
// - `b` (int): Blue component (0–255).
func RGB(r, g, b int) ColorValueFunc {
	return ColorValueFunc(func(w io.Writer) error {
		rgbStr := fmt.Sprintf("rgb(%d, %d, %d)", r, g, b)
		_, err := w.Write([]byte(rgbStr))
		return err
	})
}

// RGBA generates an RGBA color value (e.g., rgba(255, 255, 255, 0.5)).
// Parameters:
// - `r` (int): Red component (0–255).
// - `g` (int): Green component (0–255).
// - `b` (int): Blue component (0–255).
// - `alpha` (float64): Opacity (0.0–1.0).
func RGBA(r, g, b int, alpha float64) ColorValueFunc {
	return ColorValueFunc(func(w io.Writer) error {
		rgbaStr := fmt.Sprintf("rgba(%d, %d, %d, %g)", r, g, b, alpha)
		_, err := w.Write([]byte(rgbaStr))
		return err
	})
}

// HSL generates an HSL color value in the format `hsl(hue, saturation%, lightness%)`.
// Parameters:
// - `hue` (int): Represents the color angle in degrees. Valid CSS values include:
//   - Any integer value (e.g., -30, 0, 360, 450).
//   - Negative values wrap counter-clockwise (e.g., -30 is equivalent to 330).
//   - Values greater than 360 wrap clockwise (e.g., 450 is equivalent to 90).
//
// - `saturation` (float64): Represents the saturation as a percentage (e.g., `100.0` for `100%`).
//   - Values must be non-negative.
//
// - `lightness` (float64): Represents the lightness as a percentage (e.g., `50.5` for `50.5%`).
//   - Values must be non-negative.
//
// Returns:
// - A `ColorValueFunc` that outputs the HSL value as valid CSS.
func HSL(hue int, saturation, lightness float64) ColorValueFunc {
	return ColorValueFunc(func(w io.Writer) error {
		hslStr := fmt.Sprintf("hsl(%d, %g%%, %g%%)", hue, saturation, lightness)
		_, err := w.Write([]byte(hslStr))
		return err
	})
}

// HSLA generates an HSLA color value in the format `hsla(hue, saturation%, lightness%, alpha)`.
// Parameters:
// - `hue` (int): Represents the color angle in degrees. Valid CSS values include:
//   - Any integer value (e.g., -30, 0, 360, 450).
//   - Negative values wrap counter-clockwise (e.g., -30 is equivalent to 330).
//   - Values greater than 360 wrap clockwise (e.g., 450 is equivalent to 90).
//
// - `saturation` (float64): Represents the saturation as a percentage (e.g., `100.0` for `100%`).
//   - Values must be non-negative.
//
// - `lightness` (float64): Represents the lightness as a percentage (e.g., `50.5` for `50.5%`).
//   - Values must be non-negative.
//
// - `alpha` (float64): Represents the opacity of the color as a normalized value between `0.0` (fully transparent) and `1.0` (fully opaque).
// Returns:
// - A `ColorValueFunc` that outputs the HSLA value as valid CSS.
func HSLA(hue int, saturation, lightness, alpha float64) ColorValueFunc {
	return ColorValueFunc(func(w io.Writer) error {
		hslaStr := fmt.Sprintf("hsla(%d, %g%%, %g%%, %g)", hue, saturation, lightness, alpha)
		_, err := w.Write([]byte(hslaStr))
		return err
	})
}
