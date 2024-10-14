package repository

import (
	"backend/config"
	"backend/models"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	if err != nil {
		log.Printf("Error inserting application into database: %v", err)
		return err
	}
	return nil
}

func HasUserAppliedForJob(userID string, jobID primitive.ObjectID) (bool, error) {
	count, err := config.DB.Collection("applications").CountDocuments(
		context.Background(),
		bson.M{"userId": userID, "jobId": jobID},
	)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
