package handler

import (
	"Assignment/model"
	"context"
	"os"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func AddVideos(videos []model.Video, mongoClient *mongo.Client) error {
	databaseName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("COLLECTION_NAME")
	collection := mongoClient.Database(databaseName).Collection(collectionName)

	for _, video := range videos {
		filter := bson.D{{Key: "id", Value: video.ID}}

		update := bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "title", Value: video.Title},
				{Key: "description", Value: video.Description},
				{Key: "publishedat", Value: video.PublishedAt},
				{Key: "thumbnail", Value: video.Thumbnail},
			}},
		}

		if _, err := collection.UpdateOne(context.Background(), filter, update, options.Update().SetUpsert(true)); err != nil {
			return err
		}
	}

	return nil
}
