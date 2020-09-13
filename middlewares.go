package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tok, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "error": "Authentication required"})
			return
		}

		claims, err := verifyToken(tok)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "error": "Authentication required"})
			return
		}

		c.Set("email", claims["email"])
		c.Next()
	}
}
