package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID        primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title     string             `json:"title,omitempty" bson:"title,omitempty" validate:"required"`
	Content   string             `json:"content,omitempty" bson:"content,omitempty" validate:"required"`
	UserID    primitive.ObjectID `json:"userId,omitempty" bson:"userId,omitempty" validate:"required"`
	CreatedAt time.Time          `json:"createdAt,omitempty" bson:"createdAt,omitempty"`
	UpdatedAt time.Time          `json:"updatedAt,omitempty" bson:"updatedAt,omitempty"`
}
