package main

import (
	"image/color"
)

type Light interface {
	// Reset color to white with max brightness
	Reset()
	// Turn light on
	On()
	// Turn light off
	Off()
	// Set color with max brightness
	SetColor(color color.Color)
	// Set brightness manually
	SetBrightness(brightness int)
	// Set hue manually
	SetHue(hue int)
	// Set saturation manually
	SetSaturation(sat int)
}

type light struct {
	on bool
	reachable bool
	saturation int	
	brightness int
	color color.Color
}
