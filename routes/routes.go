package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/operation", controllers.Calculate)
	router.GET("/history", controllers.Historique)

	return router
}
