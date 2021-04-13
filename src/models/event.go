package models

type Event struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	ActionID uint   `json:"action_id"`
	Time     string `json:"timestamp"`
}
