package big_category

import (
	"context"
	"github.com/yuuis/PersonalDataRepository/api/utilities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type datastore struct {
	col *mongo.Collection
}

func NewDataStore(c *mongo.Client) *datastore {
	return &datastore{col: c.Database("pss").Collection("big_category")}
}

func (d *datastore) All() (*[]BigCategory, error) {
	// TODO: マスタデータを個人ごとに持っていいのか？

	bc := make([]BigCategory, 0)

	h, err := d.col.Find(nil, bson.D{})

	if err != nil {
		return nil, err
	}

	if err := h.All(context.Background(), &bc); err != nil {
		return nil, err
	}

	return &bc, nil
}

func (d *datastore) Find(id string) (*BigCategory, error) {
	bc := BigCategory{}

	findOptions := options.FindOne()
	err := d.col.FindOne(nil, bson.M{"id": id}, findOptions).Decode(&bc)

	if err == mongo.ErrNoDocuments {
		return nil, utilities.NotFoundError
	} else if err != nil {
		return nil, err
	}

	return &bc, nil
}
