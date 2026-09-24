package server

import (
	"encoding/json"
	"net/http"

	"github.com/MarkRivera/ada/internal/config"
)

type HealthResponse struct {
	Success bool `json:"success"`
	MongoDB string `json:"mongoStatus"`
}

func HealthCheckHandler(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := HealthResponse{}
		response.Success = true
		response.MongoDB = "Up"

		err := app.Db.Ping(r.Context(), nil)
		if err != nil {
			app.Logger.Error("There was an issue while trying to reach the database: " + err.Error())
			response.MongoDB = "Down"
			response.Success = false
		}

		w.Header().Set("Content-Type", "application/json")
		
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			app.ServerError(w, r, err)
		}
	}
}
