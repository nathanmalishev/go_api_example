package common

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
	mgo "gopkg.in/mgo.v2"
)

func HandleDb(sess *mgo.Session) negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		// copy mgo.Session and put into context
		fmt.Println("cop mgo.session")
		next(w, r)
	})
}
