package routes

import (
	"backend/handlers"
	"backend/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()

	// Public routes
	r.HandleFunc("/register", handlers.Register).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/forgot-password", handlers.ForgotPassword).Methods("POST")
	r.HandleFunc("/reset-password", handlers.ResetPassword).Methods("POST")

	// Protected routes
	r.HandleFunc("/jobs", middleware.AuthMiddleware(handlers.GetJobs)).Methods("GET")
	r.HandleFunc("/jobs/new", middleware.AuthMiddleware(handlers.PostJob)).Methods("POST")
	r.HandleFunc("/jobs/category/{category}", middleware.AuthMiddleware(handlers.GetJobsByCategory)).Methods("GET")
	r.HandleFunc("/candidates", middleware.AuthMiddleware(handlers.GetCandidates)).Methods("GET")
	r.HandleFunc("/candidates/new", middleware.AuthMiddleware(handlers.PostCandidate)).Methods("POST")
	r.HandleFunc("/applications", middleware.AuthMiddleware(handlers.PostApplication)).Methods("POST")
	r.HandleFunc("/profile", middleware.AuthMiddleware(handlers.GetProfile)).Methods("GET")
	r.HandleFunc("/profile", middleware.AuthMiddleware(handlers.UpdateProfile)).Methods("PUT")

	return r
}
