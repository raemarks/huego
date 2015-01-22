package main

import (
	"image/color"
)

type Light interface {
	// Reset color to white with max brightness
	Reset() error
	// Turn light on
	On() error
	// Turn light off
	Off() error
	// Set color with max brightness
	SetColor(color color.Color) error
	// Set brightness manually
	SetBrightness(brightness int) error
	// Set hue manually
	SetHue(hue int) error
	// Set saturation manually
	SetSaturation(sat int) error
	// Set light's name
	SetName(name string) error
}
