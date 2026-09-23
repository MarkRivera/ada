package server

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/MarkRivera/ada/internal/config"
)

type HealthResponse struct {
	Success bool `json:"success"`
	Mongo string `json: "mongo"`
}

func HealthCheckHandler(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := HealthResponse{}
		response.Success = true
		response.Mongo = "Up"

		err := app.Db.Ping(r.Context(), nil)
		if err != nil {
			log.Printf("There was an issue while trying to reach the database: %s", err)
			response.Mongo = "Down"
			response.Success = false
		}

		w.Header().Set("Content-Type", "application/json")
		
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			log.Printf("Error while responding to health check request: %s", err)
			return
		}
	}
}
