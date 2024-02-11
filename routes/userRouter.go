package routes

import (
	controller "organization-management-api/controllers"
	middleware "organization-management-api/middleware"

	"github.com/gin-gonic/gin"
)

// Change into ORGANIZATION
func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("users/",controller.GetUsers())
	incomingRoutes.POST("users/:user_id,controller.GetUser()")

}
