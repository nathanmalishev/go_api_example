package main

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/nathanmalishev/go_api_example/common"
	"github.com/nathanmalishev/go_api_example/models"
	mgo "gopkg.in/mgo.v2"
)

func main() {

	/*  create data store for the app */
	store := models.CreateStore(&mgo.DialInfo{
		Addrs:    []string{common.AppConfig.MongoServer},
		Username: common.AppConfig.MongoUsername,
		Database: common.AppConfig.DbName,
		Password: common.AppConfig.MongoPassword,
		Timeout:  time.Second * 5,
	}, common.AppConfig.DbName)

	/* initialize indexs */
	err := store.InitIndexs()
	if err != nil {
		log.Fatal(err)
	}

	/* create authorizer used to decode/encode jwt's */
	authModule := &common.Auth{
		Secret:        common.AppConfig.JwtSecret,
		SigningMethod: jwt.SigningMethodHS512,
	}

	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: InitRoutes(store, authModule), // routes needs a copy of the store && authModule
	}
	log.Println("Listening on ", common.AppConfig.Server)
	log.Fatal(server.ListenAndServe())
}
