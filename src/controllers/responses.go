package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/house-mates/api/src/helpers"
	"github.com/house-mates/api/src/models"
)

// POST /responses
// Add new response
func AddResponse(c *gin.Context) {
	var (
		response = models.Response{}.MapRequestToResponse(c)
		event    models.Event
		user     models.User
	)

	result := models.DB.First(&event, response.EventID)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"data": "event_id is not a valid event id"})
		return
	}

	result = models.DB.First(&user, response.UserID)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"data": "user_id is not a valid user id"})
		return
	}

	result = models.DB.Create(&response)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequest())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   1,
		Results: response,
	}})
}

// GET /responses
// Get all responses
func FindResponses(c *gin.Context) {
	var responses []models.Response
	result := models.DB.Find(&responses)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   len(responses),
		Results: responses,
	}})
}

// GET /responses/:id
// Get response
func FindResponse(c *gin.Context) {
	var response models.Response
	result := models.DB.First(&response, c.Param("id"))

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   1,
		Results: response,
	}})
}

// GET /responses_by_eid/:event_id
// Get responses by event id
func FindResponsesByEventID(c *gin.Context) {
	var responses []models.Response
	result := models.DB.Find(&responses, "event_id = ?", c.Param("event_id"))

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   len(responses),
		Results: responses,
	}})
}

// GET /responses_by_uid/:user_id
// Get responses by user id
func FindResponsesByUserID(c *gin.Context) {
	var responses []models.Response
	result := models.DB.Find(&responses, "user_id = ?", c.Param("user_id"))

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   len(responses),
		Results: responses,
	}})
}

// DELETE /responses/:id
// Delete response
func DeleteResponse(c *gin.Context) {
	var response = models.Response{
		ID: helpers.GrabIDParamAndConvertToUInt(c),
	}

	result := models.DB.First(&response)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	models.DB.Delete(&response)

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   1,
		Results: response,
	}})
}

// DELETE /responses_by_eid/:event_id
// Delete responses by event id
func DeleteResponsesByEventID(c *gin.Context) {
	var (
		eventID   = c.Param("event_id")
		responses = []models.Response{}
	)

	result := models.DB.Find(&responses, "event_id = ?", eventID)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	models.DB.Exec("DELETE FROM responses WHERE event_id = ?", eventID)

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   len(responses),
		Results: responses,
	}})
}

// DELETE /responses_by_uid/:user_id
// Delete responses by user id
func DeleteResponsesByUserID(c *gin.Context) {
	var (
		userID    = c.Param("user_id")
		responses = []models.Response{}
	)

	result := models.DB.Find(&responses, "user_id = ?", userID)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	models.DB.Exec("DELETE FROM responses WHERE user_id = ?", userID)

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   len(responses),
		Results: responses,
	}})
}
