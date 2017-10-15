package common

import (
	"context"
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

// Middleware that makes sure each request has a valid
// JWT from this server.
func WithAuth(a Authorizer) negroni.Handler {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		//check jwt
		auth := r.Header.Get("Authorization")
		if auth == "" {
			http.Error(rw, JwtHTTPError, http.StatusUnauthorized)
			return
		}

		// authorize
		token, err := a.Authorize(auth)
		if err != nil {
			fmt.Printf("Error decoding token:%s, err:%s\n", auth, err)
			http.Error(rw, JwtHTTPError, http.StatusUnauthorized)
			return
		}

		//handle claims
		username := ""
		if claims, ok := token.Claims.(AppClaims); ok {
			fmt.Printf("%v", claims.Username)
			username = claims.Username
		} else {
			fmt.Printf("Error decoding claims, token:%s, err:", token, err)
			http.Error(rw, JwtHTTPError, http.StatusUnauthorized)
			return
		}

		//set context and next
		ctx := context.WithValue(r.Context(), "user", username)
		next(rw, r.WithContext(ctx))
	})
}
