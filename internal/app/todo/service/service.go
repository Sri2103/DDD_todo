package todo_service

import todo_model "github.com/sri2103/domain_DD_todo/internal/app/todo/model"

type TodoService interface {
	CreateTodo(title string) (*todo_model.Todo,error)
	GetTodoById(id uint)(*todo_model.Todo,error)
	GetAllTodos()([]*todo_model.Todo,error)
}