package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("user_id")
		if err != nil {
			c.Redirect(http.StatusNotFound, "/login")
			c.Abort()
			return
		}
		c.Next()
	}
}
