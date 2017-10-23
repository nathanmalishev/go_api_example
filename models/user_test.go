package models_test

import (
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/nathanmalishev/go_api_example/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var Store *models.DataStore

func TestMain(m *testing.M) {
	// startup
	Store = models.CreateStore(&mgo.DialInfo{
		Addrs:   []string{"127.0.0.1"},
		Timeout: time.Second * 5,
	}, "test_db")

	retCode := m.Run()

	// teardown
	Store.Session.DB("test_db").DropDatabase()
	Store.Close()

	os.Exit(retCode)
}

var globalUser = models.User{UserName: "nathan", Email: "nathan@gmail.com", Password: "test", HashPassword: []byte(""), Id: bson.NewObjectId()}

func TestFindUser(t *testing.T) {

	t.Run("CreateUser", CreateUser(Store))
	t.Run("UserShouldExist", UserShouldExist(Store))
	t.Run("UserShouldNotExist", UserShouldNotExist(Store))

}

func CreateUser(Store models.DataStorer) func(t *testing.T) {
	return func(t *testing.T) {

		user, err := Store.CreateUser(globalUser)
		if err != nil {
			t.Error(err)
		}
		if reflect.DeepEqual(user, globalUser) {
			t.Error()
		}
	}
}

func UserShouldExist(Store models.DataStorer) func(t *testing.T) {
	return func(t *testing.T) {
		user, err := Store.FindUser(models.User{UserName: "nathan"})
		if err != nil {
			t.Error(err)
		}
		if user.HashPassword != nil {
			if user.UserName == globalUser.UserName && user.Email == globalUser.Email && user.Id == globalUser.Id {
				t.Error()
			}
		} else {
			t.Error()
		}
	}
}

func UserShouldNotExist(Store models.DataStorer) func(t *testing.T) {
	return func(t *testing.T) {
		userExists := models.User{UserName: "nathan", Email: "nathan@gmail.com", Password: "test", HashPassword: []byte(""), Id: bson.NewObjectId()}
		if err := Store.C("users").Insert(userExists); err != nil {
			t.Error(err)
		}
		_, err := Store.FindUser(models.User{UserName: "nathan1"})
		if err != nil {
			if err.Error() != "not found" {
				t.Error(err)
			} else {
				return
			}
		}
		t.Fail()
	}
}
