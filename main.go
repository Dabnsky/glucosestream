package main

import (
	"fmt"
	"glucosestream/device"
	"time"
)

func main() {
	readings := make(chan device.Reading)
	go device.SimulateDevice("device-1", readings)

	for reading := range readings {
		fmt.Printf("Reading: %+v\n", reading)
		time.Sleep(100 * time.Millisecond)
	}
}
