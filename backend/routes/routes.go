package routes

import (
	"backend/handlers"
	"backend/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Auth routes
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	// Job routes
	r.HandleFunc("/jobs", handlers.GetJobs).Methods("GET")
	r.HandleFunc("/jobs/new", middleware.AuthMiddleware(handlers.PostJob)).Methods("POST")
	r.HandleFunc("/jobs/category/{category}", handlers.GetJobsByCategory).Methods("GET")

	// Candidate routes
	r.HandleFunc("/candidates", handlers.GetCandidates).Methods("GET")
	r.HandleFunc("/candidates/new", handlers.PostCandidate).Methods("POST")

	// Application route
	r.HandleFunc("/applications", handlers.PostApplication).Methods("POST")

	return r
}
