package donut

import (
	"image/color"
	"math"
	"math/rand/v2"
)

// LUT is a lookup table of base colors for the donut. It is intentionally
// editable.
var LUT = [12]color.RGBA{
	{0xFF, 0x00, 0x00, 0xFF}, // Red
	{0xFF, 0x7F, 0x00, 0xFF}, // Orange
	{0xFF, 0xFF, 0x00, 0xFF}, // Yellow
	{0x00, 0xFF, 0x00, 0xFF}, // Green
	{0x00, 0xFF, 0xFF, 0xFF}, // Cyan
	{0x00, 0x00, 0xFF, 0xFF}, // Blue
	{0x4B, 0x00, 0x82, 0xFF}, // Indigo
	{0x94, 0x00, 0xD3, 0xFF}, // Violet
	{0xFF, 0x14, 0x93, 0xFF}, // Pink
	{0xFF, 0xFF, 0xFF, 0xFF}, // White
	{0x8B, 0x45, 0x13, 0xFF}, // Brown
	{0x00, 0x00, 0x00, 0xFF}, // Black
}

func pickColor() byte {
	// Pick a random color from the LUT
	return byte(rand.IntN(len(LUT)))
}

func nextColor(current byte) byte {
	// Cycle to the next color in the LUT
	return byte((int(current) + 1) % len(LUT))
}

func blend(base, lumi byte, α float64) color.RGBA {
	b := LUT[base]
	l := ramp[lumi]

	return color.RGBA{
		R: uint8(float64(b.R)*(1-α) + float64(l.R)*α),
		G: uint8(float64(b.G)*(1-α) + float64(l.G)*α),
		B: uint8(float64(b.B)*(1-α) + float64(l.B)*α),
		A: 0xFF,
	}
}

var ramp [12]color.RGBA

func init() {
	for i := range ramp {
		t := float64(i) / 11.0
		// Gamma + bias baked in
		t = math.Pow(t, 0.7)
		t = 0.25 + 0.75*t // lift shadows
		v := int(255 * t)
		ramp[i] = color.RGBA{uint8(v), uint8(v), uint8(v), 0xFF}
	}
}
