package main

import (
	"fmt"
	"glucosestream/device"
	"time"
)

func main() {

	// set up multiple channels (simulated multiple devices) that feed into a single output
	deviceCount := 3
	devices := make([]chan device.Reading, deviceCount)

	// start simulating device in a separate goroutine
	for i := 0; i < deviceCount; i++ {
		devices[i] = make(chan device.Reading)
		go device.SimulateDevice(fmt.Sprintf("device-%d", i+1), devices[i])
	}

	merged := fanIn(devices[0], devices[1], devices[2])

	for reading := range readings {
		fmt.Printf("Reading: %+v\n", reading)
		time.Sleep(100 * time.Millisecond)
	}
}
