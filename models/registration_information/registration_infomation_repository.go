package registration_information

import (
	"errors"
	"github.com/jinzhu/gorm"
)

type datastore struct {
	db *gorm.DB
}

func NewDataStore(d *gorm.DB) *datastore {
	return &datastore{db: d}
}

func (d *datastore) Get() (*RegistrationInformation, error) {
	var ri []RegistrationInformation

	if err := d.db.Order("created_at desc").Limit(1).Find(&ri).Error; err != nil {
		return nil, err
	}

	if len(ri) == 0 {
		return nil, errors.New("there is not location info")
	}

	return &ri[len(ri)-1], nil
}

func (d *datastore) Store(ri *RegistrationInformation) (*RegistrationInformation, error) {
	if err := d.db.Create(&ri).Error; err != nil {
		return nil, err
	}

	return ri, nil
}
