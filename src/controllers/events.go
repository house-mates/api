package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/house-mates/api/src/helpers"
	"github.com/house-mates/api/src/models"
)

// POST /events
// Add new event
func AddEvent(c *gin.Context) {
	var event = models.Event{}.MapRequestToEvent(c)

	var action models.Action
	result := models.DB.First(&action, event.ActionID)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"data": "action_id is not a valid action id"})
		return
	}

	result = models.DB.Create(&event)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequest())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   1,
		Results: event,
	}})
}

// GET /events
// Get all events
func FindEvents(c *gin.Context) {
	var events []models.Event
	result := models.DB.Find(&events)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   len(events),
		Results: events,
	}})
}

// GET /events/:id
// Get event
func FindEvent(c *gin.Context) {
	var event models.Event
	result := models.DB.First(&event, c.Param("id"))

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   1,
		Results: event,
	}})
}

// GET /events_by_aid/:action_id
// Get event by action id
func FindEventsByActionID(c *gin.Context) {
	var events []models.Event
	result := models.DB.Find(&events, "action_id = ?", c.Param("action_id"))

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   len(events),
		Results: events,
	}})
}

// GET /events_by_range
// Get event by time range
func FindEventsByTimeRange(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequest())
	}

	var timeRange = models.EventTimeRange{}

	err = json.Unmarshal(jsonData, &timeRange)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequest())
	}

	var events []models.Event

	result := models.DB.Find(
		&events,
		"time > ? AND time < ?",
		timeRange.After,
		timeRange.Before,
	)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   len(events),
		Results: events,
	}})
}

// GET /events_by_range_and_aid/:action_id
// Get event by time range and action id
func FindEventsByTimeRangeAndActionID(c *gin.Context) {
	jsonData, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequest())
	}

	var timeRange = models.EventTimeRange{}

	err = json.Unmarshal(jsonData, &timeRange)
	if err != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequest())
	}

	var events []models.Event

	result := models.DB.Find(
		&events,
		"time > ? AND time < ? AND action_id = ?",
		timeRange.After,
		timeRange.Before,
		c.Param("action_id"),
	)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   len(events),
		Results: events,
	}})
}

// PATCH /events/:id
// Edit existing event
func EditEvent(c *gin.Context) {
	var event = models.Event{}.MapRequestToEvent(c)

	event.ID = helpers.GrabIDParamAndConvertToUInt(c)

	result := models.DB.Save(&event)

	// Patch does not care whether rows rows were affected or not.

	// TODO: Is this neccessary?
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequest())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   1,
		Results: event,
	}})
}

// DELETE /events/:id
// Delete event
func DeleteEvent(c *gin.Context) {
	var event = models.Event{
		ID: helpers.GrabIDParamAndConvertToUInt(c),
	}

	result := models.DB.First(&event)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	models.DB.Delete(&event)

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   1,
		Results: event,
	}})
}

// DELETE /events_by_aid/:action_id
// Delete events by action id
func DeleteEventsByActionID(c *gin.Context) {
	var (
		actionID = c.Param("action_id")
		event    = []models.Event{}
	)

	result := models.DB.Find(&event, "action_id = ?", actionID)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	models.DB.Exec("DELETE FROM events WHERE action_id = ?", actionID)

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   len(event),
		Results: event,
	}})
}
