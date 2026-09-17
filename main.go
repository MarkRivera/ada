package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	mongo "github.com/MarkRivera/ada/mongo"
	"github.com/google/uuid"
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

	http.HandleFunc("GET /", healthcheckHandler)


	// Guests should be able to view register and login pages
	http.HandleFunc("POST /register", registerHandler)
	http.HandleFunc("POST /login", loginHandler)

	// Users need to be able to logout, view profile, update profile
	http.HandleFunc("GET /logout", logoutHandler)
	http.HandleFunc("GET /profile/{id}", viewProfileHandler)
	http.HandleFunc("PATCH /profile/{id}/edit", editProfileHandler)


    // Uploads
	// Users should be able to provision an upload, check their upload's status, delete an upload, update metadata
	// Users should be able to view videos uploaded by other users and themselves
	http.HandleFunc("POST /upload", uploadHandler)
	http.HandleFunc("GET /upload/status", statusHandler)



	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprintf(w, "{ \"success\": true }")
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Register User")
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Login User")
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Logout User")
}

func viewProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "View Profile")
}

func editProfileHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Edit Profile")
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	currentTime := time.Now().UTC()

	// What will come from the user?
	// Title, Desc and technically OwnerId

	// What will come from other systems?
	// VideoId, ManifestId, Thumbnail

	// View Count always defaults to 0

	metadata := mongo.VideoMetadata{
		Id: uuid.New(),
		Last_Updated: currentTime,
		Create_Date: currentTime,
		Published_At: currentTime,
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status Handler")
}

