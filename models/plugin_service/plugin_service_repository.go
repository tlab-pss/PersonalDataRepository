package plugin_service

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type datastore struct {
	col *mongo.Collection
}

func NewDataStore(c *mongo.Client) *datastore {
	return &datastore{col: c.Database("pss").Collection("plugin_service")}
}

func (d *datastore) All() (*[]PluginService, error) {
	ps := make([]PluginService, 0)

	h, err := d.col.Find(nil, bson.D{})

	if err != nil {
		return nil, err
	}

	if err := h.All(context.Background(), &ps); err != nil {
		return nil, err
	}

	return &ps, nil
}

func (d *datastore) Store(service *PluginService) (*PluginService, error) {
	_, err := d.col.InsertOne(context.Background(), service)

	if err != nil {
		return nil, err
	}

	return service, nil
}
