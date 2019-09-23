package trip

import (
	"github.com/gin-gonic/gin"
	"github.com/lbrulet/API-AWS-RDS/controllers"
)

// RegisterTripService add route handler from trip
func RegisterTripService(route *gin.RouterGroup) {

	route.GET("/", controllers.GetTrips)
	route.POST("/", controllers.NewTrip)
}
