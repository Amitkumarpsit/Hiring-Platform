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
	JobID         string `json:"jobId"` // Expecting JobID as a string
	FullName      string `json:"fullName"`
	Email         string `json:"email"`
	Age           int    `json:"age"`
	Course        string `json:"course"`
	CourseEndDate string `json:"courseEndDate"` // Date as a string
	Address       string `json:"address"`
	PhoneNumber   string `json:"phoneNumber"`
}

func PostApplication(w http.ResponseWriter, r *http.Request) {
	var input ApplicationInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Log incoming request payload
	log.Printf("Received application payload: %+v", input)

	jobID, err := primitive.ObjectIDFromHex(input.JobID)
	if err != nil {
		log.Printf("Error parsing JobID: %v", err)
		http.Error(w, "Invalid JobID", http.StatusBadRequest)
		return
	}
	log.Printf("Parsed JobID: %s", jobID.Hex())

	courseEndDate, err := time.Parse("2006-01-02", input.CourseEndDate)
	if err != nil {
		log.Printf("Error parsing CourseEndDate: %v", err)
		http.Error(w, "Invalid CourseEndDate format", http.StatusBadRequest)
		return
	}

	// Retrieve userID from context (assuming AuthMiddleware sets it)
	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		log.Printf("User not authenticated")
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	// Check if the user has already applied for this job
	hasApplied, err := repository.HasUserAppliedForJob(userID, jobID)
	if err != nil {
		log.Printf("Error checking existing application: %v", err)
		http.Error(w, "Error checking application status", http.StatusInternalServerError)
		return
	}
	if hasApplied {
		log.Printf("User %s has already applied for job %s", userID, jobID.Hex())
		http.Error(w, "You have already applied for this job", http.StatusConflict)
		return
	}

	// Save application to DB
	application := models.Application{
		ID:            primitive.NewObjectID(),
		UserID:        userID,
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

	err = repository.CreateApplication(application)
	if err != nil {
		log.Printf("Error creating application: %v", err)
		http.Error(w, "Failed to submit application", http.StatusInternalServerError)
		return
	}

	// Log success and respond
	log.Printf("Application created successfully for JobID: %s", jobID.Hex())
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "Application submitted successfully",
	})
}
