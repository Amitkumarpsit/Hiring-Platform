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
		AllowedOrigins: []string{"http://localhost:3000"}, // Allow requests from your React app
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		// Enable Debugging for testing, consider disabling in production
		Debug: true,
	})

	// Wrap the router with the CORS handler
	handler := c.Handler(router)

	log.Println("Server started on :8000")
	// Start server on port 8000 with CORS-enabled handler
	log.Fatal(http.ListenAndServe(":8000", handler))
}
