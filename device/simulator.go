package device

import (
	"context"
	"math/rand"
	"time"
)

// Reading represents a glucose reading from a device.
type Reading struct {
	ID           int
	DeviceID     string
	Timestamp    time.Time
	GlucoseLevel float64
}

func SimulateDevice(ctx context.Context, deviceID string, out chan<- Reading) {
	readingID := 1
	for {
		select {
		case <-ctx.Done():
			close(out)
			return
		default:
			// Simulate a glucose reading
			reading := Reading{
				ID:           readingID,
				DeviceID:     deviceID,
				Timestamp:    time.Now(),
				GlucoseLevel: 70 + rand.Float64()*100, // Random glucose level between 70 and 170
			}
			out <- reading
			readingID++
			time.Sleep(1 * time.Second)
		}

	}
}
