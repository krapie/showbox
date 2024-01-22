package models

type Ticket struct {
	ID     int64 `json:"id"`
	ShowID int64 `json:"show_id"`
	UserID int64 `json:"user_id"`
}
