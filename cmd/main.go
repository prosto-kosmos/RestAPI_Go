package main

import (
	todo "github.com/prosto-kosmos/RestAPI_Go"
	"github.com/prosto-kosmos/RestAPI_Go/pkg/handler"
	"github.com/prosto-kosmos/RestAPI_Go/pkg/repository"
	"github.com/prosto-kosmos/RestAPI_Go/pkg/service"
	"github.com/spf13/viper"

	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing congigs: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRouters()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
