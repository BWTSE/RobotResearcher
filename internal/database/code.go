package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func (d *Database) IsValidCode(code string) (bool, error) {
	cursor := d.client.Database("test").Collection("codes").FindOne(context.TODO(), bson.M{"code": code, "open": true})

	err := cursor.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
