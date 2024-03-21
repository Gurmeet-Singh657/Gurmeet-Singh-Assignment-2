package handler

import (
	"Assignment/model"
	"context"
	"os"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetVideos(page, pageSize int, mongoClient *mongo.Client) ([]model.Video, error) {
	databaseName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")
	collection := mongoClient.Database(databaseName).Collection(collectionName)

	// Fetching the videos with page and pageSize
	// in reverse order
	options := options.Find().
		SetSkip(int64((page - 1) * pageSize)).
		SetLimit(int64(pageSize)).
		SetSort(bson.D{{Key: "publishedat", Value: -1}})

	cursor, err := collection.Find(context.Background(), bson.D{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var videos []model.Video
	for cursor.Next(context.Background()) {
		var video model.Video
		if err := cursor.Decode(&video); err != nil {
			return nil, err
		}
		videos = append(videos, video)
	}

	return videos, nil
}
