package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/house-mates/api/src/helpers"
	"github.com/house-mates/api/src/models"
)

// POST /usermeta
// Add new usermeta
func AddUsermeta(c *gin.Context) {
	var usermeta = models.Usermeta{}.MapRequestToUsermeta(c)

	result := models.DB.Create(&usermeta)

	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(http.StatusOK, gin.H{"data": usermeta})
}

// GET /usermeta
// Get all usermeta
func FindUsermeta(c *gin.Context) {
	var usermeta []models.Usermeta
	models.DB.Find(&usermeta)

	c.JSON(http.StatusOK, gin.H{"data": usermeta})
}

// GET /usermeta/:id
// Get usermeta by id
func FindUsermetaByID(c *gin.Context) {
	var usermeta models.Usermeta
	models.DB.First(&usermeta, c.Param("id"))

	if usermeta.ID == 0 {
		c.JSON(http.StatusOK, helpers.NotFoundResponse())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": usermeta})
}

// GET /usermeta_by_uid/:user_id
// Get usermeta by user id
func FindUsermetaByUserID(c *gin.Context) {
	var usermeta models.Usermeta
	models.DB.First(&usermeta, "user_id = ?", c.Param("user_id"))

	if usermeta.ID == 0 {
		c.JSON(http.StatusOK, helpers.NotFoundResponse())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": usermeta})
}

// PATCH /usermeta/:id
// Edit existing usermeta
func EditUsermeta(c *gin.Context) {
	var usermeta = models.Usermeta{}.MapRequestToUsermeta(c)

	usermeta.ID = helpers.GrabIDParamAndConvertToUInt(c)

	result := models.DB.Save(&usermeta)

	if result.Error != nil {
		panic(result.Error)
	}

	c.JSON(http.StatusOK, gin.H{"data": usermeta})
}

// DELETE /usermeta/:id
// Delete usermeta
func DeleteUsermeta(c *gin.Context) {
	var usermeta = models.Usermeta{
		ID: helpers.GrabIDParamAndConvertToUInt(c),
	}

	models.DB.First(&usermeta)

	models.DB.Delete(&usermeta)

	c.JSON(http.StatusOK, gin.H{"data": usermeta})
}
