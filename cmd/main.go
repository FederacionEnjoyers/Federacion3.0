package main

import (
	federacion "github.com/Serelllka/Federacion"
	"github.com/Serelllka/Federacion/pkg/handler"
	"github.com/Serelllka/Federacion/pkg/repository"
	"github.com/Serelllka/Federacion/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(federacion.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
