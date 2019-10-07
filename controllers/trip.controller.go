package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lbrulet/API-AWS-RDS/models"
	"github.com/lbrulet/API-AWS-RDS/services"
)

// GetTrips get all trips or trips by user id
func GetTrips(c *gin.Context) {
	id := c.Query("id")
	services.GetTrips(c, id)
}

// NewTrip create a new trip
func NewTrip(c *gin.Context) {
	payload := models.Trip{}

	if err := c.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Error payload"})
		return
	}

	services.NewTrip(c, payload)
}
