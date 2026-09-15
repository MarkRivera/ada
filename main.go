package main

import (
	"context"
	"log"

	mongo "github.com/MarkRivera/ada/mongo"
)



 

func main() {	
	mongoClient, err := mongo.InitializeMongoClient();
	if err != nil {
		log.Fatalf("Could not connect client to Mongo Cluster: %s", err)
	}
	defer mongoClient.Disconnect(context.TODO())

	err = mongo.SetupMongoDB(mongoClient)
	if err != nil {
		log.Fatal(err)
	}

	// Set up HTTP Handlers
}

