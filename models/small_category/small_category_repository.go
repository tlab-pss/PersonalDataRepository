package small_category

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
	return &datastore{col: c.Database("pss").Collection("small_category")}
}

func (d *datastore) All() (*[]SmallCategory, error) {
	// TODO: マスタデータを個人ごとに持っていいのか？

	sc := make([]SmallCategory, 0)

	h, err := d.col.Find(nil, bson.D{})

	if err != nil {
		return nil, err
	}

	if err := h.All(context.Background(), &sc); err != nil {
		return nil, err
	}

	return &sc, nil
}

func (d *datastore) Find(id string) (*SmallCategory, error) {
	sc := SmallCategory{}

	findOptions := options.FindOne()
	err := d.col.FindOne(nil, bson.D{{"id", id}}, findOptions).Decode(&sc)

	if err == mongo.ErrNoDocuments {
		return nil, utilities.NotFoundError
	} else if err != nil {
		return nil, err
	}

	return &sc, nil
}
