package user_handler

import (
	"encoding/json"
	"net/http"

	user_model "github.com/sri2103/domain_DD_todo/internal/app/user/model"
	user_service "github.com/sri2103/domain_DD_todo/internal/app/user/service"
)

type UserHandler struct {
	userService user_service.UserService
}

func NewUserHandler(userService user_service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// create
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	user := &user_model.User{}

	// convert request body to json format and store in the object
	err := user.ToJSON(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	// call service layer method for creating a new user with data received as parameter
	err = h.userService.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}

// find by ID
func (h *UserHandler) FindById(w http.ResponseWriter, r *http.Request) {

}

// Update

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {}

// delete
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {}

// Get all
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	usersList, err := h.userService.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	jsonData, err := json.MarshalIndent(&usersList, "", " ") // indenting is important here!
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
