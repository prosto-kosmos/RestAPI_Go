package main

import (
	todo "github.com/prosto-kosmos/RestAPI_Go"
	"github.com/prosto-kosmos/RestAPI_Go/pkg/handler"

	"log"
)

func main() {
	handlers := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.InitRouters()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
