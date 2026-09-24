package config

import (
	"encoding/json"
	"flag"
	"log/slog"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/v2/mongo"
)


type Application struct {
	IsDev bool
	Db *mongo.Client
	Logger *slog.Logger
	MongoUrl string
	Port string
}

func InitApplication() (*Application) {
	var conf Application
	
	flag.StringVar(&conf.Port, "port", ":8081", "HTTP Port Number")
	flag.StringVar(&conf.MongoUrl, "mongoUrl", "", "Mongo Connection String")
	flag.BoolVar(&conf.IsDev, "isDev", false, "Inform runtime if you're in dev environment by supplying true or false")

	flag.Parse()

	
	slogOptions := &slog.HandlerOptions{
    	Level: slog.LevelInfo,
	}

	if conf.IsDev {
		slogOptions.Level = slog.LevelDebug
		slogOptions.AddSource = true
	}

	conf.Logger = slog.New(slog.NewJSONHandler(os.Stdout, slogOptions))


	if conf.IsDev {
		conf.Logger.Debug("Hello Dev!")
	}

	if conf.MongoUrl == "" {
		conf.Logger.Error("Missing required Mongo URL, set '-mongoUrl' when running the application")
		os.Exit(1)
	}

	return &conf
}

func (app *Application) ServerError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method 	= r.Method
		uri 	= r.URL.RequestURI()
	)

	app.Logger.Error(err.Error(), "method", method, "uri", uri)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

type ClientResponse struct {
	Ok      bool    `json:"ok"`
	Status  int     `json:"status"`
	Message string  `json:"msg"`
}
// The clientError helper sends a specific status code and corresponding description to the user
func (app *Application) ClientError(w http.ResponseWriter, status int, message string) {
	response := ClientResponse{
		Ok: false,
		Status: status,
		Message: message,
	}

	jsonRes, err := json.Marshal(response)
	if err != nil {
		http.Error(w, http.StatusText(status), status)
		return
	}

    http.Error(w, string(jsonRes), status)
}