package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	hydra "github.com/ory/hydra-client-go/client"
	"github.com/ory/hydra-client-go/client/admin"
	hydraModels "github.com/ory/hydra-client-go/models"
)

func handleUserLogin(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	if !(strings.HasSuffix(input.Email, "iiitm.ac.in") && input.Password == "master") {
		c.JSON(http.StatusForbidden, gin.H{"status": "error", "error": "Invalid email or password"})
		return
	}

	token, err := createToken(input.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": "Something went wrong"})
		log.Printf("login: create token: %s\n", err)
		return
	}

	c.SetCookie("token", token, 3600*48, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Successfully logged in"})
}

func handleLoginChallenge(c *gin.Context) {
	lc := c.Query("login_challenge")
	if lc == "" {
		renderIndex(c)
		return
	}

	adminURL, _ := url.Parse("http://localhost:4445")
	adminC := hydra.NewHTTPClientWithConfig(nil, &hydra.TransportConfig{Schemes: []string{adminURL.Scheme}, Host: adminURL.Host, BasePath: adminURL.Path})

	res, err := adminC.Admin.GetLoginRequest(admin.NewGetLoginRequestParams().WithLoginChallenge(lc))
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}

	if *res.Payload.Skip {
		accRes, err := adminC.Admin.AcceptLoginRequest(
			admin.
				NewAcceptLoginRequestParams().
				WithLoginChallenge(lc).
				WithBody(&hydraModels.AcceptLoginRequest{Subject: res.Payload.Subject}),
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "Something went wrong")
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, *accRes.Payload.RedirectTo)
	}

	renderIndex(c)
}

func handleLoginChallengePost(c *gin.Context) {
	var input struct {
		Email     string `json:"email" binding:"required"`
		Password  string `json:"password" binding:"required"`
		Challenge string `json:"challenge" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	if !(strings.HasSuffix(input.Email, "iiitm.ac.in") && input.Password == "master") {
		c.JSON(http.StatusForbidden, gin.H{"status": "error", "error": "Invalid email or password"})
		return
	}

	adminURL, _ := url.Parse("http://localhost:4445")
	adminC := hydra.NewHTTPClientWithConfig(nil, &hydra.TransportConfig{Schemes: []string{adminURL.Scheme}, Host: adminURL.Host, BasePath: adminURL.Path})

	accRes, err := adminC.Admin.AcceptLoginRequest(
		admin.
			NewAcceptLoginRequestParams().
			WithLoginChallenge(input.Challenge).
			WithBody(&hydraModels.AcceptLoginRequest{Subject: &input.Email}),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "Something went wrong")
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "redirectTo": accRes.Payload.RedirectTo})
}
