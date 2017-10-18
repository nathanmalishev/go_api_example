package controllers

import (
	"net/http"

	"github.com/nathanmalishev/taskmanager/common"
	"github.com/nathanmalishev/taskmanager/models"
)

// gets all tasks for a given user
func GetAllTasks(d models.DataStorer, w http.ResponseWriter, r *http.Request) {

	userClaims := r.Context().Value("userContext").(*common.UserClaims)

	tasks, err := d.GetAllTasksByUserId(userClaims.UserId)
	if err != nil {
		common.DisplayAppError(w, err, common.FetchError, 500)
		return
	}

	common.WriteJson(w, "success", &tasks, http.StatusOK)
}
