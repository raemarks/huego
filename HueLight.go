package huego

import (
	"errors"
	"fmt"
)

type state struct {
	On             bool
	Bri            int
	Hue            int
	Sat            int
	Xy             []float64
	Ct             int
	Alert          string
	Effect         string
	Colormode      string
	Reachable      bool
	Transitiontime int16 `json:"-"`
}

type HueLight struct {
	// Fields not contained in JSON
	hscolorf HSColorFloat `json:"-"`
	hscolori HSColorInt   `json:"-"`
	xycolor  XYColor      `json:"-"`
	Bridge   *HueBridge   `json:"-"`
	Id       string       `json:"-"`

	Name      string
	State     state
	Type      string
	Modelid   string
	Uniqueid  string
	Swversion string
}

func (hlight *HueLight) Reset() error {
	data := make(JSON)
	// These are the defaults for Hue bulbs after a power cycle
	data["on"] = true
	data["bri"] = 254
	data["hue"] = 14922
	data["sat"] = 144
	data["effect"] = "none"
	data["alert"] = "none"
	data["transitiontime"] = 4

	res, err := hlight.Bridge.UpdateLight("/lights/"+hlight.Id+"/state", data)
	if err != nil {
		return err
	}

	// Check if server threw an error
	if m, ok := res[0]["error"]; ok {
		// Go complains if I use JSON here, no aliasing
		mobj, ok := m.(map[string]interface{})
		if !ok {
			// Unexpected response, not a matching API
			panic("Unexpected response from server")
		}

		return errors.New("Error when resetting light: " +
			mobj["description"].(string))
	}

	return nil
}

//TODO: add a timer option? Add a instant-on option that modifies the fade time?
//Make On and OnWithDelay
func (hlight *HueLight) On() error {
	data := make(JSON)
	data["on"] = true

	res, err := hlight.Bridge.UpdateLight("/lights/"+hlight.Id+"/state", data)
	if err != nil {
		return err
	}

	// Check if server threw an error
	if m, ok := res[0]["error"]; ok {
		// Go complains if I use JSON here, no aliasing
		mobj, ok := m.(map[string]interface{})
		if !ok {
			// Unexpected response, not a matching API
			panic("Unexpected response from server")
		}

		return errors.New("Error when turning light on: " +
			mobj["description"].(string))
	}

	hlight.State.On = true
	return nil
}

func (hlight *HueLight) Off() error {
	data := make(JSON)
	data["on"] = false

	res, err := hlight.Bridge.UpdateLight("/lights/"+hlight.Id+"/state", data)
	if err != nil {
		return err
	}

	// Check if server threw an error
	if m, ok := res[0]["error"]; ok {
		// Go complains if I use JSON here, no aliasing
		mobj, ok := m.(map[string]interface{})
		if !ok {
			// Unexpected response, not a matching API
			panic("Unexpected response from server")
		}

		return errors.New("Error when turning light off: " +
			mobj["description"].(string))
	}

	hlight.State.On = false
	return nil
}

func (hlight *HueLight) SetTransitionTime(time int16) error {
	hlight.State.Transitiontime = time

	/*
		data := make(JSON)
		data["transitiontime"] = time
		fmt.Println("Changing transition time to %d", time)

		res, err := hlight.Bridge.UpdateLight("/lights/"+hlight.Id+"/state", data)
		if err != nil {
			return err
		}

		// Check if server threw an error
		if m, ok := res[0]["error"]; ok {
			// Go complains if I use JSON here, no aliasing
			mobj, ok := m.(map[string]interface{})
			if !ok {
				// Unexpected response, not a matching API
				panic("Unexpected response from server")
			}

			return errors.New("Error when setting color: " +
				mobj["description"].(string))
		}

	*/
	return nil
}

func (hlight *HueLight) SetColorXY(color XYColor) error {
	hlight.xycolor = color

	data := make(JSON)
	var xy []float64
	xy = append(xy, color.X)
	xy = append(xy, color.Y)
	data["xy"] = xy
	data["transitiontime"] = hlight.State.Transitiontime
	fmt.Println("Changing color to x=%lf y=%lf", color.X, color.Y)

	res, err := hlight.Bridge.UpdateLight("/lights/"+hlight.Id+"/state", data)
	if err != nil {
		return err
	}

	// Check if server threw an error
	if m, ok := res[0]["error"]; ok {
		// Go complains if I use JSON here, no aliasing
		mobj, ok := m.(map[string]interface{})
		if !ok {
			// Unexpected response, not a matching API
			panic("Unexpected response from server")
		}

		return errors.New("Error when setting color: " +
			mobj["description"].(string))
	}

	return nil
}

