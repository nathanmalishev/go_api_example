package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/nathanmalishev/taskmanager/common"
	"github.com/nathanmalishev/taskmanager/models"
)

func GetAllTasks(d *models.DataStore, w http.ResponseWriter, r *http.Request) {
	tasks, err := d.GetAllTasks()
	if err != nil {
		common.DisplayAppError(w, err, common.FetchError, 500)
		return
	}
	tasksJson, err := json.Marshal(&tasks)
	if err != nil {
		common.DisplayAppError(w, err, common.InvalidData, 500)
		return
	}

	common.WriteJson(w, http.StatusOK, tasksJson)
}
