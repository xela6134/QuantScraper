package routes

import (
	"quant-backend/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterScrapeRoutes(router *gin.Engine) {
	scrape := router.Group("/scrape")
	{
		scrape.GET("/janestreet", controllers.ScrapeJaneStreet)
		// scrape.GET("/imc", controllers.ScrapeIMC)
		// scrape.GET("/hrt", controllers.ScrapeHRT)
		// TODO: add more scrape routes here
	}
}
