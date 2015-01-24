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

	for {
		b, e := l.IsReachable()
		if e != nil {
			panic(e)
		}
		fmt.Println("Reachable: %v", b)
		time.Sleep(time.Second)
	}

	err = l.SetName("RaeLight")
	if err != nil {
		panic(err)
	}

	l.SetColor(Amber)
	time.Sleep(time.Second * 4)
	l.SetColor(AntiqueWhite)
	time.Sleep(time.Second * 4)
	l.SetColor(Aqua)
	time.Sleep(time.Second * 4)
	l.SetColor(Ash) //
	time.Sleep(time.Second * 4)
	l.SetColor(Azure)
	time.Sleep(time.Second * 4)
	l.SetColor(Black) //not black
	time.Sleep(time.Second * 4)
	l.SetColor(Blue)
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
