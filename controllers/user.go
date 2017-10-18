package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/nathanmalishev/taskmanager/common"
	"github.com/nathanmalishev/taskmanager/models"
)

func Register(authMod *common.Auth, d *models.DataStore, w http.ResponseWriter, r *http.Request) {

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
		body.Role,
		userId,
	)
	if err != nil {
		common.DisplayAppError(w, err, "fail up", http.StatusInternalServerError)
		return
	}

	createdUser := struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Role     string `json:"role"`
		JWT      string `json:"jwt"`
	}{
		body.UserName,
		body.Email,
		body.Role,
		jwt,
	}
	common.WriteJson(w, "Succesfully registered user", createdUser, http.StatusCreated)
}
