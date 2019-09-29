package location

import "time"

type Location struct {
	ID             string
	Latitude       float64
	Longitude      float64
	CreatedAt      time.Time
}
