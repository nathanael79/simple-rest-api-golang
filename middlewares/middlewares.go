package middlewares

import (
	"net/http"
	"simple-api/utils/token"

	"github.com/gin-gonic/gin"
)

// return value gin.HandlerFunc
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
