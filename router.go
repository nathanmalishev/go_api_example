package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func dummy() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Yet to be implemented")
		return
	})
}

func InitRoutes() http.Handler {
	router := mux.NewRouter().StrictSlash(false)

	/* User routes */
	router.Handle("/users/register", dummy()).Methods("POST")
	router.Handle("/users/login", dummy()).Methods("POST")

	/* Task routes  */
	taskRouter := mux.NewRouter().StrictSlash(false)
	taskRouter.Handle("/tasks", dummy()).Methods("GET")
	taskRouter.Handle("/tasks/{id}", dummy()).Methods("GET")
	taskRouter.Handle("/tasks/{id}", dummy()).Methods("DELETE")
	taskRouter.Handle("/tasks", dummy()).Methods("POST")
	taskRouter.Handle("/tasks/{id}", dummy()).Methods("PUT")
	taskRouter.Handle("/tasks/users/{id}", dummy()).Methods("GET")

	/* Notes routes  */
	notesRouter := mux.NewRouter().StrictSlash(false)
	notesRouter.Handle("/notes", dummy()).Methods("GET")
	notesRouter.Handle("/notes/{id}", dummy()).Methods("GET")
	notesRouter.Handle("/notes/{id}", dummy()).Methods("DELETE")
	notesRouter.Handle("/notes", dummy()).Methods("POST")
	notesRouter.Handle("/notes/{id}", dummy()).Methods("PUT")
	notesRouter.Handle("/notes/tasks/{id}", dummy()).Methods("GET")

	/* middleware */
	common := negroni.New(
		negroni.NewLogger(),
	)

	// will add auth middleware to these routes soon
	router.PathPrefix("/notes").Handler(negroni.New(
		negroni.Wrap(notesRouter),
	))
	router.PathPrefix("/tasks").Handler(negroni.New(
		negroni.Wrap(taskRouter),
	))

	// common wraps all routes in default middleware
	// this includes all API hits
	common.UseHandler(router)
	return common
}
