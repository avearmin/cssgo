package gomour

import (
	"fmt"
	"io"
)

// ColorValue defines an interface for types representing CSS-compatible Color values.
// It ensures that any implementing type can be rendered as valid CSS.
type ColorValue interface {
	ValueNode
	colorValue()
}

// color represents a CSS color (e.g., "blue", "red", "#ffffff", "rgb(0, 40, 50)".
type Color string

// Named colors that follow the CSS standard.
const (
	AliceBlue            Color = "aliceblue"
	AntiqueWhite         Color = "antiquewhite"
	Aqua                 Color = "aqua"
	Aquamarine           Color = "aquamarine"
	Azure                Color = "azure"
	Beige                Color = "beige"
	Bisque               Color = "bisque"
	Black                Color = "black"
	BlanchedAlmond       Color = "blanchedalmond"
	Blue                 Color = "blue"
	BlueViolet           Color = "blueviolet"
	Brown                Color = "brown"
	BurlyWood            Color = "burlywood"
	CadetBlue            Color = "cadetblue"
	Chartreuse           Color = "chartreuse"
	Chocolate            Color = "chocolate"
	Coral                Color = "coral"
	CornflowerBlue       Color = "cornflowerblue"
	Cornsilk             Color = "cornsilk"
	Crimson              Color = "crimson"
	Cyan                 Color = "cyan"
	DarkBlue             Color = "darkblue"
	DarkCyan             Color = "darkcyan"
	DarkGoldenRod        Color = "darkgoldenrod"
	DarkGray             Color = "darkgray"
	DarkGreen            Color = "darkgreen"
	DarkKhaki            Color = "darkkhaki"
	DarkMagenta          Color = "darkmagenta"
	DarkOliveGreen       Color = "darkolivegreen"
	DarkOrange           Color = "darkorange"
	DarkOrchid           Color = "darkorchid"
	DarkRed              Color = "darkred"
	DarkSalmon           Color = "darksalmon"
	DarkSeaGreen         Color = "darkseagreen"
	DarkSlateBlue        Color = "darkslateblue"
	DarkSlateGray        Color = "darkslategray"
	DarkTurquoise        Color = "darkturquoise"
	DarkViolet           Color = "darkviolet"
	DeepPink             Color = "deeppink"
	DeepSkyBlue          Color = "deepskyblue"
	DimGray              Color = "dimgray"
	DodgerBlue           Color = "dodgerblue"
	FireBrick            Color = "firebrick"
	FloralWhite          Color = "floralwhite"
	ForestGreen          Color = "forestgreen"
	Fuchsia              Color = "fuchsia"
	Gainsboro            Color = "gainsboro"
	GhostWhite           Color = "ghostwhite"
	Gold                 Color = "gold"
	GoldenRod            Color = "goldenrod"
	Gray                 Color = "gray"
	Green                Color = "green"
	GreenYellow          Color = "greenyellow"
	HoneyDew             Color = "honeydew"
	HotPink              Color = "hotpink"
	IndianRed            Color = "indianred"
	Indigo               Color = "indigo"
	Ivory                Color = "ivory"
	Khaki                Color = "khaki"
	Lavender             Color = "lavender"
	LavenderBlush        Color = "lavenderblush"
	LawnGreen            Color = "lawngreen"
	LemonChiffon         Color = "lemonchiffon"
	LightBlue            Color = "lightblue"
	LightCoral           Color = "lightcoral"
	LightCyan            Color = "lightcyan"
	LightGoldenRodYellow Color = "lightgoldenrodyellow"
	LightGray            Color = "lightgray"
	LightGreen           Color = "lightgreen"
	LightPink            Color = "lightpink"
	LightSalmon          Color = "lightsalmon"
	LightSeaGreen        Color = "lightseagreen"
	LightSkyBlue         Color = "lightskyblue"
	LightSlateGray       Color = "lightslategray"
	LightSteelBlue       Color = "lightsteelblue"
	LightYellow          Color = "lightyellow"
	Lime                 Color = "lime"
	LimeGreen            Color = "limegreen"
	Linen                Color = "linen"
	Magenta              Color = "magenta"
	Maroon               Color = "maroon"
	MediumAquaMarine     Color = "mediumaquamarine"
	MediumBlue           Color = "mediumblue"
	MediumOrchid         Color = "mediumorchid"
	MediumPurple         Color = "mediumpurple"
	MediumSeaGreen       Color = "mediumseagreen"
	MediumSlateBlue      Color = "mediumslateblue"
	MediumSpringGreen    Color = "mediumspringgreen"
	MediumTurquoise      Color = "mediumturquoise"
	MediumVioletRed      Color = "mediumvioletred"
	MidnightBlue         Color = "midnightblue"
	MintCream            Color = "mintcream"
	MistyRose            Color = "mistyrose"
	Moccasin             Color = "moccasin"
	NavajoWhite          Color = "navajowhite"
	Navy                 Color = "navy"
	OldLace              Color = "oldlace"
	Olive                Color = "olive"
	OliveDrab            Color = "olivedrab"
	Orange               Color = "orange"
	OrangeRed            Color = "orangered"
	Orchid               Color = "orchid"
	PaleGoldenRod        Color = "palegoldenrod"
	PaleGreen            Color = "palegreen"
	PaleTurquoise        Color = "paleturquoise"
	PaleVioletRed        Color = "palevioletred"
	PapayaWhip           Color = "papayawhip"
	PeachPuff            Color = "peachpuff"
	Peru                 Color = "peru"
	Pink                 Color = "pink"
	Plum                 Color = "plum"
	PowderBlue           Color = "powderblue"
	Purple               Color = "purple"
	RebeccaPurple        Color = "rebeccapurple"
	Red                  Color = "red"
	RosyBrown            Color = "rosybrown"
	RoyalBlue            Color = "royalblue"
	SaddleBrown          Color = "saddlebrown"
	Salmon               Color = "salmon"
	SandyBrown           Color = "sandybrown"
	SeaGreen             Color = "seagreen"
	SeaShell             Color = "seashell"
	Sienna               Color = "sienna"
	Silver               Color = "silver"
	SkyBlue              Color = "skyblue"
	SlateBlue            Color = "slateblue"
	SlateGray            Color = "slategray"
	Snow                 Color = "snow"
	SpringGreen          Color = "springgreen"
	SteelBlue            Color = "steelblue"
	Tan                  Color = "tan"
	Teal                 Color = "teal"
	Thistle              Color = "thistle"
	Tomato               Color = "tomato"
	Turquoise            Color = "turquoise"
	Violet               Color = "violet"
	Wheat                Color = "wheat"
	White                Color = "white"
	WhiteSmoke           Color = "whitesmoke"
	Yellow               Color = "yellow"
	YellowGreen          Color = "yellowgreen"
)

