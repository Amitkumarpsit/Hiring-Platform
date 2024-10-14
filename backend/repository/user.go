package repository

import (
	"backend/config"
	"backend/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUser(user models.User) error {
	_, err := config.DB.Collection("users").InsertOne(context.Background(), user)
	return err
}

func GetUserByUsername(username string) (models.User, error) {
	var user models.User
	err := config.DB.Collection("users").FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return user, err
	}
	return user, err
}
