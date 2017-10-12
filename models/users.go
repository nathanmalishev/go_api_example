package models

import "gopkg.in/mgo.v2/bson"

type (
	User struct {
		Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName    string        `json:"firstname"`
		LastName     string        `json:"lastname"`
		Email        string        `json:"email"`
		Password     string        `json:"password,omitempty"`
		HashPassword []byte        `json:"hashpassword,omitempty"`
	}
	UserStore interface {
		GetAll() ([]User, error)
	}
)

/* User store database interactions */
