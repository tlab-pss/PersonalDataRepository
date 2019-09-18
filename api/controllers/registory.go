package controllers

import "github.com/jinzhu/gorm"

type Registry struct {
	db *gorm.DB
}

func NewRegistry(d *gorm.DB) *Registry {
	return &Registry{db: d}
}
