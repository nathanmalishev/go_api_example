package models_test

import (
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/nathanmalishev/taskmanager/models"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/dbtest"
)

var server dbtest.DBServer

func TestMain(m *testing.M) {
	// startup
	d, _ := ioutil.TempDir(os.TempDir(), "mongotools-test")
	server = dbtest.DBServer{}
	server.SetPath(d)

	retCode := m.Run()

	// teardown
	server.Wipe()
	server.Stop()

	os.Exit(retCode)
}

var globalUser = models.User{UserName: "nathan", Email: "nathan@gmail.com", Password: "test", HashPassword: []byte(""), Id: bson.NewObjectId()}

func TestFindUser(t *testing.T) {
	dataStore := models.DataStore{}
	dataStore.Session = server.Session()

	t.Run("CreateUser", CreateUser(&dataStore))
	t.Run("UserShouldExist", UserShouldExist(&dataStore))
	t.Run("UserShouldNotExist", UserShouldNotExist(&dataStore))

	dataStore.Close() // close session
}

func CreateUser(dataStore *models.DataStore) func(t *testing.T) {
	return func(t *testing.T) {

		user, err := dataStore.CreateUser(globalUser)
		if err != nil {
			t.Error(err)
		}
		if reflect.DeepEqual(user, globalUser) {
			t.Error()
		}
	}
}

func UserShouldExist(dataStore *models.DataStore) func(t *testing.T) {
	return func(t *testing.T) {
		user, err := dataStore.FindUser(models.User{UserName: "nathan"})
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

func UserShouldNotExist(dataStore *models.DataStore) func(t *testing.T) {
	return func(t *testing.T) {
		userExists := models.User{UserName: "nathan", Email: "nathan@gmail.com", Password: "test", HashPassword: []byte(""), Id: bson.NewObjectId()}
		if err := dataStore.C("users").Insert(userExists); err != nil {
			t.Error(err)
		}
		_, err := dataStore.FindUser(models.User{UserName: "nathan1"})
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
