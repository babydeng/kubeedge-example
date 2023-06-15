package main

import (
	lightdriver "github/babydeng/kubeedge-example/led-raspberrypi/light_driver"
	"time"
)

func main() {
	for {
		lightdriver.TurnON(18)
		time.Sleep(1 * time.Second)
		lightdriver.TurnOff(18)
		time.Sleep(1 * time.Second)
	}
}
