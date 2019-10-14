package basic_location

import (
	"github.com/yuuis/PersonalDataRepository/models/location"
	"time"
)

type BasicLocation struct {
	ID        string
	House     location.Location
	Office    location.Location
	Route     []location.Location
	CreatedAt time.Time
}
