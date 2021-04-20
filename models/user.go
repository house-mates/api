package models

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	Nickname  string `json:"nickname"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (user User) MapRequestToUser(c *gin.Context) User {
	jsonData, err := c.GetRawData()
	if err != nil {
		panic("Failed to get request body")
	}

	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		panic("Failed to marshal json")
	}

	return user
}
