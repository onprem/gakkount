package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *API) registerAPIRoutes(router *gin.RouterGroup) {
	authR := router.Group("/")
	authR.Use(a.authMiddleware(false))

	adminR := router.Group("/")
	adminR.Use(a.authMiddleware(true))

	router.GET("/ping", func(c *gin.Context) { c.JSON(http.StatusOK, gin.H{"message": "pong"}) })
	router.POST("/login", a.handleUserLogin)
	router.GET("/logout", a.handleUserLogout)
	adminR.GET("/users", a.queryAllUsers)
	adminR.POST("/user", a.handleNewUser)
	authR.GET("/user", a.queryUser)
	authR.PATCH("/user", a.handleSelfUpdate)
	adminR.GET("/dummy", a.createDummy)

	adminR.GET("/courses", a.queryAllCourses)
	adminR.POST("/course", a.handleNewCourse)
	adminR.PATCH("/course", a.handleUpdateCourse)

	adminR.GET("/departments", a.queryAllDepartments)
	adminR.POST("/department", a.handleNewDepartment)
	adminR.PATCH("/department", a.handleUpdateDepartment)

	authR.GET("/clients", a.queryClients)
	authR.POST("/client", a.handleNewClient)
}

func (a *API) registerOAuthRoutes(router *gin.RouterGroup) {
	router.GET("/login", a.handleOAuthLogin)
	router.POST("/challenge", a.handleOAuthLoginPost)

	router.GET("/consent", a.handleConsent)
	router.GET("/consent/:challenge", a.handleConsentMetadata)
	router.POST("/consent", a.handleConsentPost)
}
