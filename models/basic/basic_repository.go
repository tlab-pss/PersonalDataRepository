package basic

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
	return &datastore{col: c.Database("pss").Collection("basic")}
}

func (d *datastore) Get() (*Basic, error) {
	ba := Basic{}

	findOptions := options.FindOne().SetSort(bson.D{{"createdat", -1}})
	err := d.col.FindOne(nil, bson.D{}, findOptions).Decode(&ba)

	if err == mongo.ErrNoDocuments {
		return nil, utilities.NotFoundError
	} else if err != nil {
		return nil, err
	}

	return &ba, nil
}

func (d *datastore) Store(basic *Basic) (*Basic, error) {
	_, err := d.col.InsertOne(context.Background(), basic)

	if err != nil {
		return nil, err
	}

	return basic, nil
}
