package models

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Nickname  string `json:"nickname"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Usermeta struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	UserID    uint   `json:"user_id"`
	MetaKey   string `json:"meta_key"`
	MetaValue string `json:"meta_value"`
}

type Action struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	CreatedBy uint   `json:"created_by"`
	Slug      string `json:"slug"`
	Title     string `json:"title"`
	Deleted   bool   `json:"deleted"`
}

type Event struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	ActionID uint   `json:"action_id"`
	Time     string `json:"timestamp"`
}

type Response struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	EventID  uint   `json:"event_id"`
	UserID   string `json:"user_id"`
	Response string `json:"response"`
}
