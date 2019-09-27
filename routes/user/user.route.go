package user

import (
	"github.com/gin-gonic/gin"
	"github.com/lbrulet/API-AWS-RDS/controllers"
)

// RegisterUserService add route handler from user
func RegisterUserService(route *gin.RouterGroup) {

	route.DELETE("/:id", controllers.RemoveUser)
}
