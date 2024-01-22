package routes

import (
	controllers "task_manager/app/controllers"

	"github.com/gin-gonic/gin"
)

func GetAuthRoutes(router *gin.Engine) *gin.Engine {
	authRoutes := router.Group("/auth")
	{
		authRoutes.POST("/register", controllers.Register)
		authRoutes.POST("/login", controllers.Login)
	}

	return router
}
