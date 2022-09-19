package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salimkun/Efishery-Test/Auth/common"
	"github.com/salimkun/Efishery-Test/Fetch/service"
)

func main() {
	router := gin.Default()

	version1 := router.Group("/api/v1")

	version1.Use(common.JwtAdminMiddleware())
	version1.GET("/resource", service.GetResource)
	version1.GET("/resource/agregate", service.AgregateResource)
	router.Run("localhost:8081")
}
