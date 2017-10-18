package common

import (
	"errors"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2/bson"
)

type (
	Authorizer interface {
		GenerateJWT(string, string) (string, error)
		Authorize(string) (*jwt.Token, error)
	}
	Auth struct {
		Secret        []byte
		SigningMethod jwt.SigningMethod
	}

	AppClaims struct {
		UserClaims
		jwt.StandardClaims
	}

	//Used in middleware and attached to the context
	UserClaims struct {
		Role     string
		Username string
		UserId   bson.ObjectId
	}
)

// generates a new JWT token
func (a *Auth) GenerateJWT(name string, role string, userId bson.ObjectId) (string, error) {
	claims := AppClaims{
		UserClaims{
			Username: name,
			Role:     role,
			UserId:   userId,
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 4).Unix(),
			Issuer:    "nathanmalishev/taskmanager",
		},
	}

	token := jwt.NewWithClaims(a.SigningMethod, claims)
	jwt, err := token.SignedString(a.Secret)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

// Processes jwtString and returns process token if valid
func (a *Auth) Authorize(jwtStr string) (*jwt.Token, error) {

	token, err := jwt.ParseWithClaims(jwtStr, &AppClaims{}, func(token *jwt.Token) (interface{}, error) {
		return a.Secret, nil
	})
	// Check if there was an error in parsing...
	if err != nil {
		return nil, errors.New("Error parsing jwt token")
	}

	if a.SigningMethod != nil && a.SigningMethod.Alg() != token.Header["alg"] {
		message := fmt.Sprintf("Expected %s signing method but token specified %s",
			a.SigningMethod.Alg(),
			token.Header["alg"])
		return nil, errors.New(message)
	}

	// Check if the parsed token is valid...
	if !token.Valid {
		return nil, errors.New("Token is invalid")
	}

	return token, nil
}
