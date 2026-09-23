package config

import (
	"flag"
	"log/slog"
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