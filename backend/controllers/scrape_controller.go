package controllers

import (
	"net/http"
	"quant-backend/scraper"

	"github.com/gin-gonic/gin"
)

func ScrapeJaneStreet(c *gin.Context) {
	data := scraper.ScrapeJaneStreet()
	c.JSON(http.StatusOK, data)
}

// func ScrapeIMC(c *gin.Context) {
// 	data := scraper.ScrapeIMC()
// 	c.JSON(http.StatusOK, data)
// }

// func ScrapeHRT(c *gin.Context) {
// 	data := scraper.ScrapeHRT()
// 	c.JSON(http.StatusOK, data)
// }
