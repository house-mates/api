package models

type Action struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	CreatedBy uint   `json:"created_by"`
	Slug      string `json:"slug"`
	Title     string `json:"title"`
	Deleted   bool   `json:"deleted"`
}
