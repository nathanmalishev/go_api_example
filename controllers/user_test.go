package controllers_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/nathanmalishev/taskmanager/common"
	"github.com/nathanmalishev/taskmanager/controllers"
	"github.com/nathanmalishev/taskmanager/models"
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

func TestLogin(t *testing.T) {
	r := mux.NewRouter()

	store := &models.DataStore{}
	store.Session = server.Session()

	auth := &common.Auth{
		Secret:        common.AppConfig.JwtSecret,
		SigningMethod: jwt.SigningMethodHS512,
	}

	r.HandleFunc("/users", controllers.WithDbAndAuth(auth, store, controllers.Register)).Methods("POST")
	r.HandleFunc("/users/login", controllers.WithDbAndAuth(auth, store, controllers.Login)).Methods("POST")

	t.Run("ShouldBeAbleToRegister", UserShouldRegister(r))
	t.Run("UserShouldBeAbleToLogin", UserShouldLogin(r))
	t.Run("UserShouldNotBeAbleToLogin", UserShouldNotLogin(r))

	store.Close()
}

func UserShouldRegister(r *mux.Router) func(t *testing.T) {
	return func(t *testing.T) {
		data := `{"username": "nathan", "password": "apple", "email":"nathan.email", "role":"student"}`
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
	}
}
