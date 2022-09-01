package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	Id       primitive.ObjectID `json:"id,omitempty"`
	Email    string             `json:"email,omitempty" validate:"required" email:"true" unique:"true"`
	Password string             `json:"password,omitempty" validate:"required" minLength:"6"`
}
