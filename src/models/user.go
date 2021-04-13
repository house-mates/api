package models

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Nickname  string `json:"nickname"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}
