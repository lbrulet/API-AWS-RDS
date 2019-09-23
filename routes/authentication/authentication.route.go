package authentication

import (
	"github.com/gin-gonic/gin"
	"github.com/lbrulet/API-AWS-RDS/controllers"
)

// RegisterAuthService add route handler from the authentication
func RegisterAuthService(route *gin.RouterGroup) {

	route.POST("/login", controllers.LoginController)

	route.POST("/register", controllers.RegisterController)

	route.GET("/sql", controllers.SqlController)
}
