package basic

import (
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type Basic struct {
	ID        string
	Name      string
	Birthday  time.Time
	Gender    int
	Mail      string
	Weight    string // TODO: 将来的に身体関係のテーブルができたらそっちに引っ越し
	Height    string // TODO: 将来的に身体関係のテーブルができたらそっちに引っ越し
	CreatedAt time.Time
}

type datastore struct {
	db *gorm.DB
}

func NewDataStore(d *gorm.DB) *datastore {
	return &datastore{db: d}
}

func (d *datastore) Get() (*Basic, error) {
	var ba []Basic

	if err := d.db.Find(&ba).Error; err != nil {
		return nil, err
	}

	if len(ba) == 0 {
		return nil, errors.New("There is not basic info")
	}

	return &ba[len(ba)-1], nil
}

func (d *datastore) Store(basic *Basic) (*Basic, error) {
	if err := d.db.Create(&basic).Error; err != nil {
		return nil, err
	}

	return basic, nil
}
