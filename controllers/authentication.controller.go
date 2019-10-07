package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lbrulet/API-AWS-RDS/models"
	"github.com/lbrulet/API-AWS-RDS/services"
)

// LoginController user login endpoint
func LoginController(c *gin.Context) {
	payload := models.LoginPayload{}

	if err := c.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Error payload"})
		return
	}

	services.LoginService(c, payload)
}

// RegisterController user register endpoint
func RegisterController(c *gin.Context) {
	payload := models.RegisterPayload{}

	if err := c.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Error payload"})
		return
	}

	services.RegisterService(c, payload)
}
