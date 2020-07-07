package db

import (
	"context"

	"github.com/fk/gqlplayground/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (db *Database) CreateLink(link *model.Link) (primitive.ObjectID, error) {
	result, err := db.Database.Collection("links").InsertOne(context.Background(), link)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (db *Database) GetLinkById(id string) (*model.Link, error) {
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	result := db.Database.Collection("links").FindOne(context.Background(), bson.M{"_id": userID})
	link := &model.Link{}
	result.Decode(link)
	return link, nil
}

func (db *Database) GetAllLinks() ([]*model.Link, error) {
	var results []*model.Link
	cur, err := db.Database.Collection("links").Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	for cur.Next(context.Background()) {
		var link model.Link
		err := cur.Decode(&link)
		if err != nil {
			return nil, err
		}
		results = append(results, &link)

	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	cur.Close(context.Background())
	return results, nil
}
