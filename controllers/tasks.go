package controllers

import (
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

	common.WriteJson(w, "success", &tasks, http.StatusOK)
}
