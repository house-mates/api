package models

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type Event struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	ActionID uint   `json:"action_id"`
	Time     string `json:"time"`
}

type EventTimeRange struct {
	After  string
	Before string
}

func (event Event) MapRequestToEvent(c *gin.Context) Event {
	jsonData, err := c.GetRawData()
	if err != nil {
		panic("Failed to get request body")
	}

	err = json.Unmarshal(jsonData, &event)
	if err != nil {
		panic("Failed to marshal json")
	}

	return event
}
