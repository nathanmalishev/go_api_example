package main

import (
	"log"
	"net/http"

	config "github.com/nathanmalishev/taskmanager/config"
)

func main() {

	server := &http.Server{
		Addr:    config.AppConfig.Server,
		Handler: InitRoutes(),
	}
	log.Println("Listening on ", config.AppConfig.Server)
	log.Fatal(server.ListenAndServe())
}
