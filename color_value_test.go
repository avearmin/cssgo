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
		test{"RGBA: rgba(0, 0, 0, 0)", RGBA(0, 0, 0, 0), "rgba(0, 0, 0, 0)"},
		test{"RGBA: rgba(255, 255, 255, 1)", RGBA(255, 255, 255, 1), "rgba(255, 255, 255, 1)"},
		test{"RGBA: rgba(123, 45, 67, 0.5)", RGBA(123, 45, 67, 0.5), "rgba(123, 45, 67, 0.5)"},
	)
}

func TestHSL(t *testing.T) {
	RunTests(t,
		test{"HSL: hsl(0, 0%, 0%)", HSL(0, 0, 0), "hsl(0, 0%, 0%)"},
		test{"HSL: hsl(360, 100%, 50%)", HSL(360, 100, 50), "hsl(360, 100%, 50%)"},
		test{"HSL: hsl(180, 50%, 25%)", HSL(180, 50, 25), "hsl(180, 50%, 25%)"},
		test{"HSL: hsl(360, 100.55%, 25.643%)", HSL(360, 100.55, 25.643), "hsl(360, 100.55%, 25.643%)"},
	)
}

func TestHSLA(t *testing.T) {
	RunTests(t,
		test{"HSLA: hsla(0, 0%, 0%, 0)", HSLA(0, 0, 0, 0), "hsla(0, 0%, 0%, 0)"},
		test{"HSLA: hsla(360, 100%, 50%, 1)", HSLA(360, 100, 50, 1), "hsla(360, 100%, 50%, 1)"},
		test{"HSLA: hsla(180, 50%, 25%, 0.5)", HSLA(180, 50, 25, 0.5), "hsla(180, 50%, 25%, 0.5)"},
		test{"HSLA: hsla(360, 100.55%, 25.643%, 0.3322)", HSLA(360, 100.55, 25.643, 0.3322), "hsla(360, 100.55%, 25.643%, 0.3322)"},
	)
}

