package todo_repository

import todo_model "github.com/sri2103/domain_DD_todo/internal/app/todo/model"

type TodoRepository interface {
	Save(todo *todo_model.Todo) error
	FindById(id uint)(*todo_model.Todo,error)
	FindAll()([]*todo_model.Todo,error)
}