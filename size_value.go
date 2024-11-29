package cssgo

import (
	"io"
	"strconv"
)

// SizeValue defines an interface for types representing CSS-compatible size values.
// This ensures that any type implementing this interface can be rendered as a valid CSS size.
type SizeValue interface {
	ValueNode
	sizeValue()
}

// Size represents a CSS size value (e.g., "10px", "50%", "1.5em").
// It is a concrete type that implements the SizeValue interface.
type Size string

func (s Size) RenderCSS(w io.Writer) error {
	_, err := w.Write([]byte(s))
	return err
}

func (s Size) valueNode() {}
func (s Size) sizeValue() {}

// size is a helper function that generates a Size value from a float and a unit.
// It converts the float to a string and appends the unit to form a CSS-compatible size string.
//
// Parameters:
// - value (float64): The numeric size value (e.g., 10.5, 1.2).
// - unit (string): The CSS unit (e.g., "px", "em").
//
// Returns:
// - Size: A new Size instance representing the value and unit.
func size(value float64, unit string) Size {
	floatStr := strconv.FormatFloat(value, 'f', -1, 64)
	return Size(floatStr + unit)
}

// CM generates a size in centimeters (cm).
// Example: CM(1.5) -> "1.5cm"
func CM(value float64) Size {
	return size(value, "cm")
}

// MM generates a size in millimeters (mm).
// Example: MM(10.0) -> "10mm"
func MM(value float64) Size {
	return size(value, "mm")
}

// IN generates a size in inches (in).
// Example: IN(2.0) -> "2in"
func IN(value float64) Size {
	return size(value, "in")
}

// PX generates a size in pixels (px).
// Example: PX(16.0) -> "16px"
func PX(value float64) Size {
	return size(value, "px")
}

// PT generates a size in points (pt).
// Example: PT(12.0) -> "12pt"
func PT(value float64) Size {
	return size(value, "pt")
}

// PC generates a size in picas (pc).
// Example: PC(1.0) -> "1pc"
func PC(value float64) Size {
	return size(value, "pc")
}

// EM generates a size in relative "em" units (relative to the font size of the element).
// Example: EM(1.5) -> "1.5em"
func EM(value float64) Size {
	return size(value, "em")
}

// REM generates a size in root "rem" units (relative to the root font size).
// Example: REM(1.2) -> "1.2rem"
func REM(value float64) Size {
	return size(value, "rem")
}

// VW generates a size in viewport width (vw).
// Example: VW(50.0) -> "50vw"
func VW(value float64) Size {
	return size(value, "vw")
}

// VH generates a size in viewport height (vh).
// Example: VH(100.0) -> "100vh"
func VH(value float64) Size {
	return size(value, "vh")
}

// PCT generates a size in percentages (%).
// Example: PCT(75.5) -> "75.5%"
func PCT(value float64) Size {
	return size(value, "%")
}

// VMIN generates a size in viewport minimum (vmin).
// Example: VMIN(10.0) -> "10vmin"
func VMIN(value float64) Size {
	return size(value, "vmin")
}

// VMAX generates a size in viewport maximum (vmax).
// Example: VMAX(20.0) -> "20vmax"
func VMAX(value float64) Size {
	return size(value, "vmax")
}
