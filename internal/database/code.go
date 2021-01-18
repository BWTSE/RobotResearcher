package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"

	"go.mongodb.org/mongo-driver/bson"
)

func (d *Database) IsValidCode(code string) (bool, bool, error) {
	cursor := d.client.Database("test").Collection("codes").FindOne(context.TODO(), bson.M{"code": code, "open": true})

	err := cursor.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, false, nil
		}

		return false, false, err
	}

	var res struct {
		Ignore bool `bson:"ignore"`
	}

	err = cursor.Decode(&res)
	if err != nil {
		return false, false, err
	}

	return true, res.Ignore, nil
}
