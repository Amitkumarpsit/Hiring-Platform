package handlers

import (
	"backend/models"
	"backend/repository"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ApplicationInput struct {
	JobID         string `json:"jobId"`
	FullName      string `json:"fullName"`
	Email         string `json:"email"`
	Age           int    `json:"age"`
	Course        string `json:"course"`
	CourseEndDate string `json:"courseEndDate"`
	Address       string `json:"address"`
	PhoneNumber   string `json:"phoneNumber"`
}

func PostApplication(w http.ResponseWriter, r *http.Request) {
	var input ApplicationInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		log.Printf("Error decoding application JSON: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Parse the JobID
	jobID, err := primitive.ObjectIDFromHex(input.JobID)
	if err != nil {
		log.Printf("Error parsing JobID: %v", err)
		http.Error(w, "Invalid JobID", http.StatusBadRequest)
		return
	}

	// Parse the CourseEndDate
	courseEndDate, err := time.Parse("2006-01-02", input.CourseEndDate)
	if err != nil {
		log.Printf("Error parsing CourseEndDate: %v", err)
		http.Error(w, "Invalid CourseEndDate format", http.StatusBadRequest)
		return
	}

	application := models.Application{
		ID:            primitive.NewObjectID(),
		JobID:         jobID,
		FullName:      input.FullName,
		Email:         input.Email,
		Age:           input.Age,
		Course:        input.Course,
		CourseEndDate: courseEndDate,
		Address:       input.Address,
		PhoneNumber:   input.PhoneNumber,
		AppliedAt:     time.Now(),
	}

	// Save the application to the database
	err = repository.CreateApplication(application)
	if err != nil {
		log.Printf("Error creating application: %v", err)
		http.Error(w, "Failed to submit application", http.StatusInternalServerError)
		return
	}

	// Send acknowledgment to the user
	response := map[string]string{
		"message": "Application submitted successfully",
	}

	// Respond to the user with a success message
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
