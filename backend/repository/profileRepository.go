// backend/repository/profileRepository.go
package repository

import (
	"backend/config"
	"backend/models"
	"context"
)

// CreateCandidate inserts a new candidate into the database
func CreateCandidate(candidate models.Candidate) error {
	_, err := config.DB.Collection("candidates").InsertOne(context.Background(), candidate)
	return err
}

// GetAllCandidates retrieves all candidates from the database
func GetAllCandidates() ([]models.Candidate, error) {
	cursor, err := config.DB.Collection("candidates").Find(context.Background(), map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var candidates []models.Candidate
	err = cursor.All(context.Background(), &candidates)
	return candidates, err
}
