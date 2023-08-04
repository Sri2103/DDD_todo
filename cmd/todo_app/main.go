package main

import (
	"log"
	"net/http"

	todo_repository "github.com/sri2103/domain_DD_todo/internal/app/todo/repository"
	todo_service "github.com/sri2103/domain_DD_todo/internal/app/todo/service"
	"github.com/sri2103/domain_DD_todo/internal/db"
	todo_handler "github.com/sri2103/domain_DD_todo/internal/delivery/http"
)

func main() {
	database, err := db.ConnectToDB()
    if err != nil {
        log.Fatal("Failed to connect to the database:", err)
    }

	if err := db.MigrateModels(database); err != nil {
        log.Fatal("Failed to perform auto-migration:", err)
    }
	todoPgImpl := todo_repository.NewTodoPostgresImpl(database)
	// toRepo := todo_repository.NewTodoRepositoryImpl()
	todoService:=todo_service.NewTodoServiceImpl(todoPgImpl)
	todoHandler := todo_handler.NewTodoHandler(todoService)
	http.HandleFunc("/todos", todoHandler.CreateTodo)
	http.HandleFunc("/todos/{id}", todoHandler.GetTodoById)
	http.HandleFunc("/todo",todoHandler.GetAllTodos)
	log.Print("todo DDD")
	log.Fatal(http.ListenAndServe(":8080", nil))
}