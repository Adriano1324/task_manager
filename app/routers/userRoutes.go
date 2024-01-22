package routes

import (
	controllers "task_manager/app/controllers"
	middlewares "task_manager/app/middlewares"

	"github.com/gin-gonic/gin"
)

func GetUserRoutes(router *gin.Engine) *gin.Engine {
	userRoutes := router.Group("/user")
	{
		userRoutes.Use(middlewares.JwtAuthMiddleware())
		userRoutes.GET("", controllers.CurrentUser)
	}

	return router
}
