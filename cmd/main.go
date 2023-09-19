package main

import (
	"github.com/EvgeniyBudaev/golang-zhashkevych-todo-app/internal/app/server"
	"github.com/EvgeniyBudaev/golang-zhashkevych-todo-app/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(server.Server)
	if err := srv.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())
	}
}
