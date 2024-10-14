// backend/repository/profileRepository.go
package repository

import (
	"backend/config"
	"backend/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CreateCandidate inserts a new candidate into the database
func CreateCandidate(candidate models.Candidate) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := config.DB.Collection("candidates").InsertOne(ctx, candidate)
	return err
}

// GetAllCandidates retrieves all candidates from the database
func GetAllCandidates() ([]models.Candidate, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := config.DB.Collection("candidates").Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	var candidates []models.Candidate
	if err = cursor.All(ctx, &candidates); err != nil {
		return nil, err
	}
	return candidates, nil
}

// GetUserProfile fetches a user profile by ID from the database
func GetUserProfile(userID string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	var user models.User
	err = config.DB.Collection("users").FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// UpdateUserProfile updates an existing user's profile in the database
func UpdateUserProfile(profile models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := config.DB.Collection("users").UpdateOne(
		ctx,
		bson.M{"_id": profile.ID},
		bson.M{"$set": bson.M{
			"name":  profile.FullName,
			"email": profile.Email,
			// Add other fields you wish to update
		}},
	)
	return err
}
