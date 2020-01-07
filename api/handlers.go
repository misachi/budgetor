package api

import (
	"budgetor/utils"
	"net/http"

	"github.com/gorilla/mux"
)

var conn = utils.Conn

func Handlers() *mux.Router {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/v1").Subrouter()
	subRouter.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		ListUsers(conn, w, r)
	}).Methods("GET")
	subRouter.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetUser(conn, w, r)
	}).Methods("GET")
	subRouter.HandleFunc("/users/create", func(w http.ResponseWriter, r *http.Request) {
		CreateUser(conn, w, r)
	}).Methods("POST")

	return subRouter
}
