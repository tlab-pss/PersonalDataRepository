package health

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

func (d *datastore) GetLatest() (*Health, error) {
	var he []Health

	if err := d.db.Order("created_at desc").Limit(1).Find(&he).Error; err != nil {
		return nil, err
	}

	if len(he) == 0 {
		return nil, errors.New("there is not health info")
	}

	return &he[len(he)-1], nil
}

func (d *datastore) Store(health *Health) (*Health, error) {
	if err := d.db.Create(&health).Error; err != nil {
		return nil, err
	}

	return health, nil
}
