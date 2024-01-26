package main

import (
	"github.com/SanjaySinghRajpoot/FileRead/routes"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	routes.DataRoute(r)

	r.Run(":8080")
}
