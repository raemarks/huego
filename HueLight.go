package main

import (
	"image/color"
	"errors"
)

type state struct {
	On bool
	Bri int
	Hue int
	Sat int
	Xy []float64
	Ct int
	Alert string
	Effect string
	Colormode string
	Reachable bool
}

type HueLight struct {
	Light

	// Fields not contained in JSON
	color color.Color `json:"-"`
	Bridge *HueBridge `json:"-"`
	Id string `json:"-"`

	Name string
	State state
	Type string
	Modelid string
	Uniqueid string
	Swversion string
}

func (hlight *HueLight) Reset() error {
	return nil
}

//TODO: add a timer option? Add a instant-on option that modifies the fade time?
//Make On and OnWithDelay
func (hlight *HueLight) On() error {
	data := make(JSON)
	data["on"] = true

	res, err := hlight.Bridge.UpdateLight(hlight.Id, data)
	if err != nil {
		return err
	}

	// Check if server threw an error
	if m, ok := res[0]["error"]; ok {
		// Go complains if I use JSON here, no aliasing
		mobj, ok := m.(map[string] interface{})
		if !ok {
			// Unexpected response, not a matching API
			panic("Unexpected response from server")
		}

		return errors.New("Error when turning light on: " +
			mobj["description"].(string))
	}

	return nil
}

func (hlight *HueLight) Off() error {
	data := make(JSON)
	data["on"] = false

	res, err := hlight.Bridge.UpdateLight(hlight.Id, data)
	if err != nil {
		return err
	}

	// Check if server threw an error
	if m, ok := res[0]["error"]; ok {
		// Go complains if I use JSON here, no aliasing
		mobj, ok := m.(map[string] interface{})
		if !ok {
			// Unexpected response, not a matching API
			panic("Unexpected response from server")
		}

		return errors.New("Error when turning light off: " +
			mobj["description"].(string))
	}

	return nil
}

func (hlight *HueLight) SetColor(color color.Color) error {
	return nil
}

func (hlight *HueLight) SetBrightness(brightness int) error {
	return nil
}

func (hlight *HueLight) SetHue(hue int) error {
	return nil
}

func (hlight *HueLight) SetSaturation(sat int) error {
	return nil
}

func (hlight *HueLight) SetName(name string) error {
	return nil
}
