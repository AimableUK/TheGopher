package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Author struct {
	ID   primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name string             `json:"name" bson:"name"`
	Bio  string             `json:"bio" bson:"bio"`
}
