// backend/handlers/jobHandler.go
package handlers

import (
	"backend/models"
	"backend/repository"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get all jobs
func GetJobs(w http.ResponseWriter, r *http.Request) {
	jobs, err := repository.GetAllJobs()
	if err != nil {
		http.Error(w, "Unable to fetch jobs", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(jobs)
}

// Post a new job
func PostJob(w http.ResponseWriter, r *http.Request) {
	var job models.Job
	err := json.NewDecoder(r.Body).Decode(&job)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Validate the job category
	switch models.JobCategory(job.Category) {
	case models.ITJob, models.SalesJob, models.MarketingJob, models.HRJob, models.EngineeringJob, models.OtherJob:
		// Category is valid
	default:
		http.Error(w, "Invalid job category", http.StatusBadRequest)
		return
	}

	job.ID = primitive.NewObjectID()

	// Get the user ID from the context (set by the AuthMiddleware)
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}
	job.CreatedBy = userID

	err = repository.CreateJob(job)
	if err != nil {
		http.Error(w, "Failed to post job", http.StatusInternalServerError)
		return
	}

	response := map[string]string{
		"message": "Job submitted successfully",
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// GetJobsByCategory handles retrieving jobs by category
func GetJobsByCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	category := vars["category"]

	jobs, err := repository.GetJobsByCategory(category)
	if err != nil {
		http.Error(w, "Failed to get jobs", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jobs)
}

func ApplyForJob(w http.ResponseWriter, r *http.Request) {
	var application models.Application

	// Decode the request body into application model
	if err := json.NewDecoder(r.Body).Decode(&application); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Retrieve userID from context (assuming AuthMiddleware sets it)
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		log.Printf("User not authenticated")
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	// Convert JobID from string to ObjectID (ensure frontend sends a valid string)
	jobID, err := primitive.ObjectIDFromHex(application.JobID.Hex())
	if err != nil {
		log.Printf("Invalid JobID format: %v", err)
		http.Error(w, "Invalid JobID", http.StatusBadRequest)
		return
	}
	application.JobID = jobID

	// Check if the user has already applied for this job
	applied, err := repository.HasUserAppliedForJob(userID, application.JobID)
	if err != nil {
		log.Printf("Error checking if user has applied: %v", err)
		http.Error(w, "Error checking application status", http.StatusInternalServerError)
		return
	}
	if applied {
		log.Printf("User has already applied for this job")
		http.Error(w, "Already applied for this job post", http.StatusConflict)
		return
	}

	// Set additional fields
	application.UserID = userID
	application.AppliedAt = time.Now()

	// Save application to the database
	if err := repository.CreateApplication(application); err != nil {
		log.Printf("Error creating application: %v", err)
		http.Error(w, "Error creating application", http.StatusInternalServerError)
		return
	}

	log.Printf("Application submitted successfully for user %s", userID)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Application submitted successfully"})
}
