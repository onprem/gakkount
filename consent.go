package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ory/hydra-client-go/client/admin"
	hydraModels "github.com/ory/hydra-client-go/models"
)

func handleConsentMetadata(c *gin.Context) {
	cc := c.Param("challenge")

	ac := getHydraAdminClient()
	res, err := ac.Admin.GetConsentRequest(admin.NewGetConsentRequestParams().WithConsentChallenge(cc))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"scope":  res.Payload.RequestedScope,
		"user":   res.Payload.Subject,
		"client": res.Payload.Client,
	})
}

func handleConsent(c *gin.Context) {
	cc := c.Query("consent_challenge")
	if cc == "" {
		renderIndex(c)
		return
	}

	ac := getHydraAdminClient()
	res, err := ac.Admin.GetConsentRequest(admin.NewGetConsentRequestParams().WithConsentChallenge(cc))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": "Something went wrong"})
		return
	}

	if res.Payload.Skip {
		accRes, err := ac.Admin.AcceptConsentRequest(
			admin.
				NewAcceptConsentRequestParams().
				WithConsentChallenge(cc).WithBody(&hydraModels.AcceptConsentRequest{
				GrantScope:               res.Payload.RequestedScope,
				GrantAccessTokenAudience: res.Payload.RequestedAccessTokenAudience,
				Session: &hydraModels.ConsentRequestSession{
					IDToken: map[string]string{
						"name": "Dummy Dammani",
					},
				},
			}),
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": "Something went wrong"})
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, *accRes.Payload.RedirectTo)
	}

	renderIndex(c)
}

func handleConcentPost(c *gin.Context) {
	var input struct {
		Allow     bool   `json:"allow" binding:"required"`
		Challenge string `json:"challenge" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	ac := getHydraAdminClient()

	if !input.Allow {
		res, err := ac.Admin.RejectConsentRequest(
			admin.
				NewRejectConsentRequestParams().
				WithConsentChallenge(input.Challenge).WithBody(&hydraModels.RejectRequest{
				Error:            "access_denied",
				ErrorDescription: "The resource owner denied the request",
			}),
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": "Something went wrong"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "redirectTo": res.Payload.RedirectTo})
	}

	res, err := ac.Admin.GetConsentRequest(admin.NewGetConsentRequestParams().WithConsentChallenge(input.Challenge))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": "Something went wrong"})
		return
	}

	accRes, err := ac.Admin.AcceptConsentRequest(
		admin.
			NewAcceptConsentRequestParams().
			WithConsentChallenge(input.Challenge).WithBody(&hydraModels.AcceptConsentRequest{
			GrantScope:               res.Payload.RequestedScope,
			GrantAccessTokenAudience: res.Payload.RequestedAccessTokenAudience,
			Remember:                 true,
			RememberFor:              3600 * 48,
			Session: &hydraModels.ConsentRequestSession{
				IDToken: map[string]string{
					"name": "Dummy Dammani",
				},
			},
		}),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "error": "Something went wrong"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "redirectTo": accRes.Payload.RedirectTo})
}
