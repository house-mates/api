package main

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
