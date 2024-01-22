package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Show struct {
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title string             `json:"title" bson:"title,omitempty"`
}

func NewShow(title string) *Show {
	return &Show{
		Title: title,
	}
}
