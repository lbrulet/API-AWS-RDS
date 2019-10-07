package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lbrulet/API-AWS-RDS/models"
	"github.com/lbrulet/API-AWS-RDS/services"
)

// RemoveUser remove a user
func RemoveUser(c *gin.Context) {
	var user models.BindingUser
	if err := c.ShouldBindUri(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Error uri"})
		return
	}
	services.RemoveUser(c, user)
}
