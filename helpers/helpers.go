package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GrabIDParamAndConvertToUInt(c *gin.Context) uint {
	intValue, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic("Failed to convert string to int")
	}

	return uint(intValue)
}

func NoResults() gin.H {
	return gin.H{"data": Results{}}
}

func BadRequest() gin.H {
	return gin.H{"data": "Bad Request"}
}

type Results struct {
	Count   int         `json:"count"`
	Results interface{} `json:"results"`
}
