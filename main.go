package main

import (
	"context"
	"fmt"
	"glucosestream/device"
	"os"
	"os/signal"
	"time"
)

func fanIn(channels ...<-chan device.Reading) <-chan device.Reading {
	out := make(chan device.Reading)
	done := make(chan struct{})

	go func() {
		for _, ch := range channels {
			go func(c <-chan device.Reading) {
				for reading := range c {
					out <- reading
				}
				done <- struct{}{}
			}(ch)
		}
		for i := 0; i < len(channels); i++ {
			<-done
		}
		close(out)
	}()

	return out
}

func main() {

	// add context to allow terminal signals
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Listen for interrupt signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)

	// set up multiple channels (simulated multiple devices) that feed into a single output
	deviceCount := 3
	devices := make([]chan device.Reading, deviceCount)

	// start simulating device in a separate goroutine
	for i := 0; i < deviceCount; i++ {
		devices[i] = make(chan device.Reading)
		go device.SimulateDevice(ctx, fmt.Sprintf("device-%d", i+1), devices[i])
	}

	merged := fanIn(devices[0], devices[1], devices[2])

	// handle shutdown on interrupt signal
	go func() {
		<-sig
		fmt.Println("Shutting down")
		cancel()
	}()

	// read from merged channel

	for reading := range merged {
		fmt.Printf("Reading: %+v\n", reading)
		time.Sleep(100 * time.Millisecond)
	}
}
