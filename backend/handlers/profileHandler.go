// backend/handlers/profileHandler.go
package handlers

import (
	"backend/models"
	"backend/repository"
	"encoding/json"
	"log"
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(candidates)
}

// Post a new candidate
func PostCandidate(w http.ResponseWriter, r *http.Request) {
	var candidate models.Candidate
	err := json.NewDecoder(r.Body).Decode(&candidate)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	candidate.ID = primitive.NewObjectID()
	err = repository.CreateCandidate(candidate)
	if err != nil {
		http.Error(w, "Failed to create candidate", http.StatusInternalServerError)
		return
	}

	// Send acknowledgment to the user
	response := map[string]string{
		"message": "Candidate created successfully",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

// Fetch the user's profile
func GetProfile(w http.ResponseWriter, r *http.Request) {
	// Ensure user_id is in the context, otherwise handle error
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "User ID not found", http.StatusUnauthorized)
		return
	}

	profile, err := repository.GetUserProfile(userID)
	if err != nil {
		log.Println("Error fetching profile:", err)
		http.Error(w, "Failed to fetch profile", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(profile)
}

// Update the user's profile
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var profile models.User
	err := json.NewDecoder(r.Body).Decode(&profile)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = repository.UpdateUserProfile(profile)
	if err != nil {
		http.Error(w, "Failed to update profile", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Profile updated successfully"})
}
