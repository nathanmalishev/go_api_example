package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Server      string
	MongoServer string
}

var AppConfig Config

func init() {
	file, err := os.Open("config/config.json")
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	decoder := json.NewDecoder(file)
	config := Config{}
	if err := decoder.Decode(&config); err != nil {
		log.Fatal(err)
	}
	AppConfig = config
}
