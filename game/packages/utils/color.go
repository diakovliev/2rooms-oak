package utils

import "image/color"

// InverseColor returns the inverse color of the given color.RGBA.
//
// It takes a color.RGBA as a parameter and returns a color.RGBA.
func InverseColor(color color.RGBA) (ret color.RGBA) {
	ret = color
	ret.R = ret.A - ret.R
	ret.G = ret.A - ret.G
	ret.B = ret.A - ret.B
	return
}
