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
	// This server is evolving to handle initial upload requests and user registration
    // Maybe in the future those domains can be split into their own services but for now
    // its fine.

    // Users
	// Guests should be able to view register and login pages
	// Users need to be able to logout, view profile, update profile

    // Uploads
	// Users should be able to provision an upload, check their upload's status, delete an upload, update metadata
	// Users should be able to view videos uploaded by other users and themselves
}

