package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/juliotorresmoreno/server-api/controllers"
)

func NewRouter(next http.Handler) http.Handler {
	server := mux.NewRouter().StrictSlash(false)
	server.HandleFunc("/api/users", controllers.GetUsers(next)).Methods("GET")
	server.HandleFunc("/api/users/{user}", controllers.PutUser(next)).Methods("PUT")
	server.PathPrefix("/").HandlerFunc(router(next))
	return server
}

func router(next http.Handler) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	}
}
