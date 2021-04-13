package models

type Usermeta struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	UserID    uint   `json:"user_id"`
	MetaKey   string `json:"meta_key"`
	MetaValue string `json:"meta_value"`
}