func (hlight *HueLight) SetColorFloatHS(color HSColorFloat) error {
	// Hue only accepts hsv, so need to convert from rgb.
	hlight.hscolorf = color

	data := make(JSON)
	data["hue"] = int(color.H * 65535.0)
	data["sat"] = int(color.S * 255.0)
	data["transitiontime"] = hlight.State.Transitiontime
	fmt.Println("Changing color to %d %d", int(color.H*65535.0), int(color.S*255.0))

	res, err := hlight.Bridge.UpdateLight("/lights/"+hlight.Id+"/state", data)
	if err != nil {
		return err
	}

	// Check if server threw an error
	if m, ok := res[0]["error"]; ok {
		// Go complains if I use JSON here, no aliasing
		mobj, ok := m.(map[string]interface{})
		if !ok {
			// Unexpected response, not a matching API
			panic("Unexpected response from server")
		}

		return errors.New("Error when setting color: " +
			mobj["description"].(string))
	}

	return nil
}

func (hlight *HueLight) SetColorIntHS(color HSColorInt) error {
	// Hue only accepts hsv, so need to convert from rgb.
	hlight.hscolori = color

	data := make(JSON)
	data["hue"] = color.H
	data["sat"] = color.S
	data["transitiontime"] = hlight.State.Transitiontime
	fmt.Println("Changing color to %d %d", color.H, color.S)

	res, err := hlight.Bridge.UpdateLight("/lights/"+hlight.Id+"/state", data)
	if err != nil {
		return err
	}

	// Check if server threw an error
	if m, ok := res[0]["error"]; ok {
		// Go complains if I use JSON here, no aliasing
		mobj, ok := m.(map[string]interface{})
		if !ok {
			// Unexpected response, not a matching API
			panic("Unexpected response from server")
		}

		return errors.New("Error when setting color: " +
			mobj["description"].(string))
	}

	return nil
}

func (hlight *HueLight) SetBrightness(brightness int) error {
	if brightness < 0 || brightness > 254 {
		return errors.New("Brightness value not in range 0-254")
	}

	data := make(JSON)
	data["bri"] = brightness
	data["transitiontime"] = hlight.State.Transitiontime

	res, err := hlight.Bridge.UpdateLight("/lights/"+hlight.Id+"/state", data)
	if err != nil {
		return err
	}

	// Check if server threw an error
	if m, ok := res[0]["error"]; ok {
		// Go complains if I use JSON here, no aliasing
		mobj, ok := m.(map[string]interface{})
		if !ok {
			// Unexpected response, not a matching API
			panic("Unexpected response from server")
		}

		return errors.New("Error when setting brightness: " +
			mobj["description"].(string))
	}

	return nil
}

func (hlight *HueLight) SetHue(hue int) error {
	if hue < 0 || hue > 65535 {
		return errors.New("Hue value not in range 0-65535")
	}

	data := make(JSON)
	data["hue"] = hue
	data["transitiontime"] = hlight.State.Transitiontime

	res, err := hlight.Bridge.UpdateLight("/lights/"+hlight.Id+"/state", data)
	if err != nil {
		return err
	}

	// Check if server threw an error
	if m, ok := res[0]["error"]; ok {
		// Go complains if I use JSON here, no aliasing
		mobj, ok := m.(map[string]interface{})
		if !ok {
			// Unexpected response, not a matching API
			panic("Unexpected response from server")
		}

		return errors.New("Error when setting hue: " +
			mobj["description"].(string))
	}

	return nil
}

func (hlight *HueLight) SetSaturation(sat int) error {
	if sat < 0 || sat > 255 {
		return errors.New("Sat value not in range 0-255")
	}

	data := make(JSON)
	data["sat"] = sat
	data["transitiontime"] = hlight.State.Transitiontime

	res, err := hlight.Bridge.UpdateLight("/lights/"+hlight.Id+"/state", data)
	if err != nil {
		return err
	}

	// Check if server threw an error
	if m, ok := res[0]["error"]; ok {
		// Go complains if I use JSON here, no aliasing
		mobj, ok := m.(map[string]interface{})
		if !ok {
			// Unexpected response, not a matching API
			panic("Unexpected response from server")
		}

		return errors.New("Error when setting saturation: " +
			mobj["description"].(string))
	}

	return nil
}

func (hlight *HueLight) SetName(name string) error {
	data := make(JSON)
	data["name"] = name

	res, err := hlight.Bridge.UpdateLight("/lights/"+hlight.Id, data)
	if err != nil {
		return err
	}

	// Check if server threw an error
	if m, ok := res[0]["error"]; ok {
		// Go complains if I use JSON here, no aliasing
		mobj, ok := m.(map[string]interface{})
		if !ok {
			// Unexpected response, not a matching API
			panic("Unexpected response from server")
		}

		return errors.New("Error when setting name: " +
			mobj["description"].(string))
	}

	return nil
}

// Unfortunately this is only a guideline. It takes the bridge a while to
// refresh
func (hlight *HueLight) IsReachable() (bool, error) {
	// Get itself, just to check its state.
	l, err := hlight.Bridge.GetLight(hlight.Id)
	if err != nil {
		return false, err
	}

	return l.State.Reachable, nil
}
