package location

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Location struct {
	ID        string
	Latitude  string
	Longitude string
	CreatedAt time.Time
}

type datastore struct {
	db *gorm.DB
}

func (d *datastore) GetLatest() (*Location, error) {
	var la []Location

	if err := d.db.Find(&la).Error; err != nil {
		return nil, err
	}

	if len(la) == 0 {
		return nil, errors.New("There is not location info")
	}

	return &la[len(la)-1], nil
}

func (d *datastore) Store(location *Location) (*Location, error) {
	if err := d.db.Create(&location).Error; err != nil {
		return nil, err
	}

	return location, nil
}
