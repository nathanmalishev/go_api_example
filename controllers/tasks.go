package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/nathanmalishev/taskmanager/common"
	"github.com/nathanmalishev/taskmanager/models"
)

func GetAllTasks(d *models.DataStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tasks, err := d.GetAllTasks()
		if err != nil {
			common.DisplayAppError(w, err, common.FetchError, 500)
			return
		}
		tasksJson, err := json.Marshal(&tasks)
		err = errors.New("error")
		if err != nil {
			common.DisplayAppError(w, err, common.InvalidData, 500)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		w.Write(tasksJson)

	})
}
