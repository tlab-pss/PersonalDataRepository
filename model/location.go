package model

import "time"

type Location struct {
	ID string
	Latitude string
	Longitude string
	CreatedAt time.Time
}
