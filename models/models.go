package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName    string        `json:"firstname"`
		LastName     string        `json:"lastname"`
		Email        string        `json:"email"`
		Password     string        `json:"password,omitempty"`
		HashPassword []byte        `json:"hashpassword,omitempty"`
	}
	TaskNote struct {
		Id          bson.ObjectId `bson:"_id,omitempty" json:"id"`
		TaskId      bson.ObjectId `json:"taskid"`
		Description string        `json:"description"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
	}
	UserStore interface {
		GetAll(*DataStore) ([]User, error)
	}
)

func (*DataStore) getAll() ([]User, error) {
	return nil, nil
}
