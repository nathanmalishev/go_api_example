package common

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/urfave/negroni"
)

// Middleware that makes sure each request has a valid
// JWT from this server.
func WithAuth(a Authorizer) negroni.Handler {
	return negroni.HandlerFunc(func(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		//check jwt
		var token string
		tokens, ok := r.Header["Authorization"]
		if ok && len(tokens) >= 1 {
			token = tokens[0]
			token = strings.TrimPrefix(token, "Bearer ")
		}
		if token == "" {
			DisplayAppError(rw, errors.New(""), JwtHTTPError, http.StatusUnauthorized)
			return
		}

		// authorize
		jwtToken, err := a.Authorize(token)
		if err != nil {
			fmt.Printf("Error decoding token:%s, err:%s\n", jwtToken, err)
			DisplayAppError(rw, err, JwtHTTPError, http.StatusUnauthorized)
			return
		}

		userContext := UserClaims{}
		// handle claims
		if claims, ok := jwtToken.Claims.(*AppClaims); ok {
			userContext = UserClaims{
				Username: claims.Username,
				UserId:   claims.UserId,
			}
		} else {
			fmt.Printf("Error decoding claims, token:%s, err:", token, err)
			DisplayAppError(rw, err, JwtHTTPError, http.StatusUnauthorized)
			return
		}

		//set context and next
		ctx := context.WithValue(r.Context(), "userContext", &userContext)

		next(rw, r.WithContext(ctx))
	})
}
