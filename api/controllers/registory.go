package controllers

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Registry struct {
	db *mongo.Client
}

func NewRegistry(c *mongo.Client) *Registry {
	return &Registry{db: c}
}
