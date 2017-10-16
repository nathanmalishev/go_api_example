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

	common.WriteJson(w, http.StatusCreated, []byte("Successfully created user"))
}
