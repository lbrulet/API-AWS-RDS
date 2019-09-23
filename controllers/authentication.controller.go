package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lbrulet/API-AWS-RDS/database"
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

func SqlController(c *gin.Context) {
	db := database.DBManager.DB
	stmt, err := db.Prepare("CREATE Table Trip(id int NOT NULL AUTO_INCREMENT,start_lat float NOT NULL,start_lng float NOT NULL,end_lat float NOT NULL,	end_lng float NOT NULL, id_user int NOT NULL, PRIMARY KEY (id), FOREIGN KEY (id_user) REFERENCES User(id));")
	if err != nil {
		fmt.Println(err.Error())
	}

	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Table created successfully..")
	}
	c.JSON(http.StatusBadRequest, gin.H{"success": true, "message": "Table created successfully.."})
}
