package main

import (
	"github.com/SanjaySinghRajpoot/FileRead/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	r := gin.Default()

	// Global middleware: recovery and logging
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	// Register routes
	registerRoutes(r)

	// Serve Swagger documentation
	serveSwagger(r)

	r.Run(":8080")
}

// registerRoutes registers application-specific routes
func registerRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to FileRead!",
		})
	})

	routes.DataRoute(r)
}

func serveSwagger(r *gin.Engine) {
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
