package todo_handler

import (
	"net/http"

	"github.com/gorilla/mux"
	todo_repository "github.com/sri2103/domain_DD_todo/internal/app/todo/repository"
	todo_service "github.com/sri2103/domain_DD_todo/internal/app/todo/service"
)

func SetUpTodoRoutes(todoRepo todo_repository.TodoRepository, r *mux.Router) {
	todoService := todo_service.NewTodoServiceImpl(todoRepo)
	handlers := NewTodoHandler(todoService)
	r.HandleFunc("/todos", handlers.CreateTodo).Methods(http.MethodPost)
	r.HandleFunc("/todos/{id}", handlers.GetTodoById).Methods(http.MethodGet)
	r.HandleFunc("/todo", handlers.GetAllTodos).Methods(http.MethodGet)
}
