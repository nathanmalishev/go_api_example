package common

import (
	"context"
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
)

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

		userContext := UserClaims{}
		// handle claims
		if claims, ok := token.Claims.(AppClaims); ok {
			fmt.Printf("%v", claims.Username)
			userContext = UserClaims{
				Username: claims.Username,
				Role:     claims.Role,
				UserId:   claims.UserId,
			}
		} else {
			fmt.Printf("Error decoding claims, token:%s, err:", token, err)
			http.Error(rw, JwtHTTPError, http.StatusUnauthorized)
			return
		}

		//set context and next
		ctx := context.WithValue(r.Context(), "userContext", userContext)

		next(rw, r.WithContext(ctx))
	})
}
