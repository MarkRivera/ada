package main

import (
	"context"
	"log"
	"net/http"

	"github.com/MarkRivera/ada/internal/auth"
	"github.com/MarkRivera/ada/internal/config"
	ada "github.com/MarkRivera/ada/internal/mongo"
	"github.com/MarkRivera/ada/internal/uploads"
)
 


func main() {
	config := config.InitConfig()

	
	mongoClient, err := ada.InitializeMongoClient(config.MongoUrl);
	if err != nil {
		log.Fatalf("Could not connect client to Mongo Cluster: %s", err)
	}
	defer mongoClient.Disconnect(context.Background())

	// err = ada.SetupMongoDB(mongoClient)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// env := &config.Env{ Db: mongoClient }

	// Set up HTTP Handlers
	// This server is evolving to handle initial upload requests and user registration
    // Maybe in the future those domains can be split into their own services but for now
    // its fine.


	mux := http.NewServeMux()
	
	// server.RegisterRoutes(env, mux)
	auth.RegisterRoutes(mux)
	uploads.RegisterRoutes(mux)

	log.Printf("Connecting to %s \n", config.Port)
	log.Fatal(http.ListenAndServe(config.Port, mux))
}


