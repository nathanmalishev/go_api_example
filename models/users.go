package models

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	User struct {
		Id           bson.ObjectId `bson:"_id,omitempty" json:"id,omitempty"`
		UserName     string        `json:"username"`
		Email        string        `json:"email"`
		Password     string        `json:"password,omitempty"`
		HashPassword []byte        `json:"hashPassword,omitempty"`
	}
	UserStore interface {
		//GetAll() ([]User, error)
		CreateUser(User) (bson.ObjectId, error)
		FindUser(User) (User, error)
	}
)

const cNameUsers = "users"

/* User store database interactions */

//Creates a user in the database, must have email, username & password set
func (d *DataStore) CreateUser(user User) (bson.ObjectId, error) {

	//Data check
	//Must have UserName, password & email --- in this example we can have weak passwords or invalid emails 😿
	if user.Password == "" || user.Email == "" || user.UserName == "" {
		return "", errors.New("Invalid user data, field is missing")
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	user.HashPassword = hash
	user.Id = bson.NewObjectId()
	user.Password = ""

	err = d.C(cNameUsers).Insert(user)
	if err != nil {
		return "", err
	}
	return user.Id, nil
}

// Find a specific user
func (d *DataStore) FindUser(user User) (User, error) {
	if user.UserName == "" {
		return User{}, errors.New("Invaid user data")
	}

	foundUser := User{}
	err := d.C(cNameUsers).Find(bson.M{"username": user.UserName}).One(&foundUser)
	if err != nil {
		return User{}, err
	}

	return foundUser, nil
}

// Creates appropiate indexs for the user collection
// Creates an index on 'username' && 'email' field
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
