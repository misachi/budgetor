package api

import (
	"budgetor/utils"
	"net/http"

	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"
)

var conn = utils.Conn

func Handlers(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/v1").Subrouter()
	subRouter.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		ListUsers(db, w, r)
	}).Methods("GET")
	subRouter.HandleFunc("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
		GetUser(db, w, r)
	}).Methods("GET")
	subRouter.HandleFunc("/users/create", func(w http.ResponseWriter, r *http.Request) {
		CreateUser(db, w, r)
	}).Methods("POST")

	return subRouter
}
