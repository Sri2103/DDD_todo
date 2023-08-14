package todo_repository

import (
	"log"

	todo_model "github.com/sri2103/domain_DD_todo/internal/app/todo/model"
	"gorm.io/gorm"
)

type Pg_Todo struct {
	*gorm.Model
	ID        uint `gorm:"primarykey"`
	Title     string
	Completed bool
}

type TodoPostgresImpl struct {
	db *gorm.DB
}

func NewTodoPostgresImpl(db *gorm.DB) *TodoPostgresImpl {
	return &TodoPostgresImpl{db: db}
}

func (r *TodoPostgresImpl) Save(todo *todo_model.Todo) error {
	log.Println("Creating a todo")
	pT := &Pg_Todo{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: todo.Completed,
	}
	err := r.db.Create(&pT).Error
	return err
}

func (r *TodoPostgresImpl) FindById(id uint) (*todo_model.Todo, error) {
	var pT = new(Pg_Todo)
	if err := r.db.Where("id=?", id).First(&pT).Error; err != nil {
		return nil, err
	}
	t := &todo_model.Todo{
		ID:        pT.ID,
		Title:     pT.Title,
		Completed: pT.Completed,
	}
	return t, nil
}

func (r *TodoPostgresImpl) FindAll() ([]*todo_model.Todo, error) {
	var Pg_todos []*Pg_Todo
	if err := r.db.Find(&Pg_todos).Error; err != nil {
		return nil, err
	}
	var todos []*todo_model.Todo

	for _, t := range Pg_todos {
		todos = append(todos, &todo_model.Todo{
			ID:        t.ID,
			Title:     t.Title,
			Completed: t.Completed,
		})
	}
	return todos, nil
}
