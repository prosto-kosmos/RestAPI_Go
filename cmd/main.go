package main

import (
	todo "github.com/prosto-kosmos/RestAPI_Go"
	"github.com/prosto-kosmos/RestAPI_Go/pkg/handler"
	"github.com/prosto-kosmos/RestAPI_Go/pkg/repository"
	"github.com/prosto-kosmos/RestAPI_Go/pkg/service"

	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRouters()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