func (c Color) RenderCSS(w io.Writer) error {
	_, err := w.Write([]byte(c))
	return err
}

func (c Color) valueNode()  {}
func (c Color) colorValue() {}

// Hex generates a hexadecimal Color value (e.g., #ffffff).
// Parameter:
// - `hexdec` (int): An integer representing the hexadecimal Color value (e.g., 0xffffff).
func Hex(hexdec int) Color {
	return Color(fmt.Sprintf("#%06x", hexdec))
}

// RGB generates an RGB Color value (e.g., rgb(255, 255, 255)).
// Parameters:
// - `r` (int): Red component (0–255).
// - `g` (int): Green component (0–255).
// - `b` (int): Blue component (0–255).
func RGB(r, g, b int) Color {
	return Color(fmt.Sprintf("rgb(%d, %d, %d)", r, g, b))
}

// RGBA generates an RGBA Color value (e.g., rgba(255, 255, 255, 0.5)).
// Parameters:
// - `r` (int): Red component (0–255).
// - `g` (int): Green component (0–255).
// - `b` (int): Blue component (0–255).
// - `alpha` (float64): Opacity (0.0–1.0).
func RGBA(r, g, b int, alpha float64) Color {
	return Color(fmt.Sprintf("rgba(%d, %d, %d, %g)", r, g, b, alpha))
}

// HSL generates an HSL Color value in the format `hsl(hue, saturation%, lightness%)`.
// Parameters:
// - `hue` (int): Represents the Color angle in degrees. Valid CSS values include:
//   - Any integer value (e.g., -30, 0, 360, 450).
//   - Negative values wrap counter-clockwise (e.g., -30 is equivalent to 330).
//   - Values greater than 360 wrap clockwise (e.g., 450 is equivalent to 90).
//
// - `saturation` (float64): Represents the saturation as a percentage (e.g., `100.0` for `100%`).
//   - Values must be non-negative.
//
// - `lightness` (float64): Represents the lightness as a percentage (e.g., `50.5` for `50.5%`).
//   - Values must be non-negative.
func HSL(hue int, saturation, lightness float64) Color {
	return Color(fmt.Sprintf("hsl(%d, %g%%, %g%%)", hue, saturation, lightness))
}

// HSLA generates an HSLA Color value in the format `hsla(hue, saturation%, lightness%, alpha).`.
//
// Parameters:
// - `hue` (int): Represents the Color angle in degrees. Valid CSS values include:
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
// - `alpha` (float64): Represents the opacity of the Color as a normalized value between `0.0` (fully transparent) and `1.0` (fully opaque).
func HSLA(hue int, saturation, lightness, alpha float64) Color {
	return Color(fmt.Sprintf("hsla(%d, %g%%, %g%%, %g)", hue, saturation, lightness, alpha))
}
