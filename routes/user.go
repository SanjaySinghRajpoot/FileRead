package routes

import (
	"github.com/SanjaySinghRajpoot/FileRead/controller"
	"github.com/gin-gonic/gin"
)

func DataRoute(router *gin.Engine) {

	// All user routes
	router.GET("/data", controller.GetData)
}
