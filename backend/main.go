package main

import (
	"quant-backend/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Attach all your route groups here
	routes.RegisterRoutes(router)

	router.Run(":8080")
}
