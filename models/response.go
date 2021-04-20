package models

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type Response struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	EventID  uint   `json:"event_id"`
	UserID   uint   `json:"user_id"`
	Response string `json:"response"`
}

func (response Response) MapRequestToResponse(c *gin.Context) Response {
	jsonData, err := c.GetRawData()
	if err != nil {
		panic("Failed to get request body")
	}

	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		panic("Failed to marshal json")
	}

	return response
}
