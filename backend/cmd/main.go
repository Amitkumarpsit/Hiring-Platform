package main

import (
	"backend/config"
	"backend/routes"
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	// Initialize database connection
	config.ConnectDB()

	// Set up routes
	router := routes.SetupRoutes()

	// Create a new CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
		Debug:            true,
	})

	// Wrap the router with the CORS handler
	handler := c.Handler(router)

	log.Println("Server started on :8000")
	// Start server on port 8000 with CORS-enabled handler
	log.Fatal(http.ListenAndServe(":8000", handler))
}
