package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	authorized := router.Group("/")
	authorized.Use(authMiddleware())

	oauthR := router.Group("/oauth")

	authorized.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})
	oauthR.GET("/login", handleLoginChallenge)
	oauthR.POST("/login", handleUserLogin)
	oauthR.POST("/challenge", handleLoginChallengePost)

	oauthR.GET("/consent", handleConsent)
	oauthR.GET("/consent/:challenge", handleConsentMetadata)
	oauthR.POST("/consent", handleConcentPost)

	router.NoRoute(serveUI)

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,

		// Good practice to set timeouts to avoid Slowloris attacks.
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s\n", err)
	}
	log.Println("Server exiting")
}
