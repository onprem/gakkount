package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *API) authMiddleware(onlyAdmin bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		tok, err := c.Cookie("token")
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "error": "Authentication required"})
			return
		}

		claims, err := a.verifyToken(tok)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "error": "Authentication required"})
			return
		}

		if onlyAdmin && claims["role"] != "admin" {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "error": "Authentication required"})
			return
		}

		c.Set("email", claims["email"])
		c.Set("role", claims["role"])
		c.Next()
	}
}
