package common

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Server        string
	MongoServer   string
	MongoUsername string
	MongoPassword string
	DbName        string
	JwtSecret     []byte
}

/* Global for common package is the AppConfig */
var AppConfig *Config

func readConfig(filename string) *Config {
	file, err := os.Open(filename)
	if err != nil {
		log.Println("no config file found, using defaults")
		return &Config{}
	}
	decoder := json.NewDecoder(file)

	config := Config{}
	if err := decoder.Decode(&config); err != nil {
		log.Fatal(err)
	}
	return &config
}

func readEnv(c *Config) {
	c.Server = os.Getenv("SERVER_ADDR")
	c.MongoUsername = os.Getenv("MONGO_USERNAME")
	c.MongoPassword = os.Getenv("MONGO_PASSWORD")
	if c.MongoPassword == "" || c.MongoUsername == "" {
		log.Println("Please note these env variables haven't been set, MONGO_USERNAME & MONGO_PASSWORD")
	}

	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		log.Fatal("You must set your env variable JWT_SECRET")
	}
	c.JwtSecret = []byte(jwtSecret)
}

func init() {
	AppConfig = readConfig("common/config.json")
	readEnv(AppConfig)
}
