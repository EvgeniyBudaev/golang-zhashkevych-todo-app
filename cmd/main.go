package main

import (
	"github.com/EvgeniyBudaev/golang-zhashkevych-todo-app/internal/app/server"
	"github.com/EvgeniyBudaev/golang-zhashkevych-todo-app/pkg/handler"
	"github.com/EvgeniyBudaev/golang-zhashkevych-todo-app/pkg/repository"
	"github.com/EvgeniyBudaev/golang-zhashkevych-todo-app/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())
	}
}
