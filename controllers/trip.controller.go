package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lbrulet/API-AWS-RDS/models"
	"github.com/lbrulet/API-AWS-RDS/services"
)

func GetTrips(c *gin.Context) {
	id := c.Query("id")
	if len(id) > 0 {
		fmt.Println("id=", id)
	}
	services.GetTrips(c, id)
}

func NewTrip(c *gin.Context) {
	payload := models.Trip{}

	if err := c.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Error payload"})
		return
	}

	services.NewTrip(c, payload)
}
