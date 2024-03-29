package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/lbrulet/API-AWS-RDS/routes/authentication"
	"github.com/lbrulet/API-AWS-RDS/routes/trip"
	"github.com/lbrulet/API-AWS-RDS/routes/user"
)

// CORS allow request from outside
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers, authorization, content-type")
		c.Header("Content-Type", "application/json")
	}
}

// InitRouter return a server
func InitRouter() *gin.Engine {
	router := gin.New()

	router.Use(CORS(), gin.Logger(), gin.Recovery())

	api := router.Group("/api")

	auth := api.Group("/auth")
	users := api.Group("/user")
	trips := api.Group("/trips")

	authentication.RegisterAuthService(auth)
	trip.RegisterTripService(trips)
	user.RegisterUserService(users)

	return router
}
