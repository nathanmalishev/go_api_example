package common

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// generates a new JWT token
func GenerateJWT(name string, role string, signKey []byte) (string, error) {
	claims := struct {
		Role string
		Name string
		jwt.StandardClaims
	}{
		name,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 4).Unix(),
			Issuer:    "nathanmalishev/taskmanager",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	jwt, err := token.SignedString(signKey)
	if err != nil {
		return "", err
	}

	return jwt, nil
}
