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
	todo_handler "github.com/sri2103/domain_DD_todo/internal/app/todo/handlers"
	todo_repository "github.com/sri2103/domain_DD_todo/internal/app/todo/repository"
	user_handler "github.com/sri2103/domain_DD_todo/internal/app/user/handlers"
	user_repository "github.com/sri2103/domain_DD_todo/internal/app/user/repository"
	"github.com/sri2103/domain_DD_todo/internal/db"
)

func SetUp() *mux.Router {
	database, err := db.ConnectToDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	if err := db.MigrateModels(database); err != nil {
		log.Fatal("Failed to perform auto-migration:", err)
	}

	todoPgImpl := todo_repository.NewTodoPostgresImpl(database)
	userPgImpl := user_repository.NewUserPostgresImpl(database)

	r := mux.NewRouter()

	r.Use(common.LoggingMiddleWare)
	r.Use(mux.CORSMethodMiddleware(r))

	todo_handler.SetUpTodoRoutes(todoPgImpl, r)
	user_handler.SetUpUserRoutes(userPgImpl,r)

	return r
}
func main() {
	r := SetUp()
	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)

	sig := <-c
	log.Println("Received terminate, graceful shutdown", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	srv.Shutdown(ctx)
}
