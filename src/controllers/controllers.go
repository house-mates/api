package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/house-mates/api/src/models"
)

// GET /users
// Get all users
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GET /users/:id
// Get a user by id
func FindUser(c *gin.Context) {
	var user models.User
	models.DB.First(&user, c.Param("id"))

	c.JSON(http.StatusOK, gin.H{"data": user})
}
