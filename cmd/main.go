package main

import (
	"github.com/joho/godotenv"
	"log"
	"todo-list/internal/http"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	route := http.NewRoute()
	route.RouteRun()

}
