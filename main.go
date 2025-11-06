package main

import (
	"fmt"
	"glucosestream/device"
	"time"
)

func main() {

	// set up channel for receiving readings from simulated device
	readings := make(chan device.Reading)

	// start simulating device in a separate goroutine
	go device.SimulateDevice("device-1", readings)

	for reading := range readings {
		fmt.Printf("Reading: %+v\n", reading)
		time.Sleep(100 * time.Millisecond)
	}
}
