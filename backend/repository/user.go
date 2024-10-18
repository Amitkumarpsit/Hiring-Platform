package repository

import (
	"backend/config"
	"backend/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
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

func ValidateAuthToken(authToken string) (primitive.ObjectID, error) {
	var user models.User
	filter := bson.M{
		"authToken":           authToken,
		"authTokenExpiration": bson.M{"$gt": time.Now()},
	}
	err := config.DB.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return user.ID, nil
}

func VerifyUser(userID primitive.ObjectID) error {
	update := bson.M{
		"$set": bson.M{
			"isVerified":          true,
			"authToken":           nil,
			"authTokenExpiration": nil,
		},
	}
	_, err := config.DB.Collection("users").UpdateOne(context.Background(), bson.M{"_id": userID}, update)
	return err
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	filter := bson.M{"email": email}
	err := config.DB.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	return user, err
}

func SaveResetToken(userID primitive.ObjectID, resetToken string, expirationTime time.Time) error {
	update := bson.M{
		"$set": bson.M{
			"resetToken":           resetToken,
			"resetTokenExpiration": expirationTime,
		},
	}
	_, err := config.DB.Collection("users").UpdateOne(context.Background(), bson.M{"_id": userID}, update)
	return err
}

func ValidateResetToken(resetToken string) (primitive.ObjectID, error) {
	var user models.User
	filter := bson.M{
		"resetToken":           resetToken,
		"resetTokenExpiration": bson.M{"$gt": time.Now()},
	}
	err := config.DB.Collection("users").FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return user.ID, nil
}

func UpdatePassword(userID primitive.ObjectID, newPassword string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": bson.M{
			"password":             string(hashedPassword),
			"resetToken":           nil,
			"resetTokenExpiration": nil,
		},
	}
	_, err = config.DB.Collection("users").UpdateOne(context.Background(), bson.M{"_id": userID}, update)
	return err
}
