package main

import (
	"log"
	"os"

	server "github.com/WORUS/arithmetic-progression"
	handler "github.com/WORUS/arithmetic-progression/internal/app/handler"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	handler := new(handler.Handler)

	serv := new(server.Server)

	func() {
		if err := serv.Run(os.Getenv("port"), handler.InitRoutes()); err != nil {
			log.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()

}
