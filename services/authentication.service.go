package services

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/lbrulet/API-AWS-RDS/database"
	"github.com/lbrulet/API-AWS-RDS/models"
)

func LoginService(c *gin.Context, payload models.LoginPayload) {
	db := database.DBManager.DB
	var user models.User

	// Check if the user exist
	row := db.QueryRow(`SELECT * FROM User WHERE username = ?`, payload.Username)
	if err := row.Scan(&user.ID, &user.Username, &user.Password); err != nil {
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
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": user, "token": tokenString, "expiration_time": expirationTime.Unix()})
}

func RegisterService(c *gin.Context, payload models.RegisterPayload) {
	db := database.DBManager.DB
	var count int

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
