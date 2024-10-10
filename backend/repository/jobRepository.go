// backend/repository/jobRepository.go
package repository

import (
	"backend/config"
	"backend/models"
	"context"
)

// CreateJob inserts a new job into the database
func CreateJob(job models.Job) error {
	_, err := config.DB.Collection("jobs").InsertOne(context.Background(), job)
	return err
}

// GetAllJobs retrieves all jobs from the database
func GetAllJobs() ([]models.Job, error) {
	cursor, err := config.DB.Collection("jobs").Find(context.Background(), map[string]interface{}{})
	if err != nil {
		return nil, err
	}

	var jobs []models.Job
	err = cursor.All(context.Background(), &jobs)
	return jobs, err
}

func CreateApplication(application models.Application) error {
	_, err := config.DB.Collection("applications").InsertOne(context.Background(), application)
	return err
}
