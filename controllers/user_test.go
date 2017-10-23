package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/nathanmalishev/go_api_example/common"
	"github.com/nathanmalishev/go_api_example/controllers"
	"github.com/nathanmalishev/go_api_example/models"
	mgo "gopkg.in/mgo.v2"
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

func TestLogin(t *testing.T) {
	r := mux.NewRouter()

	auth := &common.Auth{
		Secret:        common.AppConfig.JwtSecret,
		SigningMethod: jwt.SigningMethodHS512,
	}

	r.HandleFunc("/users", controllers.WithDbAndAuth(auth, Store, controllers.Register)).Methods("POST")
	r.HandleFunc("/users/login", controllers.WithDbAndAuth(auth, Store, controllers.Login)).Methods("POST")

	t.Run("ShouldBeAbleToRegister", UserShouldRegister(r))
	t.Run("UserShouldBeAbleToLogin", UserShouldLogin(r))
	t.Run("UserShouldNotBeAbleToLogin", UserShouldNotLogin(r))

}

func UserShouldRegister(r *mux.Router) func(t *testing.T) {
	return func(t *testing.T) {
		data := `{"username": "nathan", "password": "apple", "email":"nathan.email"}`
		req, err := http.NewRequest("POST", "/users", strings.NewReader(data))
		if err != nil {
			t.Error(err)
		}
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		if w.Code != 201 {
			t.Errorf("HTTP Status expected: 201, got: %d", w.Code)
		}
	}
}

func UserShouldNotLogin(r *mux.Router) func(t *testing.T) {
	return func(t *testing.T) {
		data := `{"username": "nathan", "password": "fakepassword"}`
		req, err := http.NewRequest("POST", "/users/login", strings.NewReader(data))
		if err != nil {
			t.Error(err)
		}
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		if w.Code != 500 {
			t.Errorf("HTTP Status expected: 500, got: %d", w.Code)
		}
	}
}
func UserShouldLogin(r *mux.Router) func(t *testing.T) {
	return func(t *testing.T) {
		data := `{"username": "nathan", "password": "apple"}`
		req, err := http.NewRequest("POST", "/users/login", strings.NewReader(data))
		if err != nil {
			t.Error(err)
		}
		w := httptest.NewRecorder()

		r.ServeHTTP(w, req)
		if w.Code != 200 {
			t.Errorf("HTTP Status expected: 200, got: %d", w.Code)
		}
		// if all correct fields are not present
		result := struct {
			Message string
			Data    controllers.CreatedUser
		}{}
		err = json.NewDecoder(w.Body).Decode(&result)
		if err != nil {
			t.Error(err)
		}
		if result.Data.Username != "nathan" || result.Data.Email != "nathan.email" {
			t.Fail()
		}
		if result.Data.JWT == "" || result.Message != "success" {
			t.Fail()
		}
	}
}
