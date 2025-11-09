package main

import (
	"fmt"
	"glucosestream/device"
	"time"
)

func fanIn(channels ...<-chan device.Reading) <-chan device.Reading {
	out := make(chan device.Reading)
	for _, ch := range channels {
		go func(c <-chan device.Reading) {
			for reading := range c {
				out <- reading
			}
		}(ch)
	}
	return out
}

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

	for reading := range merged {
		fmt.Printf("Reading: %+v\n", reading)
		time.Sleep(100 * time.Millisecond)
	}
}
