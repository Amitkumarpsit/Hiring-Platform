// backend/handlers/jobHandler.go
package handlers

import (
	"backend/models"
	"backend/repository"
	"encoding/json"
	"net/http"

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
	json.NewDecoder(r.Body).Decode(&job)
	job.ID = primitive.NewObjectID()
	err := repository.CreateJob(job)
	if err != nil {
		http.Error(w, "Failed to post job", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
