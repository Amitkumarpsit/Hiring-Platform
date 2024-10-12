// backend/handlers/jobHandler.go
package handlers

import (
	"backend/models"
	"backend/repository"
	"encoding/json"
	"net/http"

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
	err = repository.CreateJob(job)
	if err != nil {
		http.Error(w, "Failed to post job", http.StatusInternalServerError)
		return
	}
	// Send acknowledgment to the user
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
