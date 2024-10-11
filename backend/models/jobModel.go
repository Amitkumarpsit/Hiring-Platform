package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Job struct defines a job document in MongoDB
type Job struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Title            string             `bson:"title"`
	Company          string             `bson:"company"`
	Responsibilities string             `bson:"responsibilities"`
	Qualifications   string             `bson:"qualifications"`
	Location         string             `bson:"location"`
	Category         string             `bson:"category"`
}

// JobCategory defines the available job categories
type JobCategory string

const (
	ITJob          JobCategory = "IT"
	SalesJob       JobCategory = "Sales"
	MarketingJob   JobCategory = "Marketing"
	HRJob          JobCategory = "HR"
	EngineeringJob JobCategory = "Engineering"
	OtherJob       JobCategory = "Other"
)
