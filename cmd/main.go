package main

import (
	"log"
	"net/http"

	handler "github.com/WORUS/arithmetic-progression/internal/app/handler"
	"github.com/joho/godotenv"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	handler := new(handler.Handler)

	http.HandleFunc("/task", handler.TaskHandler)

	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
