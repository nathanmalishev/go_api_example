package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		Id           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		UserName     string        `json:"username"`
		Email        string        `json:"email"`
		Role         string        `json:"role"`
		Password     string        `json:"password"`
		HashPassword []byte        `json:"hashPassword,omitempty"`
	}
	UserStore interface {
		GetAll() ([]User, error)
		CreateUser(User) error
	}
)

const cNameUsers = "users"

/* User store database interactions */
func (d *DataStore) CreateUser(user User) error {

	//Data check
	//Must have UserName, password, email & role --- in this example we can obv have weak passwords 😿
	if user.Password == "" || user.Role == "" || user.Email == "" || user.UserName == "" {
		return errors.New("Invalid user data, field is missing")
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.HashPassword = hash

	err = d.C(cNameUsers).Insert(user)
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
