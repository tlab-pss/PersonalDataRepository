package hotpepper

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
	return &datastore{col: c.Database("pss").Collection("hotpepper")}
}

func (d *datastore) Get() (*Intake, error) {
	i := Intake{}

	findOptions := options.FindOne().SetSort(bson.D{{"createdat", -1}})
	err := d.col.FindOne(nil, bson.D{}, findOptions).Decode(&i)

	if err == mongo.ErrNoDocuments {
		return nil, utilities.NotFoundError
	} else if err != nil {
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
