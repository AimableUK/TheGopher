package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Book struct {
	ID            primitive.ObjectID     `json:"id,omitempty" bson:"_id,omitempty"`
	Title         string                 `json:"title" bson:"title"`
	AuthorID      primitive.ObjectID     `json:"author_id" bson:"author_id"`
	PublishedDate time.Time              `json:"published_date" bson:"published_date"`
	Details       map[string]interface{} `json:"details" bson:"details"`
	Category      string                 `json:"category" bson:"category"`
}
