// backend/models/candidateModel.go
package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Candidate struct defines a candidate document in MongoDB
type Candidate struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Name   string             `bson:"name"`
	Skills []string           `bson:"skills"`
	Resume string             `bson:"resume"`
}
