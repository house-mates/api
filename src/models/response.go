package models

type Response struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	EventID  uint   `json:"event_id"`
	UserID   string `json:"user_id"`
	Response string `json:"response"`
}
