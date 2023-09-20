package main

import (
	"github.com/EvgeniyBudaev/golang-zhashkevych-todo-app/internal/app/server"
	"github.com/EvgeniyBudaev/golang-zhashkevych-todo-app/pkg/handler"
	"github.com/EvgeniyBudaev/golang-zhashkevych-todo-app/pkg/repository"
	"github.com/EvgeniyBudaev/golang-zhashkevych-todo-app/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.InitRoutes()); err != nil {
		log.Fatal("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
