package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/salimkun/Efishery-Test/Auth/common/util"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := util.TokenValid(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "UnAuthorized"})
			return
		}
		c.Next()
	}
}
