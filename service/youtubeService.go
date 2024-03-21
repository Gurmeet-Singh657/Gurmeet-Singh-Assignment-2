package service

import (
	"Assignment/handler"
	"Assignment/model"
	"Assignment/util"
	"context"
	"log"
	"os"
	"sync"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/api/googleapi"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"github.com/joho/godotenv"
)

var (
	mongoClient   *mongo.Client
	youtubeClient *youtube.Service
	apiKeys       []string
	apiKeyIndex   int
	apiKeyMutex   sync.Mutex
)

/************* Setting Up Youtube Client *****************/
/** 1. Set up Mongodb Url **/
/** 2. Check Available API Key **/
/** 3. Set up Youtube Client **/

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file, Please add a .env file if not Added")
	}

	mongoURI, exists := os.LookupEnv("MONGODB_URL")
	if !exists {
		log.Fatal("Mongodb Url Not Set Properly")
	}

	// Getting all available API Keys
	apiKeys = util.GetAPIKeys()

	// Connecting to Mongodb
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatalf("Error creating MongoDB client: %v", err)
	}

	mongoClient = client

	// Trying Youtube Client with first api Key
	callYoutubeClient(apiKeys[0])
}

func callYoutubeClient(apiKey string) {
	var err error
	youtubeClient, err = youtube.NewService(context.Background(), option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Unable to create YouTube Client: %v", err)
	}
}

func StoreYoutubeDataAsynchronously(query string) error {
	call := youtubeClient.Search.List([]string{"snippet"}).
		Q(query).
		MaxResults(20)

	response, err := call.Do()

	videos, _ := util.GetFormattedVideos(response)

	if err != nil {
		if apiErr, ok := err.(*googleapi.Error); ok && apiErr.Code == 403 {
			log.Printf("YouTube API Error: %v", apiErr)
			switchAPIKey()
			return StoreYoutubeDataAsynchronously(query)
		}
		return err
	}

	if err := handler.AddVideos(videos, mongoClient); err != nil {
		return err
	}

	return nil
}

// switching the api keys in case if one api key is not working
func switchAPIKey() {
	apiKeyMutex.Lock()
	defer apiKeyMutex.Unlock()

	apiKeyIndex = (apiKeyIndex + 1) % len(apiKeys)
	callYoutubeClient(apiKeys[apiKeyIndex])
}

func GetVideosHandler(page, pageSize int) ([]model.Video, error) {

	videos, err := handler.GetVideos(page, pageSize, mongoClient)

	if err != nil {
		log.Printf("Can't Fetch videos %v", err)
		return nil, err
	}

	return videos, nil
}
