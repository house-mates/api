package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/house-mates/api/src/helpers"
	"github.com/house-mates/api/src/models"
)

// POST /users
// Add a new user
func AddUser(c *gin.Context) {
	var user = models.User{}.MapRequestToUser(c)

	result := models.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequest())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   1,
		Results: user,
	}})
}

// GET /users
// Get all users
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   len(users),
		Results: users,
	}})
}

// GET /users/:id
// Get a user by id
func FindUser(c *gin.Context) {
	var user models.User
	models.DB.First(&user, c.Param("id"))

	if user.ID == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   1,
		Results: user,
	}})
}

// PATCH /users/:id
// Edit an existing user
func EditUser(c *gin.Context) {
	var user = models.User{}.MapRequestToUser(c)

	user.ID = helpers.GrabIDParamAndConvertToUInt(c)

	result := models.DB.Save(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequest())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   1,
		Results: user,
	}})
}

// DELETE /users/:id
// Delete a user
func DeleteUser(c *gin.Context) {
	var user = models.User{
		ID: helpers.GrabIDParamAndConvertToUInt(c),
	}

	models.DB.First(&user)

	if user.Nickname == "" {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
