package controllers

import (
	"net/http"

	"github.com/nathanmalishev/taskmanager/common"
	"github.com/nathanmalishev/taskmanager/models"
)

// With db wraps each controller that needs the db with a new session
// this is important to handle requests concurrently
// This is only used by the task handlers, which only ever need the TaskStore implementation
// Going to build this function specifically for that as its then easier to test those functions
func WithDb(store models.DataStorer, fn func(models.TaskStore, http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newStore := store.GetStore() // when we return the store, we copy the session
		defer newStore.Close()       // must close the session, or we will leave connections open
		fn(newStore, w, r)
	})
}

//We will need withAuth module, for register/login routes
// trying not to populate namespace, so adding another wrapper to deliver db/auth
func WithDbAndAuth(
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
