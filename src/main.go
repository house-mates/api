package main

import (
	"github.com/gin-gonic/gin"

	"github.com/house-mates/api/src/controllers"
	"github.com/house-mates/api/src/models"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/users", controllers.FindUsers)

	r.Run()
}
