package common

import (
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

func readEnv(c *Config) {
	c.DbName = os.Getenv("DB_NAME")
	if c.DbName == "" {
		log.Println("DB_NAME not set using default, taskmanager")
		c.DbName = "taskmanager"
	}

	c.MongoServer = os.Getenv("MONGO_SERVER")
	if c.MongoServer == "" {
		log.Println("MONGO_SERVER not set using default, 127.0.0.1")
		c.MongoServer = "127.0.0.1"
	}

	c.Server = os.Getenv("SERVER_ADDR")
	if c.Server == "" {
		log.Println("SERVER_ADDR not setting using default, localhost:8080")
		c.Server = "localhost:8080"
	}

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
	AppConfig = &Config{}
	readEnv(AppConfig)
}
