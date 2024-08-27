package controller

import (
	"GoPractice/entry"
	"GoPractice/manager"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

type UserController struct {
	Manager *manager.UserManager
	Log     *logrus.Logger
}

func NewUserController(manager *manager.UserManager, log *logrus.Logger) *UserController {
	return &UserController{Manager: manager, Log: log}
}

func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId, err := strconv.ParseInt(vars["userId"], 10, 64)
	if err != nil {
		c.Log.Errorf("Error converting userId to int: %v", err)
		http.Error(w, "Invalid userId", http.StatusBadRequest)
		return
	}

	user, err := c.Manager.GetUserByID(userId)
	if err != nil {
		c.Log.Errorf("Error fetching user by ID: %v", err)
		http.Error(w, "Error fetching data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {

	var userEntry entry.UserEntry
	err := json.NewDecoder(r.Body).Decode(&userEntry)
	if err != nil {
		c.Log.Errorf("Error converting userId to int: %v", err)
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := c.Manager.CreateUser(userEntry)
	if err != nil {
		http.Error(w, "Error creating data", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(user)

}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Implementation for updating user
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	// Implementation for deleting user
}
