package ada

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type VideoMetadata struct {
	Id uuid.UUID;
	VideoId uuid.UUID;
	ManifestId uuid.UUID;
	Title string;
	View_Count int;
	Create_Date time.Time;
	Last_Updated time.Time;
	Published_At time.Time;
	OwnerId uuid.UUID;
	Description string;
	Thumbnail string;
}

func InitializeMongoClient() (*mongo.Client, error) {
	uri, ok := os.LookupEnv("MONGO_URI")
	if !ok {
		uri = "mongodb://localhost:27017"
	}

	clientOptions := options.Client().ApplyURI(uri)
	mongoClient, err := mongo.Connect(clientOptions)

	if err != nil {
		return nil, err
	}
	
	return mongoClient, nil;
}

func SetupMongoDB(client *mongo.Client) error {
	err := client.Ping(context.TODO(), nil)
	if err != nil {
		log.Print("Pinging the Mongo Cluster failed")
		return err
	}

	log.Println("Connected to Metadata DB")
	
	videoCollection := client.Database("video-metadata-db").Collection("videos")
	result, err := videoCollection.InsertOne(context.TODO(), VideoMetadata{
		Id: uuid.New(),
		VideoId: uuid.New(),
		ManifestId: uuid.New(),
		OwnerId: uuid.New(),
		View_Count: 0,
		Create_Date: time.Now(),
		Last_Updated: time.Now(),
		Title: "Test123",
		Description: "AlsoTesting123",
	})

	if err != nil {
		log.Printf("There was an issue while inserting into the collection: %s", err)
		return err
	}

	log.Print("Inserted document with this id: ", result.InsertedID)

	return nil
}

func InsertIntoMemoryStore(memoryStore []VideoMetadata, item VideoMetadata) {
	memoryStore[0] = item
}

func UpdateMemoryStoreItem(memoryStore []VideoMetadata, key string) {}

func DeleteMemoryStoreItem(memoryStore []VideoMetadata, index int) {}