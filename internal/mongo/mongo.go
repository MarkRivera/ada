package ada

import (
	"context"
	"log"
	"time"

	"github.com/MarkRivera/ada/internal/config"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/v2/bson"
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
	FileType string; // MP4
	Chunks uint;
	Status string; // Pending, Uploading, Error, Complete, Cancelled
}


type VideoRepository struct {
	collection *mongo.Collection
}

func InitializeMongoClient(app *config.Application) error {
	clientOptions := options.Client().ApplyURI(app.MongoUrl)

	app.Logger.Info("Attempting Mongo DB Connection")
	mongoClient, err := mongo.Connect(clientOptions)
	if err != nil {
		return err
	}
	app.Logger.Info("Mongo DB Connection Successful!")
	app.Db = mongoClient
	
	return nil;
}

func NewMongoRepository(client *mongo.Client) *VideoRepository {
	return &VideoRepository{
		collection: client.Database("video-metadata-db").Collection("videos"),
	}
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

func (repo *VideoRepository) InsertVideoMetadata(ctx context.Context, item VideoMetadata) (*mongo.InsertOneResult, error) {
	result, err := repo.collection.InsertOne(ctx, item)
	if err != nil {
		log.Printf("There was an issue while inserting into the collection: %s", err)
		return nil, err
	}


	log.Print("Inserted document with this id: ", result.InsertedID)
	return result, nil
}

func UpdateVideoMetadata(client *mongo.Client, id bson.ObjectID) {} // TODO: What is the update parameter look like?

func DeleteVideoMetadata(client *mongo.Client, id bson.ObjectID) {}