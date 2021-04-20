package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/house-mates/api/helpers"
	"github.com/house-mates/api/models"
)

// POST /usermeta
// Add new usermeta
func AddUsermeta(c *gin.Context) {
	var usermeta = models.Usermeta{}.MapRequestToUsermeta(c)

	result := models.DB.Create(&usermeta)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequest())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   1,
		Results: usermeta,
	}})
}

// GET /usermeta
// Get all usermeta
func FindUsermeta(c *gin.Context) {
	var usermeta []models.Usermeta
	result := models.DB.Find(&usermeta)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   len(usermeta),
		Results: usermeta,
	}})
}

// GET /usermeta/:id
// Get usermeta by id
func FindUsermetaByID(c *gin.Context) {
	var usermeta models.Usermeta
	result := models.DB.First(&usermeta, c.Param("id"))

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   1,
		Results: usermeta,
	}})
}

// GET /usermeta_by_uid/:user_id
// Get usermeta by user id
func FindUsermetaByUserID(c *gin.Context) {
	var usermeta []models.Usermeta
	result := models.DB.Find(&usermeta, "user_id = ?", c.Param("user_id"))

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   len(usermeta),
		Results: usermeta,
	}})
}

// PATCH /usermeta/:id
// Edit existing usermeta
func EditUsermeta(c *gin.Context) {
	var usermeta = models.Usermeta{}.MapRequestToUsermeta(c)

	usermeta.ID = helpers.GrabIDParamAndConvertToUInt(c)

	result := models.DB.Save(&usermeta)

	// Patch does not care whether rows rows were affected or not.

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, helpers.BadRequest())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   1,
		Results: usermeta,
	}})
}

// DELETE /usermeta/:id
// Delete usermeta
func DeleteUsermeta(c *gin.Context) {
	var usermeta = models.Usermeta{
		ID: helpers.GrabIDParamAndConvertToUInt(c),
	}

	result := models.DB.First(&usermeta)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	if usermeta.UserID == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	models.DB.Delete(&usermeta)

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   1,
		Results: usermeta,
	}})
}

// DELETE /usermeta_by_uid/:user_id
// Delete usermeta by user id
func DeleteUsermetaByUserID(c *gin.Context) {
	var (
		userID   = c.Param("user_id")
		usermeta = []models.Usermeta{}
	)

	result := models.DB.Find(&usermeta, "user_id = ?", userID)

	if result.RowsAffected == 0 {
		c.JSON(http.StatusOK, helpers.NoResults())
		return
	}

	models.DB.Exec("DELETE FROM usermeta WHERE user_id = ?", userID)

	c.JSON(http.StatusOK, gin.H{"data": helpers.Results{
		Count:   len(usermeta),
		Results: usermeta,
	}})
}
