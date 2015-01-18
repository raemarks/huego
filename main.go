package main

import (
		"fmt"
       )

func main() {
	fmt.Println("Hello world, this is huego!")
	bridge := new(Bridge)
	//err := bridge.GetIPAddress()
	err := bridge.SetupBridge()
	if err != nil {
		fmt.Println(err)
	}
}
