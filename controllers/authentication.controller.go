package controllers

import (
	"fmt"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/lbrulet/API-AWS-RDS/database"
	"github.com/lbrulet/API-AWS-RDS/models"
)

// LoginController user login endpoint
func LoginController(c *gin.Context) {
	db := database.DBManager.DB
	payload := models.LoginPayload{}
	var user models.User

	if err := c.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Error payload"})
		return
	}

	// Check if the user exist
	row := db.QueryRow(`SELECT * FROM User WHERE username = ?`, payload.Username)
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "User does not exist"})
		return
	}
	if user.Password != payload.Password {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Username or password are invalid"})
		return
	}

	// Create token
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := models.Claims{ID: user.ID, StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime.Unix()}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("pingouin123"))

	c.JSON(http.StatusOK, gin.H{"success": true, "message": user, "token": tokenString, "expiration_time": expirationTime})
}

// RegisterController user register endpoint
func RegisterController(c *gin.Context) {
	db := database.DBManager.DB
	payload := models.RegisterPayload{}
	var count int

	if err := c.ShouldBindBodyWith(&payload, binding.JSON); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "Error payload"})
		return
	}

	// Count rows
	row := db.QueryRow(`SELECT COUNT(*) FROM User WHERE username = ?`, payload.Username)
	err := row.Scan(&count)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "User already exist."})
		return
	}

	//Insert new user
	_, err = db.Exec(`INSERT INTO User (username, password)	VALUES (?, ?)`, payload.Username, payload.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": "Internal error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "User created with success"})
}
