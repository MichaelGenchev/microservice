package main

import (
	"log"

	"github.com/MichaelGenchev/microservice/service/config"
)


func main() {
	// Load configuration
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Create a new router
	// router := gin.New()

	// Set up the API routes and handlers
	api.SetupRoutes(router)

	// Create a new HTTP server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Server.Port),
		Handler: router,
	}

	// Start the server in a goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	// Wait for a shutdown signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	// Shut down the server gracefully
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
}
