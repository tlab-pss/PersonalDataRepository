package location

import "time"

type Location struct {
	ID             string
	Latitude       string
	Longitude      string
	Transportation string
	CreatedAt      time.Time
}
