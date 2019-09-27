package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lbrulet/API-AWS-RDS/database"
	"github.com/lbrulet/API-AWS-RDS/models"
)

func RemoveUser(c *gin.Context, user models.BindingUser) {
	db := database.DBManager.DB

	//Insert new user
	_, err := db.Exec(`DELETE FROM User WHERE id = ?`, user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Internal error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Success"})
}