func TestNamedColors(t *testing.T) {
	RunTests(t,
		test{
			"aliceblue constant",
			AliceBlue,
			"aliceblue",
		},
		test{
			"antiquewhite constant",
			AntiqueWhite,
			"antiquewhite",
		},
		test{
			"aqua constant",
			Aqua,
			"aqua",
		},
		test{
			"aquamarine constant",
			Aquamarine,
			"aquamarine",
		},
		test{
			"azure constant",
			Azure,
			"azure",
		},
		test{
			"beige constant",
			Beige,
			"beige",
		},
		test{
			"bisque constant",
			Bisque,
			"bisque",
		},
		test{
			"black constant",
			Black,
			"black",
		},
		test{
			"blanchedalmond constant",
			BlanchedAlmond,
			"blanchedalmond",
		},
		test{
			"blue constant",
			Blue,
			"blue",
		},
		test{
			"blueviolet constant",
			BlueViolet,
			"blueviolet",
		},
		test{
			"brown constant",
			Brown,
			"brown",
		},
		test{
			"burlywood constant",
			BurlyWood,
			"burlywood",
		},
		test{
			"cadetblue constant",
			CadetBlue,
			"cadetblue",
		},
		test{
			"chartreuse constant",
			Chartreuse,
			"chartreuse",
		},
		test{
			"chocolate constant",
			Chocolate,
			"chocolate",
		},
		test{
			"coral constant",
			Coral,
			"coral",
		},
		test{
			"cornflowerblue constant",
			CornflowerBlue,
			"cornflowerblue",
		},
		test{
			"cornsilk constant",
			Cornsilk,
			"cornsilk",
		},
		test{
			"crimson constant",
			Crimson,
			"crimson",
		},
		test{
			"cyan constant",
			Cyan,
			"cyan",
		},
		test{
			"darkblue constant",
			DarkBlue,
			"darkblue",
		},
		test{
			"darkcyan constant",
			DarkCyan,
			"darkcyan",
		},
		test{
			"darkgoldenrod constant",
			DarkGoldenRod,
			"darkgoldenrod",
		},
		test{
			"darkgray constant",
			DarkGray,
			"darkgray",
		},
		test{
			"darkgreen constant",
			DarkGreen,
			"darkgreen",
		},
		test{
			"darkkhaki constant",
			DarkKhaki,
			"darkkhaki",
		},
		test{
			"darkmagenta constant",
			DarkMagenta,
			"darkmagenta",
		},
		test{
			"darkolivegreen constant",
			DarkOliveGreen,
			"darkolivegreen",
		},
		test{
			"darkorange constant",
			DarkOrange,
			"darkorange",
		},
		test{
			"darkorchid constant",
			DarkOrchid,
			"darkorchid",
		},
		test{
			"darkred constant",
			DarkRed,
			"darkred",
		},
		test{
			"darksalmon constant",
			DarkSalmon,
			"darksalmon",
		},
		test{
			"darkseagreen constant",
			DarkSeaGreen,
			"darkseagreen",
		},
		test{
			"darkslateblue constant",
			DarkSlateBlue,
			"darkslateblue",
		},
		test{
			"darkslategray constant",
			DarkSlateGray,
			"darkslategray",
		},
		test{
			"darkturquoise constant",
			DarkTurquoise,
			"darkturquoise",
		},
		test{
			"darkviolet constant",
			DarkViolet,
			"darkviolet",
		},
		test{
			"deeppink constant",
			DeepPink,
			"deeppink",
		},
		test{
			"deepskyblue constant",
			DeepSkyBlue,
			"deepskyblue",
		},
		test{
			"dimgray constant",
			DimGray,
			"dimgray",
		},
		test{
			"dodgerblue constant",
			DodgerBlue,
			"dodgerblue",
		},
		test{
			"firebrick constant",
			FireBrick,
			"firebrick",
		},
		test{
			"floralwhite constant",
			FloralWhite,
			"floralwhite",
		},
		test{
			"forestgreen constant",
			ForestGreen,
			"forestgreen",
		},
		test{
			"fuchsia constant",
			Fuchsia,
			"fuchsia",
		},
		test{
			"gainsboro constant",
			Gainsboro,
			"gainsboro",
		},
		test{
			"ghostwhite constant",
			GhostWhite,
			"ghostwhite",
		},
		test{
			"gold constant",
			Gold,
			"gold",
		},
		test{
			"goldenrod constant",
			GoldenRod,
			"goldenrod",
		},
		test{
			"gray constant",
			Gray,
			"gray",
		},
		test{
			"green constant",
			Green,
			"green",
		},
		test{
			"greenyellow constant",
			GreenYellow,
			"greenyellow",
		},
		test{
			"honeydew constant",
			HoneyDew,
			"honeydew",
		},
		test{
			"hotpink constant",
			HotPink,
			"hotpink",
		},
		test{
			"indianred constant",
			IndianRed,
			"indianred",
		},
		test{
			"indigo constant",
			Indigo,
			"indigo",
		},
		test{
			"ivory constant",
			Ivory,
			"ivory",
		},
		test{
			"khaki constant",
			Khaki,
			"khaki",
		},
		test{
			"lavender constant",
			Lavender,
			"lavender",
		},
		test{
			"lavenderblush constant",
			LavenderBlush,
			"lavenderblush",
		},
		test{
			"lawngreen constant",
			LawnGreen,
			"lawngreen",
		},
		test{
			"lemonchiffon constant",
			LemonChiffon,
			"lemonchiffon",
		},
		test{
			"lightblue constant",
			LightBlue,
			"lightblue",
		},
		test{
			"lightcoral constant",
			LightCoral,
			"lightcoral",
		},
		test{
			"lightcyan constant",
			LightCyan,
			"lightcyan",
		},
		test{
			"lightgoldenrodyellow constant",
			LightGoldenRodYellow,
			"lightgoldenrodyellow",
		},
		test{
			"lightgray constant",
			LightGray,
			"lightgray",
		},
		test{
			"lightgreen constant",
			LightGreen,
			"lightgreen",
		},
		test{
			"lightpink constant",
			LightPink,
			"lightpink",
		},
		test{
			"lightsalmon constant",
			LightSalmon,
			"lightsalmon",
		},
		test{
			"lightseagreen constant",
			LightSeaGreen,
			"lightseagreen",
		},
		test{
			"lightskyblue constant",
			LightSkyBlue,
			"lightskyblue",
		},
		test{
			"lightslategray constant",
			LightSlateGray,
			"lightslategray",
		},
		test{
			"lightsteelblue constant",
			LightSteelBlue,
			"lightsteelblue",
		},
		test{
			"lightyellow constant",
			LightYellow,
			"lightyellow",
		},
		test{
			"lime constant",
			Lime,
			"lime",
		},
		test{
			"limegreen constant",
			LimeGreen,
			"limegreen",
		},
		test{
			"linen constant",
			Linen,
			"linen",
		},
		test{
			"magenta constant",
			Magenta,
			"magenta",
		},
		test{
			"maroon constant",
			Maroon,
			"maroon",
		},
		test{
			"mediumaquamarine constant",
			MediumAquaMarine,
			"mediumaquamarine",
		},
		test{
			"mediumblue constant",
			MediumBlue,
			"mediumblue",
		},
		test{
			"mediumorchid constant",
			MediumOrchid,
			"mediumorchid",
		},
		test{
			"mediumpurple constant",
			MediumPurple,
			"mediumpurple",
		},
		test{
			"mediumseagreen constant",
			MediumSeaGreen,
			"mediumseagreen",
		},
		test{
			"mediumslateblue constant",
			MediumSlateBlue,
			"mediumslateblue",
		},
		test{
			"mediumspringgreen constant",
			MediumSpringGreen,
			"mediumspringgreen",
		},
		test{
			"mediumturquoise constant",
			MediumTurquoise,
			"mediumturquoise",
		},
		test{
			"mediumvioletred constant",
			MediumVioletRed,
			"mediumvioletred",
		},
		test{
			"midnightblue constant",
			MidnightBlue,
			"midnightblue",
		},
		test{
			"mintcream constant",
			MintCream,
			"mintcream",
		},
		test{
			"mistyrose constant",
			MistyRose,
			"mistyrose",
		},
		test{
			"moccasin constant",
			Moccasin,
			"moccasin",
		},
		test{
			"navajowhite constant",
			NavajoWhite,
			"navajowhite",
		},
		test{
			"navy constant",
			Navy,
			"navy",
		},
		test{
			"oldlace constant",
			OldLace,
			"oldlace",
		},
		test{
			"olive constant",
			Olive,
			"olive",
		},
		test{
			"olivedrab constant",
			OliveDrab,
			"olivedrab",
		},
		test{
			"orange constant",
			Orange,
			"orange",
		},
		test{
			"orangered constant",
			OrangeRed,
			"orangered",
		},
		test{
			"orchid constant",
			Orchid,
			"orchid",
		},
		test{
			"palegoldenrod constant",
			PaleGoldenRod,
			"palegoldenrod",
		},
		test{
			"palegreen constant",
			PaleGreen,
			"palegreen",
		},
		test{
			"paleturquoise constant",
			PaleTurquoise,
			"paleturquoise",
		},
		test{
			"palevioletred constant",
			PaleVioletRed,
			"palevioletred",
		},
		test{
			"papayawhip constant",
			PapayaWhip,
			"papayawhip",
		},
		test{
			"peachpuff constant",
			PeachPuff,
			"peachpuff",
		},
		test{
			"peru constant",
			Peru,
			"peru",
		},
		test{
			"pink constant",
			Pink,
			"pink",
		},
		test{
			"plum constant",
			Plum,
			"plum",
		},
		test{
			"powderblue constant",
			PowderBlue,
			"powderblue",
		},
		test{
			"purple constant",
			Purple,
			"purple",
		},
		test{
			"rebeccapurple constant",
			RebeccaPurple,
			"rebeccapurple",
		},
		test{
			"red constant",
			Red,
			"red",
		},
		test{
			"rosybrown constant",
			RosyBrown,
			"rosybrown",
		},
		test{
			"royalblue constant",
			RoyalBlue,
			"royalblue",
		},
		test{
			"saddlebrown constant",
			SaddleBrown,
			"saddlebrown",
		},
		test{
			"salmon constant",
			Salmon,
			"salmon",
		},
		test{
			"sandybrown constant",
			SandyBrown,
			"sandybrown",
		},
		test{
			"seagreen constant",
			SeaGreen,
			"seagreen",
		},
		test{
			"seashell constant",
			SeaShell,
			"seashell",
		},
		test{
			"sienna constant",
			Sienna,
			"sienna",
		},
		test{
			"silver constant",
			Silver,
			"silver",
		},
		test{
			"skyblue constant",
			SkyBlue,
			"skyblue",
		},
		test{
			"slateblue constant",
			SlateBlue,
			"slateblue",
		},
		test{
			"slategray constant",
			SlateGray,
			"slategray",
		},
		test{
			"snow constant",
			Snow,
			"snow",
		},
		test{
			"springgreen constant",
			SpringGreen,
			"springgreen",
		},
		test{
			"steelblue constant",
			SteelBlue,
			"steelblue",
		},
		test{
			"tan constant",
			Tan,
			"tan",
		},
		test{
			"teal constant",
			Teal,
			"teal",
		},
		test{
			"thistle constant",
			Thistle,
			"thistle",
		},
		test{
			"tomato constant",
			Tomato,
			"tomato",
		},
		test{
			"turquoise constant",
			Turquoise,
			"turquoise",
		},
		test{
			"violet constant",
			Violet,
			"violet",
		},
		test{
			"wheat constant",
			Wheat,
			"wheat",
		},
		test{
			"white constant",
			White,
			"white",
		},
		test{
			"whitesmoke constant",
			WhiteSmoke,
			"whitesmoke",
		},
		test{
			"yellow constant",
			Yellow,
			"yellow",
		},
		test{
			"yellowgreen constant",
			YellowGreen,
			"yellowgreen",
		},
	)
}
