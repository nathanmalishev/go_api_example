package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/nathanmalishev/taskmanager/common"
	"github.com/nathanmalishev/taskmanager/models"
)

func Register(d *models.DataStore, w http.ResponseWriter, r *http.Request) {

	//get data from request
	decoder := json.NewDecoder(r.Body)
	body := models.User{}
	err := decoder.Decode(&body)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid user data", http.StatusInternalServerError)
		return
	}

	//create new user
	err = d.CreateUser(body)
	if err != nil {
		common.DisplayAppError(w, err, "Invalid user data", http.StatusInternalServerError)
		return
	}

	createdUser := struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     string `json:"role"`
	}{body.UserName, body.Email, body.Role}

	common.WriteJson(w, "Succesfully registered user", createdUser, http.StatusCreated)
}
