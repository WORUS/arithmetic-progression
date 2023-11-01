package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	server "github.com/WORUS/arithmetic-progression"
	"github.com/WORUS/arithmetic-progression/internal/app/cache"
	handler "github.com/WORUS/arithmetic-progression/internal/app/handler"
	"github.com/WORUS/arithmetic-progression/internal/app/service"
	"github.com/WORUS/arithmetic-progression/internal/app/task"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

const (
	defaultN = 7
)

func main() {
	ctx := context.Background()

	N := flag.Int("n", defaultN, "max number of goroutines for tasks")
	flag.Parse()

	var queue []*task.Task
	qready := make(chan bool, 1000)
	goroutines := make(chan bool, *N)

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	cache := cache.NewCache()
	service := service.NewService(cache, qready, goroutines, queue)
	handler := handler.NewHandler(service)
	serv := new(server.Server)

	go service.QueueListener(ctx)

	go func() {
		if err := serv.Run(os.Getenv("port"), handler.InitRoutes()); err != nil {
			log.Fatalf("error occurred while running http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := serv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

}
