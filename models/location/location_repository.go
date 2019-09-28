package location

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
	return &datastore{col: c.Database("pss").Collection("location")}
}

func (d *datastore) GetLatest() (*Location, error) {
	lo := Location{}

	findOptions := options.FindOne().SetSort(bson.D{{"createdat", -1}})
	err := d.col.FindOne(nil, bson.D{}, findOptions).Decode(&lo)

	if err == mongo.ErrNoDocuments {
		return nil, utilities.NotFoundError
	} else if err != nil {
		return nil, err
	}

	return &lo, nil
}

func (d *datastore) Store(location *Location) (*Location, error) {
	_, err := d.col.InsertOne(context.Background(), location)

	if err != nil {
		return nil, err
	}

	return location, nil
}
