package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"

	"budgetor/users"
	"budgetor/utils"
)

// CreateUser Add new user
func CreateUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	repo := users.NewUserRepository(db)
	user := &users.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(user); err != nil {
		utils.JsonResponse(w, utils.ErrorResponse{Error: err.Error()}, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	repo.AddUser(user)
	utils.JsonResponse(w, user, http.StatusCreated)
}

// ListUsers Get all users
func ListUsers(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	repo := users.NewUserRepository(db)
	utils.JsonResponse(w, repo.GetAllUsers([]users.User{}), http.StatusOK)
}

// GetUser Retrieve user given the user ID
func GetUser(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	repo := users.NewUserRepository(db)
	vars := mux.Vars(r)

	target, _ := strconv.Atoi(vars["id"])
	user, changed := repo.GetUser(users.User{}, target)

	if !changed {
		utils.JsonResponse(w, utils.ErrorResponse{Error: "User matching query not found"}, http.StatusNotFound)
	} else {
		utils.JsonResponse(w, user, http.StatusOK)
	}
}
