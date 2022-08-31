package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todos struct {
	Id    primitive.ObjectID `json:"id,omitempty"`
	Title string             `json:"title,omitempty" validate:"required"`
}
