package common

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Server        string
	MongoServer   string
	MongoUsername string
	MongoPassword string
	DbName        string
}

/* Global for common package is the AppConfig */
var AppConfig *Config

func readConfig(filename string) *Config {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)

	config := Config{}
	if err := decoder.Decode(&config); err != nil {
		log.Fatal(err)
	}
	return &config
}

func readEnv(c *Config) {
	c.MongoUsername = os.Getenv("MONGO_USERNAME")
	c.MongoPassword = os.Getenv("MONGO_PASSWORD")
}

func init() {
	AppConfig = readConfig("common/config.json")
	readEnv(AppConfig)
	fmt.Println(AppConfig)
}
