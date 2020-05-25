package main

import (
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"github.com/willbarkoff/crazyhairdontcare/api"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := httprouter.New()

	api.Initalize(router)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:1234"},
		AllowCredentials: true,
	}).Handler(router)

	log.Fatal(http.ListenAndServe(":9696", cors))
}
