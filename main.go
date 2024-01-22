package main

import (
	"fmt"
	models "task_manager/app/models"
	routes "task_manager/app/routers"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("from main")

	router := gin.Default()

	models.ConnectDatabase()

	routes.GetAuthRoutes(router)
	routes.GetUserRoutes(router)
	router.Run(":8080")
}
