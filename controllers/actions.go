package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/house-mates/api/helpers"
	"github.com/house-mates/api/models"
)

// POST /actions
// Add new action
func AddAction(c *gin.Context) {
	var action = models.Action{}.MapRequestToAction(c)

	var user models.User
	result := models.DB.First(&user, action.CreatedBy)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"data": "created_by is not a valid user id"})
		return
	}

	result = models.DB.Create(&action)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequest())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   1,
		Results: action,
	}})
}

// GET /actions
// Get all actions
func FindActions(c *gin.Context) {
	var action []models.Action
	result := models.DB.Find(&action)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   len(action),
		Results: action,
	}})
}

// GET /actions/:id
// Get action
func FindAction(c *gin.Context) {
	var action models.Action
	result := models.DB.First(&action, c.Param("id"))

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   1,
		Results: action,
	}})
}

// GET /actions_by_uid/:user_id
// Get action by user id
func FindActionByUserID(c *gin.Context) {
	var action []models.Action
	result := models.DB.Find(&action, "created_by = ?", c.Param("user_id"))

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   len(action),
		Results: action,
	}})
}

// PATCH /actions/:id
// Edit existing action
func EditAction(c *gin.Context) {
	var action = models.Action{}.MapRequestToAction(c)

	action.ID = helpers.GrabIDParamAndConvertToUInt(c)

	result := models.DB.Save(&action)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	// TODO: Is this neccessary?
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequest())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   1,
		Results: action,
	}})
}

// DELETE /actions/:id
// Delete action
func DeleteAction(c *gin.Context) {
	var action = models.Action{
		ID: helpers.GrabIDParamAndConvertToUInt(c),
	}

	result := models.DB.First(&action)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	models.DB.Delete(&action)

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   1,
		Results: action,
	}})
}

// DELETE /actions_by_uid/:user_id
// Delete actions by user id
func DeleteActionsByUserID(c *gin.Context) {
	var (
		userID = c.Param("user_id")
		action = []models.Action{}
	)

	result := models.DB.Find(&action, "created_by = ?", userID)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	models.DB.Exec("DELETE FROM actions WHERE created_by = ?", userID)

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   len(action),
		Results: action,
	}})
}
