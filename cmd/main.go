package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/otumian-empire/go-weather-api-consumer/api"
)

func Recover() {
	if err := recover(); err != nil {
		log.Printf("an error occurred: %v", err)
	}
}

func main() {
	log.Println("Consuming weather API")

	// Catch unhandled errors
	defer Recover()

	// Load env variables
	ENV_CONST, err := godotenv.Read()

	if err != nil {
		log.Fatal(err)
		log.Fatal("an error occurred reading server credentials")
	}

	// get the port and host
	ADDRESS := fmt.Sprintf("%v:%v", ENV_CONST["HOST"], ENV_CONST["PORT"])

	// define the endpoint for making api request
	http.HandleFunc("/", api.QueryHandler)
	http.HandleFunc("/format", api.FormatQueryHandler)

	// start server
	log.Printf("Running server on: %v", ADDRESS)
	log.Fatal(http.ListenAndServe(ADDRESS, nil))
}
