package health

import "time"

type Health struct {
	ID        string
	Weight    float64
	Height    float64
	HeartRate int
	CreatedAt time.Time
}
