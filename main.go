package main

import (
	"go-lang-website/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.RegisterRoutes(router)

	// Serve static files and the React app
	router.Static("/static", "./static")
	router.NoRoute(func(c *gin.Context) {
		c.File("./static/index.html")
	})

	router.Run(":8080")
}
