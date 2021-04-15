package models

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

type Usermeta struct {
	ID        uint   `json:"id" gorm:"primary_key"`
	UserID    uint   `json:"user_id"`
	MetaKey   string `json:"meta_key"`
	MetaValue string `json:"meta_value"`
}

func (usermeta Usermeta) MapRequestToUsermeta(c *gin.Context) Usermeta {
	jsonData, err := c.GetRawData()
	if err != nil {
		panic("Failed to get request body")
	}

	err = json.Unmarshal(jsonData, &usermeta)
	if err != nil {
		panic("Failed to marshal json")
	}

	return usermeta
}
