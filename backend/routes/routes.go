// backend/routes/routes.go
package routes

import (
	"backend/handlers"

	"github.com/gorilla/mux"
)

// SetupRoutes initializes the routes
func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Job routes
	r.HandleFunc("/jobs", handlers.GetJobs).Methods("GET")
	r.HandleFunc("/jobs/new", handlers.PostJob).Methods("POST")

	// Candidate routes
	r.HandleFunc("/candidates", handlers.GetCandidates).Methods("GET")
	r.HandleFunc("/candidates/new", handlers.PostCandidate).Methods("POST")

	// New application route
	r.HandleFunc("/applications", handlers.PostApplication).Methods("POST")

	return r
}
