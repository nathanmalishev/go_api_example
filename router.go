package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nathanmalishev/taskmanager/common"
	"github.com/nathanmalishev/taskmanager/controllers"
	"github.com/nathanmalishev/taskmanager/models"
	"github.com/urfave/negroni"
)

func dummy() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Yet to be implemented")
		return
	})
}

func InitRoutes(store models.DataStorer, authModule common.Authorizer) http.Handler {
	router := mux.NewRouter().StrictSlash(false)

	/* User routes */
	router.Handle("/users", controllers.WithDbAndAuth(authModule, store, controllers.Register)).Methods("POST")
	router.Handle("/users/login", controllers.WithDbAndAuth(authModule, store, controllers.Login)).Methods("POST")
	//router.Handle("/users", withDb(store, controllers.GetUser)).Methods("GET")

	/* Task routes  */
	taskRouter := mux.NewRouter().StrictSlash(false)
	taskRouter.Handle("/tasks", controllers.WithDb(store, controllers.GetAllTasks)).Methods("GET")
	taskRouter.Handle("/tasks/{id}", dummy()).Methods("GET")
	taskRouter.Handle("/tasks/{id}", dummy()).Methods("DELETE")
	taskRouter.Handle("/tasks", dummy()).Methods("POST")
	taskRouter.Handle("/tasks/{id}", dummy()).Methods("PUT")
	taskRouter.Handle("/tasks/users/{id}", dummy()).Methods("GET")

	/* middleware */
	commonMidleware := negroni.New(
		negroni.NewLogger(),
	)

	router.PathPrefix("/tasks").Handler(negroni.New(
		common.WithAuth(authModule),
		negroni.Wrap(taskRouter),
	))

	// common wraps all routes in default middleware
	// this includes all API hits
	commonMidleware.UseHandler(router)

	return commonMidleware
}
