package controllers

import (
	"net/http"

	"github.com/nathanmalishev/taskmanager/common"
	"github.com/nathanmalishev/taskmanager/models"
)

// With db wraps each controller that needs the db with a new session
// this is important to handle requests concurrently
// We want the actual function to recieve dataStore so i don't think you can middleware it
func WithDb(store models.DataStorer, fn func(models.DataStorer, http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		newStore := store.GetStore() // when we return the store, we copy the session
		defer newStore.Close()       // must close the session, or we will leave connections open
		fn(newStore, w, r)
	})
}

//We will need withAuth module, for register/login routes
//trying to not pollute global name space so going to need another middleware
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
