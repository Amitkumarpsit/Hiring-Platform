// backend/models/jobModel.go
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
}
