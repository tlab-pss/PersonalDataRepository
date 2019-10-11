package hotpepper

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type datastore struct {
	col *mongo.Collection
}

func NewDataStore(c *mongo.Client) *datastore {
	return &datastore{col: c.Database("pss").Collection("hotpepper")}
}

func (d *datastore) All() (*[]Intake, error) {
	i := make([]Intake, 0)

	h, err := d.col.Find(nil, bson.D{})

	if err != nil {
		return nil, err
	}

	if err := h.All(context.Background(), &i); err != nil {
		return nil, err
	}

	return &i, nil
}

func (d *datastore) Store(intake *Intake) (*Intake, error) {
	_, err := d.col.InsertOne(context.Background(), intake)

	if err != nil {
		return nil, err
	}

	return intake, nil
}

// todo: 日付のfindも必要だよねー
