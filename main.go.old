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

	l := lights[0]
	lights[1].Off()
	lights[2].Off()
	l.Reset()
	time.Sleep(time.Second * 4)

	err = l.SetName("RaeLight")
	if err != nil {
		panic(err)
	}

	l.SetTransitionTime(0)
	l.SetColorIntHS(HSColorInt{65535, 254})
	time.Sleep(time.Second * 4)
	l.SetColorFloatHS(Amber)
	time.Sleep(time.Second * 4)
	l.SetColorFloatHS(AntiqueWhite)
	time.Sleep(time.Second * 4)
	l.SetColorFloatHS(Aqua)
	time.Sleep(time.Second * 4)
	l.SetColorFloatHS(Ash) //
	time.Sleep(time.Second * 4)
	l.SetColorFloatHS(Azure)
	time.Sleep(time.Second * 4)
	l.SetColorFloatHS(Black) //not black
	time.Sleep(time.Second * 4)
	l.SetColorFloatHS(Blue)
	time.Sleep(time.Second * 4)

	l.Reset()

	/*
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
	*/
}
