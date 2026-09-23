package config

import (
	"flag"
	"log"

	"go.mongodb.org/mongo-driver/v2/mongo"
)


type Config struct {
	IsDev bool
	Db *mongo.Client
	MongoUrl string
	Port string
}

func InitConfig() (*Config) {
	var conf Config
	
	flag.StringVar(&conf.Port, "port", ":8081", "HTTP Port Number")
	flag.StringVar(&conf.MongoUrl, "mongoUrl", "", "Mongo Connection String")
	flag.BoolVar(&conf.IsDev, "IsDev", false, "Inform runtime if you're in dev environment by supplying true or false")

	flag.Parse()

	if conf.IsDev {
		log.Println("Hello Dev!")
	}

	if conf.MongoUrl == "" {
		log.Fatal("Missing required Mongo URL, set '-mongoUrl' when running the application")
	}

	return &conf
}