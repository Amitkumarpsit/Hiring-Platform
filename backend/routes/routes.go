package routes

import (
	"backend/handlers"
	"backend/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	api := mux.NewRouter()

	// Public routes
	api.HandleFunc("/register", handlers.Register).Methods("POST")
	api.HandleFunc("/login", handlers.Login).Methods("POST")
	api.HandleFunc("/forgot-password", handlers.ForgotPassword).Methods("POST")
	api.HandleFunc("/reset-password", handlers.ResetPassword).Methods("POST")

	// Protected routes
	api.HandleFunc("/jobs", middleware.AuthMiddleware(handlers.GetJobs)).Methods("GET")
	api.HandleFunc("/jobs/new", middleware.AuthMiddleware(handlers.PostJob)).Methods("POST")
	api.HandleFunc("/jobs/category/{category}", middleware.AuthMiddleware(handlers.GetJobsByCategory)).Methods("GET")
	api.HandleFunc("/candidates", middleware.AuthMiddleware(handlers.GetCandidates)).Methods("GET")
	api.HandleFunc("/candidates/new", middleware.AuthMiddleware(handlers.PostCandidate)).Methods("POST")
	api.HandleFunc("/applications", middleware.AuthMiddleware(handlers.PostApplication)).Methods("POST")
	api.HandleFunc("/profile", middleware.AuthMiddleware(handlers.GetProfile)).Methods("GET")
	api.HandleFunc("/profile", middleware.AuthMiddleware(handlers.UpdateProfile)).Methods("PUT")

	return api
	//added
}
