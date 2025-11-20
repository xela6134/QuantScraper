package routes

import (
	"quant-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Just an example route
	ping := router.Group("/ping")
	{
		ping.GET("", controllers.Ping)
	}

	RegisterScrapeRoutes(router)
}
