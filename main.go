package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello world, this is huego!")
	bridge := new(HueBridge)
	//err := bridge.GetIPAddress()
	err := bridge.SetupBridge()
	if err != nil {
		fmt.Println(err)
	}

	lights, err := bridge.GetLights()
	if err != nil {
		panic(err)
	}

	for _, l := range lights {
		err := l.Off()
		if err != nil {
			panic(err)
		}
	}
	time.Sleep(time.Second * 2)

	for _, l := range lights {
		err := l.On()
		if err != nil {
			panic(err)
		}
	}
	time.Sleep(time.Second * 2)
	for _, l := range lights {
		err := l.SetHue(1000)
		if err != nil {
			panic(err)
		}
	}
	time.Sleep(time.Second * 2)
	for _, l := range lights {
		err := l.SetBrightness(100)
		if err != nil {
			panic(err)
		}
	}
	time.Sleep(time.Second * 2)
	for _, l := range lights {
		err := l.Reset()
		if err != nil {
			panic(err)
		}
	}

}
