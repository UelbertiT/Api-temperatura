package routes

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		home := main.Group("home")
		{
			home.GET("/", controllers.Home)
		}
		temperatura := main.Group("temperatura")
		{
			temperatura.POST("/", controllers.CreateTemperatura)
			temperatura.GET("/", controllers.GetTemperatura)
		}
	}

	return router
}
