package models

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		UserName     string        `json:"username"`
		Email        string        `json:"email"`
		Role         string        `json:"role"`
		Password     string        `json:"password,omitempty"`
		HashPassword []byte        `json:"hashpassword,omitempty"`
	}
	UserStore interface {
		GetAll() ([]User, error)
		CreateUser(User) error
	}
)

const cNameUsers = "users"

/* User store database interactions */
func (d *DataStore) CreateUser(user User) error {

	err := d.C(cNameUsers).Insert(user)
	if err != nil {
		return err
	}
	return nil
}

func (d *DataStore) UserIndexs() error {
	index := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   false,
		Background: true, // See notes.
		Sparse:     true,
	}
	err := d.C(cNameUsers).EnsureIndex(index)
	if err != nil {
		return err
	}

	index = mgo.Index{
		Key:        []string{"email"},
		Unique:     true,
		DropDups:   false,
		Background: true, // See notes.
		Sparse:     true,
	}
	return d.C(cNameUsers).EnsureIndex(index)

}
