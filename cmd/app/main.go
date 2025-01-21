package main

import (
	"log"

	"github.com/kibo13/todo-app"
	"github.com/kibo13/todo-app/internal/handler"
	"github.com/kibo13/todo-app/internal/repository"
	"github.com/kibo13/todo-app/internal/service"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
