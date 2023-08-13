package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/sri2103/domain_DD_todo/internal/app/common"
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
	r := mux.NewRouter()
	r.Use(common.LoggingMiddleWare)
	r.Use(mux.CORSMethodMiddleware(r))
	r.HandleFunc("/todos", todoHandler.CreateTodo)
	r.HandleFunc("/todos/{id}", todoHandler.GetTodoById)
	r.HandleFunc("/todo",todoHandler.GetAllTodos)


	log.Print("todo DDD")
	srv := &http.Server{
		Handler: r,
		Addr:":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout: 15 * time.Second,
	}

	go func(){
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c,os.Interrupt)

	sig:=<-c
	log.Println("Received terminate, graceful shutdown", sig)

	ctx,cancel := context.WithTimeout(context.Background(),30*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}