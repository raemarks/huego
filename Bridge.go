package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var _ = fmt.Println

const bridgeDiscoveryIP = "https://www.meethue.com/api/nupnp"
const username = "newdeveloper"

type HueBridge struct {
	ipaddr string
	baddr  string
	user   string
	passwd string
}

// Get IP address of the bridge
func (bridge *HueBridge) GetIPAddress() error {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	client := &http.Client{Transport: tr}
	resp, err := client.Get(bridgeDiscoveryIP)
	if err != nil {
		return err
	}

	dec := json.NewDecoder(resp.Body)
	var responses []JSON
	err = dec.Decode(&responses)
	if err != nil {
		return err
	}

	bridge.ipaddr = "http://" + responses[0]["internalipaddress"].(string)
	fmt.Println("Ipaddress: " + bridge.ipaddr)
	return nil
}

//io.Copy(os.Stdout, resp.Body)
func (bridge *HueBridge) isSetup() (bool, error) {
	resp, err := http.Get(bridge.ipaddr + "/api/" + username)
	if err != nil {
		return false, err
	}

	// Hue is a piece of crap and you don't know if it will send an array or
	// not, so now I have to check the freaking first byte thus this mess
	mybuf := new(bytes.Buffer)
	io.Copy(mybuf, resp.Body)
	if mybuf.Bytes()[0] == '[' {
		dec := json.NewDecoder(mybuf)
		var responses []JSON
		err = dec.Decode(&responses)
		if err != nil {
			return false, err
		}

		if e, ok := responses[0]["error"]; ok {
			m := e.(JSON)
			if m["description"] == "unauthorized user" {
				// No error, but bridge is not already set up,
				// user does not exist.
				return false, nil
			} else {
				return false, errors.New("Unexpected error " +
					"description from bridge: " +
					m["description"].(string))
			}
		} else {
			return false, errors.New("Completely unknown issue")
		}

	} else {
		// Hue bridge already has the user registered, return true
		fmt.Println("Already set up.")
		return true, nil
	}
}

func (bridge *HueBridge) SetupBridge() error {
	if bridge.ipaddr == "" {
		bridge.GetIPAddress()
	}
	if b, e := bridge.isSetup(); b && e == nil {
		// Already set up, username already valid.
		return nil
	}

	data := make(JSON)
	data["devicetype"] = "test user"
	data["username"] = username

	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(data)
	// Need to POST twice, so copy before expending buf
	b := buf.Bytes()

	resp, err := http.Post(bridge.ipaddr+"/api", "application/json", buf)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("Bad status code from POST")
	}

	dec := json.NewDecoder(resp.Body)
	var responses []JSON
	err = dec.Decode(&responses)
	if err != nil {
		return err
	}

	if e, ok := responses[0]["error"]; ok {
		m := e.(JSON)
		// Only error should be the user needing to press the link button
		if m["description"] != "link button not pressed" {
			return errors.New("Error when setting up bridge: " +
				m["description"].(string))
		} else {
		}
	} else if _, ok := responses[0]["success"]; ok {
		// Good to go, no need to press link button
		return nil
	}

	rd := bytes.NewReader(b)
	// Need to press link button and POST again
	fmt.Println("Press link button on bridge, then press enter...")
	fmt.Scanln()

	resp, err = http.Post(bridge.ipaddr+"/api", "application/json", rd)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("Bad status code from POST")
	}

	dec = json.NewDecoder(resp.Body)
	err = dec.Decode(&responses)
	if err != nil {
		return err
	}

	if _, ok := responses[0]["success"]; ok {
	} else {
		m := responses[0]["error"].(JSON)
		return errors.New("Error when setting up bridge: " +
			m["description"].(string))
	}
	return nil
}

// Get all lights connected
func (bridge *HueBridge) GetLights() ([]Light, error) {
	resp, err := http.Get(bridge.ipaddr + "/api/" + username + "/lights")
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)
	var lights map[string]*HueLight
	err = dec.Decode(&lights)
	if err != nil {
		return nil, err
	}

	var lights_array []Light
	for k, l := range lights {
		l.Bridge = bridge
		l.Id = k
		lights_array = append(lights_array, l)
	}

	return lights_array, nil
}

func (bridge *HueBridge) UpdateLight(id string, data JSON) ([]JSON, error) {
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.Encode(data)
	req, err := http.NewRequest("PUT", bridge.ipaddr+"/api/"+
		username+"/lights/"+id+"/state", buf)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	dec := json.NewDecoder(resp.Body)
	var res []JSON
	err = dec.Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
