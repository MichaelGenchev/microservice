package main

import (

	"log"


	"github.com/MichaelGenchev/microservice/internal/service/api"
	"github.com/MichaelGenchev/microservice/internal/service/config"
	"github.com/go-chi/chi/v5"
)


func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new router
	router := chi.NewMux()

	// Set up the API routes and handlers
	api.SetupRoutesAndRun(router, cfg.Server.ListenAddr, nil)

}
