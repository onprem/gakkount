package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *API) registerAPIRoutes(router *gin.RouterGroup) {
	authR := router.Group("/")
	authR.Use(a.authMiddleware())

	router.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "pong"}) })
	router.POST("/login", a.handleUserLogin)
	authR.GET("/users", a.queryAllUsers)
}

func (a *API) registerOAuthRoutes(router *gin.RouterGroup) {
	router.GET("/login", a.handleOAuthLogin)
	router.POST("/challenge", a.handleOAuthLoginPost)

	router.GET("/consent", a.handleConsent)
	router.GET("/consent/:challenge", a.handleConsentMetadata)
	router.POST("/consent", a.handleConsentPost)
}
