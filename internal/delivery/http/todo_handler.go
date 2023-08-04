package todo_handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	todo_service "github.com/sri2103/domain_DD_todo/internal/app/todo/service"
)

type TodoHandler struct {
	todoService todo_service.TodoService
}

func NewTodoHandler(todoService todo_service.TodoService)*TodoHandler{
	return &TodoHandler{todoService}
}

func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w,"Invalid request method",http.StatusMethodNotAllowed)
		return
	}
	title := r.FormValue("title")

	if title == "" {
		http.Error(w,"title cannot be empty",http.StatusBadRequest)
		return
	}

	todo,err := h.todoService.CreateTodo(title)
	if err != nil {
		http.Error(w, err.Error(),http.StatusInternalServerError)
	}
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) GetTodoById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed",http.StatusMethodNotAllowed)
		return
	}
	idStr := r.FormValue("id")
	id,err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w,"Invalid Todo ID",http.StatusBadRequest)
		return
	}
	todo,err := h.todoService.GetTodoById(uint(id))
	if err != nil {
		http.Error(w,err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-type","application/json")
	json.NewEncoder(w).Encode(todo)
}

func (h *TodoHandler) GetAllTodos(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w,"Method is not valid",http.StatusMethodNotAllowed)
		return
	}
	todos,err := h.todoService.GetAllTodos()
	if err != nil {
		http.Error(w,err.Error(),http.StatusInternalServerError)
		return 
	}

	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(todos)
	
}