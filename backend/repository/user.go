package repository

import (
	"backend/config"
	"backend/models"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

func CreateUser(user models.User) error {
	_, err := config.DB.Collection("users").InsertOne(context.Background(), user)
	return err
}

func GetUserByLoginID(loginID string) (models.User, error) {
	var user models.User
	filter := bson.M{"$or": []bson.M{
		{"phoneNumber": loginID},
		{"email": loginID},
	}}
	err := config.DB.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	return user, err
}
