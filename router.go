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
	router.Handle("/users", withDbAndAuth(authModule, store, controllers.Register)).Methods("POST")
	router.Handle("/users/login", withDbAndAuth(authModule, store, controllers.Login)).Methods("POST")
	//router.Handle("/users", withDb(store, controllers.GetUser)).Methods("GET")

	/* Task routes  */
	taskRouter := mux.NewRouter().StrictSlash(false)
	taskRouter.Handle("/tasks", withDb(store, controllers.GetAllTasks)).Methods("GET")
	taskRouter.Handle("/tasks/{id}", dummy()).Methods("GET")
	taskRouter.Handle("/tasks/{id}", dummy()).Methods("DELETE")
	taskRouter.Handle("/tasks", dummy()).Methods("POST")
	taskRouter.Handle("/tasks/{id}", dummy()).Methods("PUT")
	taskRouter.Handle("/tasks/users/{id}", dummy()).Methods("GET")

	/* middleware */
	commonMidleware := negroni.New(
		negroni.NewLogger(),
	)

	// will add auth middleware to these routes soon
	router.PathPrefix("/tasks").Handler(negroni.New(
		common.WithAuth(authModule),
		negroni.Wrap(taskRouter),
	))
	// common wraps all routes in default middleware
	// this includes all API hits
	commonMidleware.UseHandler(router)

	return commonMidleware
}

// With db wraps each controller that needs the db with a new session
// this is important to handle requests concurrently
// We want the actual function to recieve dataStore so i don't think you can middleware it
func withDb(store models.DataStorer, fn func(models.DataStorer, http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newStore := store.GetStore() // when we return the store, we copy the session
		defer newStore.Close()       // must close the session, or we will leave connections open
		fn(newStore, w, r)
	})
}

//We will need withAuth module, for register/login routes
//trying to not pollute global name space so going to need another middleware
func withDbAndAuth(
	authModule common.Authorizer,
	store models.DataStorer,
	fn func(common.Authorizer, models.UserStore, http.ResponseWriter, *http.Request),
) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newStore := store.GetStore()
		defer newStore.Close()
		fn(authModule, models.UserStore(newStore), w, r)
	})
}
