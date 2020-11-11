package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	client *mongo.Client
}

func NewDatabase() (*Database, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:pass@localhost:27017"))
	if err != nil {
		return nil, err
	}

	err = client.Connect(context.TODO())
	if err != nil {
		return nil, err
	}

	return &Database{
		client: client,
	}, nil
}

func (d *Database) Close() error {
	return d.client.Disconnect(context.TODO())
}
