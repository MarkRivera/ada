package main

import (
	"context"
	"encoding/json"
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

type UploadRequestBody struct {
	Title string 			`json:"title"`
	Description string 		`json:"description"`
	FileType string 		`json:"fileType"`
	TotalChunks int 		`json:"totalChunks"`
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	
	// TODO REQUEST VALIDATION: 
	// Ownership, ensure videos are uploaded by users logged in

	var requestBody UploadRequestBody;
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestBody); err != nil {
		// Logging?
		fmt.Println(err)
		http.Error(w, "There was an issue processing your request. Please try again later", http.StatusInternalServerError)
		return
	}

	// TODO: REQUEST VALIDATION
	// What will come from the user?
	// Title, Desc

	currentTime := time.Now().UTC()




	// What will come from other systems?
	// ManifestId, Thumbnail

	// View Count always defaults to 0

	metadata := mongo.VideoMetadata{
		Id: uuid.New(),
		Last_Updated: currentTime,
		Create_Date: currentTime,
		Published_At: currentTime,
		Status: "pending",
		View_Count: 0,
		

		Title: requestBody.Title,
		Description: requestBody.Description,
		Chunks: uint(requestBody.TotalChunks),
		FileType: requestBody.FileType,


		VideoId: uuid.New(),
		ManifestId: uuid.New(),
	}
}

func statusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Status Handler")
}

