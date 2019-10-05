package user_like

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type datastore struct {
	col *mongo.Collection
}

func NewDataStore(c *mongo.Client) *datastore {
	return &datastore{col: c.Database("pss").Collection("user_like")}
}

func (d *datastore) All() (*[]UserLike, error) {
	ul := make([]UserLike, 0)

	h, err := d.col.Find(nil, bson.D{})

	if err != nil {
		return nil, err
	}

	if err := h.All(context.Background(), &ul); err != nil {
		return nil, err
	}

	return &ul, nil
}

func (d *datastore) Store(like *UserLike) (*UserLike, error) {
	_, err := d.col.InsertOne(context.Background(), like)

	if err != nil {
		return nil, err
	}

	return like, nil
}
