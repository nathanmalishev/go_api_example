package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/nathanmalishev/taskmanager/common"
	"github.com/nathanmalishev/taskmanager/models"
	"golang.org/x/crypto/bcrypt"
)

type CreatedUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	JWT      string `json:"jwt"`
}

/* Registers a user that includes all the common fields from models.User */
func Register(authMod common.Authorizer, d models.UserStore, w http.ResponseWriter, r *http.Request) {

	//get data from request
	decoder := json.NewDecoder(r.Body)
	body := models.User{}
	err := decoder.Decode(&body)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid user data", http.StatusInternalServerError)
		return
	}

	//create new user
	userId, err := d.CreateUser(body)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid user data", http.StatusInternalServerError)
		return
	}

	//create JWT
	jwt, err := authMod.GenerateJWT(
		body.UserName,
		userId,
	)
	if err != nil {
		common.DisplayAppError(w, err, "fail up", http.StatusInternalServerError)
		return
	}
	returnUser := CreatedUser{
		body.UserName,
		body.Email,
		jwt,
	}
	common.WriteJson(w, "Succesfully registered user", returnUser, http.StatusCreated)
}

func GetUser(d *models.DataStore, w http.ResponseWriter, r *http.Request) {

	//get user info from request
	userInfo := r.Context().Value("userContext").(*common.UserClaims)
	common.WriteJson(w, "success", userInfo, http.StatusOK)
}

// Login, attempts to log in a user and writes the response back the ResponseWriter
func Login(authMod common.Authorizer, d models.UserStore, w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	body := models.User{}
	if err := decoder.Decode(&body); err != nil {
		common.DisplayAppError(w, err, "Invalid user data", http.StatusInternalServerError)
		return
	}

	user, err := d.FindUser(body)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid user data", http.StatusInternalServerError)
		return
	}

	// compare password
	err = bcrypt.CompareHashAndPassword(user.HashPassword, []byte(body.Password))
	if err != nil {
		common.DisplayAppError(w, err, "Incorrect username and password", http.StatusInternalServerError)
		return
	}

	// generate a fresh JWT
	jwt, err := authMod.GenerateJWT(
		user.UserName,
		user.Id,
	)
	if err != nil {
		common.DisplayAppError(w, err, "Please try again later", http.StatusInternalServerError)
		return
	}
	returnUser := CreatedUser{
		user.UserName,
		user.Email,
		jwt,
	}

	common.WriteJson(w, "success", returnUser, http.StatusOK)

}
