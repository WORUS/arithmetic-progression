package main

import (
	"context"
	"flag"
	"log"
	"os"

	server "github.com/WORUS/arithmetic-progression"
	"github.com/WORUS/arithmetic-progression/internal/app/cache"
	handler "github.com/WORUS/arithmetic-progression/internal/app/handler"
	"github.com/WORUS/arithmetic-progression/internal/app/service"
	"github.com/WORUS/arithmetic-progression/internal/app/task"
	"github.com/joho/godotenv"
)

var defaultN = 7

func main() {
	ctx := context.Background()
	var N int
	flag.Int("N", defaultN, "number of goroutines")
	queue := make(chan *task.Task, 7)
	goroutines := make(chan bool, N)

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	cache := cache.NewCache()
	service := service.NewService(cache, queue, goroutines)
	handler := handler.Newhandler(service)
	serv := new(server.Server)

	go service.QueueListener(ctx)

	func() {
		if err := serv.Run(os.Getenv("port"), handler.InitRoutes()); err != nil {
			log.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()

}
