package main

import (
	"context"
	"net/http"
	"os"

	"github.com/MarkRivera/ada/internal/auth"
	"github.com/MarkRivera/ada/internal/config"
	"github.com/MarkRivera/ada/internal/uploads"

	ada "github.com/MarkRivera/ada/internal/mongo"
	"github.com/MarkRivera/ada/internal/server"
)
 


func main() {
	app := config.InitApplication()

	app.Logger.Info("Application Starting")
	
	err := ada.InitializeMongoClient(app);
	if err != nil {
		app.Logger.Error("Failed to connect client to Mongo Cluster")
		app.Logger.Error(err.Error())
		os.Exit(1)
	}
	defer app.Db.Disconnect(context.Background())


	// Set up HTTP Handlers
	// This server is evolving to handle initial upload requests and user registration
    // Maybe in the future those domains can be split into their own services but for now
    // its fine.


	mux := http.NewServeMux()
	
	mux.HandleFunc("GET /{$}", server.HealthCheckHandler(app))

	mux.HandleFunc("POST /register", auth.RegisterHandler(app))
	mux.HandleFunc("POST /login", auth.LoginHandler(app))
	mux.HandleFunc("GET /logout", auth.LogoutHandler(app))
	mux.HandleFunc("GET /profile/{id}", auth.ViewProfileHandler(app))
	mux.HandleFunc("PATCH /profile/{id}/edit", auth.EditProfileHandler(app))


	mux.HandleFunc("POST /upload", uploads.UploadHandler(app))
	mux.HandleFunc("GET /upload/status", uploads.StatusHandler(app))

	

	app.Logger.Info("Listening for requests", "port", app.Port)
	err = http.ListenAndServe(app.Port, mux)
	app.Logger.Error(err.Error())
	os.Exit(1)
}


