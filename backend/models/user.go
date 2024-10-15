package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID                   primitive.ObjectID `bson:"_id,omitempty"`
	FullName             string             `bson:"fullName"`
	PhoneNumber          string             `bson:"phoneNumber"`
	Email                string             `bson:"email"`
	Address              string             `bson:"address"`
	Skills               []string           `bson:"skills"`
	Course               string             `bson:"course"`
	Specialization       string             `bson:"specialization"`
	Password             string             `bson:"password"`
	ResetToken           string             `bson:"resetToken,omitempty"`
	ResetTokenExpiration time.Time          `bson:"resetTokenExpiration,omitempty"`
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
