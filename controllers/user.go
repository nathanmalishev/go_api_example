package controllers

import (
	"net/http"

	mgo "gopkg.in/mgo.v2"
)

func RegisterAdapter(*mgo.Session) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
