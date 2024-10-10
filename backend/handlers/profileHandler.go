// backend/handlers/profileHandler.go
package handlers

import (
	"backend/models"
	"backend/repository"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Get all candidates
func GetCandidates(w http.ResponseWriter, r *http.Request) {
	candidates, err := repository.GetAllCandidates()
	if err != nil {
		http.Error(w, "Unable to fetch candidates", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(candidates)
}

// Post a new candidate
func PostCandidate(w http.ResponseWriter, r *http.Request) {
	var candidate models.Candidate
	json.NewDecoder(r.Body).Decode(&candidate)
	candidate.ID = primitive.NewObjectID()
	err := repository.CreateCandidate(candidate)
	if err != nil {
		http.Error(w, "Failed to create candidate", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
