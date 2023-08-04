package todo_service

import (
	"log"

	todo_model "github.com/sri2103/domain_DD_todo/internal/app/todo/model"
	todo_repository "github.com/sri2103/domain_DD_todo/internal/app/todo/repository"
)

type TodoServiceImpl struct {
	todoRepo todo_repository.TodoRepository
}

func NewTodoServiceImpl(todoRepo todo_repository.TodoRepository) *TodoServiceImpl {
	return &TodoServiceImpl{todoRepo: todoRepo}
}

func (s *TodoServiceImpl) CreateTodo(title string)(*todo_model.Todo,error){
	t,_ := s.todoRepo.FindAll()
	todo := &todo_model.Todo{
		ID:uint(len(t) +1), 
		Title: title,
		Completed: false,
	}
	if err := s.todoRepo.Save(todo); err != nil {
		return nil, err
	}
	log.Println(todo)
	return todo, nil

}

func (s *TodoServiceImpl) GetTodoById(id uint) (*todo_model.Todo, error) {
	return s.todoRepo.FindById(id)
}

func (s *TodoServiceImpl) GetAllTodos()([]*todo_model.Todo,error){
	return s.todoRepo.FindAll()
}