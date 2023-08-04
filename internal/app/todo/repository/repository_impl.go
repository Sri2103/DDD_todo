package todo_repository

import todo_model "github.com/sri2103/domain_DD_todo/internal/app/todo/model"

type TodoRepositoryImpl struct {
	todos []*todo_model.Todo
}

func NewTodoRepositoryImpl()*TodoRepositoryImpl{
	return &TodoRepositoryImpl{todos: []*todo_model.Todo{}}
}

func (r *TodoRepositoryImpl) Save(todo *todo_model.Todo) error {
	r.todos = append(r.todos, todo)
	return nil
}

func (r *TodoRepositoryImpl) FindById(id uint)(*todo_model.Todo,error) {
	for _,todo := range r.todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return nil, nil
}

func (r *TodoRepositoryImpl) FindAll()([]*todo_model.Todo,error){
	return r.todos,nil
}