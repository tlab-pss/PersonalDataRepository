package conversation

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
	return &datastore{col: c.Database("pss").Collection("conversation")}
}

func (d *datastore) Get() (*Conversation, error) {
	c := Conversation{}

	findOptions := options.FindOne().SetSort(bson.D{{"createdat", -1}})
	err := d.col.FindOne(nil, bson.D{}, findOptions).Decode(&c)

	if err == mongo.ErrNoDocuments {
		return nil, utilities.NotFoundError
	} else if err != nil {
		return nil, err
	}

	return &c, nil
}

func (d *datastore) FindByTransactionId(id string) (*[]Conversation, error) {
	c := make([]Conversation, 0)

	findOptions := options.Find()
	cur, err := d.col.Find(nil, bson.M{"transactionid": id}, findOptions)

	if err == mongo.ErrNoDocuments {
		return nil, utilities.NotFoundError
	} else if err != nil {
		return nil, err
	}

	if err := cur.All(context.Background(), &c); err != nil {
		return nil, err
	}

	return &c, nil
}

func (d *datastore) Store(conversation *Conversation) (*Conversation, error) {
	_, err := d.col.InsertOne(context.Background(), conversation)

	if err != nil {
		return nil, err
	}

	return conversation, nil
}
