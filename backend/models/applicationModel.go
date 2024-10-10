package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Application struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	JobID         primitive.ObjectID `bson:"jobId"`
	FullName      string             `bson:"fullName"`
	Email         string             `bson:"email"`
	Age           int                `bson:"age"`
	Course        string             `bson:"course"`
	CourseEndDate time.Time          `bson:"courseEndDate"`
	Address       string             `bson:"address"`
	PhoneNumber   string             `bson:"phoneNumber"`
	AppliedAt     time.Time          `bson:"appliedAt"`
}
