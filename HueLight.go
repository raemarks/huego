package main

type state struct {
	On bool
	Bri int
	Hue int
	Sat int
	Xy []int
	Ct int
	Alert string
	Effect string
	Colormode string
	Reachable bool
}

type HueLight struct {
	light `json:"-"`

	State state
	Type string
	Name string
	Modelid string
	Uniqueid string
	Swversion string
}


/*
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
*/
