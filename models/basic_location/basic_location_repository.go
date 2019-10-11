package basic_location

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
	return &datastore{col: c.Database("pss").Collection("basic_location")}
}

func (d *datastore) Get() (*BasicLocation, error) {
	bl:= BasicLocation{}

	findOptions := options.FindOne().SetSort(bson.D{{"createdat", -1}})
	err := d.col.FindOne(nil, bson.D{}, findOptions).Decode(&bl)

	if err == mongo.ErrNoDocuments {
		return nil, utilities.NotFoundError
	} else if err != nil {
		return nil, err
	}

	return &bl, nil
}

func (d *datastore) Store(location *BasicLocation) (*BasicLocation, error) {
	_, err := d.col.InsertOne(context.Background(), location)

	if err != nil {
		return nil, err
	}

	return location, nil
}
