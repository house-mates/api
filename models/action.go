package models

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type Action struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	CreatedBy uint   `json:"created_by"`
	Slug      string `json:"slug"`
	Title     string `json:"title"`
	Deleted   bool   `json:"deleted"`
}

func (action Action) MapRequestToAction(c *gin.Context) Action {
	jsonData, err := c.GetRawData()
	if err != nil {
		panic("Failed to get request body")
	}

	err = json.Unmarshal(jsonData, &action)
	if err != nil {
		panic("Failed to marshal json")
	}

	return action
}
