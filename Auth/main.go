package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salimkun/Efishery-Test/Auth/common"
	"github.com/salimkun/Efishery-Test/Auth/service"
)

func main() {
	router := gin.Default()

	version1 := router.Group("/api/v1")
	version1.POST("/auth/register", service.RegisterUser) // here!
	version1.POST("/auth/login", service.LoginUser)

	protected := version1.Group("/user")
	protected.Use(common.JwtAuthMiddleware())
	protected.GET("/me", service.GetUserByToken)
	router.Run("localhost:8080")
}
