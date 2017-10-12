package main

import (
	"log"
	"net/http"

	"github.com/nathanmalishev/taskmanager/common"
	"github.com/nathanmalishev/taskmanager/models"
	mgo "gopkg.in/mgo.v2"
)

func main() {

	/*  create data store for the app */
	store := models.CreateStore(&mgo.DialInfo{
		Addrs:    []string{common.AppConfig.MongoServer},
		Username: common.AppConfig.MongoUsername,
		Database: common.AppConfig.DbName,
		Password: common.AppConfig.MongoPassword,
	})

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: InitRoutes(store), // routes needs a copy of the store
	}
	log.Println("Listening on ", common.AppConfig.Server)
	log.Fatal(server.ListenAndServe())
}
