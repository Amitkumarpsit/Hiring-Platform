package repository

import (
	"backend/config"
	"backend/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// GetJobsByCategory retrieves jobs from the database by category
func GetJobsByCategory(category string) ([]models.Job, error) {
	var jobs []models.Job
	cursor, err := config.DB.Collection("jobs").Find(context.Background(), bson.M{"category": category})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.Background(), &jobs)
	return jobs, err
}
