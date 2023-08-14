package user_handler

import (
	"net/http"

	"github.com/gorilla/mux"
	user_repository "github.com/sri2103/domain_DD_todo/internal/app/user/repository"
	user_service "github.com/sri2103/domain_DD_todo/internal/app/user/service"
)

func SetUpUserRoutes(userRepo user_repository.UserRepository, r *mux.Router) {
	userService := user_service.NewUserServiceImpl(userRepo)
	handlers := NewUserHandler(userService)
	r.HandleFunc("/users/create", handlers.CreateUser).Methods(http.MethodPost)
	r.HandleFunc("/users/all", handlers.GetAllUsers).Methods(http.MethodGet)
	r.HandleFunc("/users/{id}", handlers.FindById).Methods(http.MethodGet)
	r.HandleFunc("/users/update",handlers.UpdateUser).Methods(http.MethodPut)
}
