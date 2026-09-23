package uploads

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"github.com/MarkRivera/ada/internal/config"
	ada "github.com/MarkRivera/ada/internal/mongo"
)

// Uploads
// Users should be able to provision an upload, check their upload's status, delete an upload, update metadata
// Users should be able to view videos uploaded by other users and themselves

func UploadHandler(app *config.Application) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
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
		// Title, Desc, FileType and TotalChunks
		if len(requestBody.Title) > MAX_TITLE_SIZE {
			http.Error(w, "Title is too long", http.StatusBadRequest)
			return
		}

		if len(requestBody.Description) ==  MAX_DESCRIPTION_SIZE {
			http.Error(w, "Description is too long", http.StatusBadRequest)
			return
		}

		if requestBody.FileType != "video/mp4" {
			http.Error(w, "We currently only support MP4 Videos", http.StatusBadRequest)
			return
		}

		currentTime := time.Now().UTC()




		// What will come from other systems?
		// ManifestId, Thumbnail

		// View Count always defaults to 0

		// called metadata
		_ = ada.VideoMetadata{
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

		// mongo.InsertVideoMetadata(r.Context(), metadata)
	}
}

func StatusHandler(app *config.Application) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Status Handler")
	}
}

