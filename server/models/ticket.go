package models

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	AVAILABLE = "available"
	RESERVED  = "reserved"
	BOUGHT    = "bought"
)

type Ticket struct {
	ID     primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ShowID primitive.ObjectID `json:"show_id" bson:"show_id"`
	UserID primitive.ObjectID `json:"user_id" bson:"user_id"`
}

func NewTicket(ID, ShowID, UserID primitive.ObjectID) *Ticket {
	return &Ticket{
		ID:     ID,
		ShowID: ShowID,
		UserID: UserID,
	}
}
