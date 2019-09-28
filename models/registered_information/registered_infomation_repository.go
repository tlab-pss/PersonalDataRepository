package registered_information

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
	return &datastore{col: c.Database("pss").Collection("registered_infomation")}
}

func (d *datastore) Get() (*RegisteredInformation, error) {
	ri := RegisteredInformation{}

	findOptions := options.FindOne().SetSort(bson.D{{"createdat", -1}})
	err := d.col.FindOne(nil, bson.D{}, findOptions).Decode(&ri)

	if err == mongo.ErrNoDocuments {
		return nil, utilities.NotFoundError
	} else if err != nil {
		return nil, err
	}

	return &ri, nil
}

func (d *datastore) Store(ri *RegisteredInformation) (*RegisteredInformation, error) {
	_, err := d.col.InsertOne(context.Background(), ri)

	if err != nil {
		return nil, err
	}

	return ri, nil
}
