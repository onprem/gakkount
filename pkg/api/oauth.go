package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/prmsrswt/edu-accounts/ent/user"
	"github.com/prmsrswt/edu-accounts/pkg/ui"

	"github.com/gin-gonic/gin"
	"github.com/ory/hydra-client-go/client/admin"
	hydraModels "github.com/ory/hydra-client-go/models"
)

func (a *API) handleOAuthLogin(c *gin.Context) {
	lc := c.Query("login_challenge")
	if lc == "" {
		ui.RenderIndex(c)
		return
	}

	res, err := a.hc.Admin.GetLoginRequest(admin.NewGetLoginRequestParams().WithLoginChallenge(lc))
	if err != nil {
		respInternalServerErr(fmt.Errorf("login: hydra: get login req: %w", err), c)
		return
	}

	if *res.Payload.Skip {
		accRes, err := a.hc.Admin.AcceptLoginRequest(
			admin.
				NewAcceptLoginRequestParams().
				WithLoginChallenge(lc).
				WithBody(&hydraModels.AcceptLoginRequest{Subject: res.Payload.Subject}),
		)
		if err != nil {
			respInternalServerErr(fmt.Errorf("login: hydra: accept login req: %w", err), c)
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, *accRes.Payload.RedirectTo)
	}

	ui.RenderIndex(c)
}

func (a *API) handleOAuthLoginPost(c *gin.Context) {
	var input struct {
		Email     string `json:"email" binding:"required"`
		Password  string `json:"password" binding:"required"`
		Challenge string `json:"challenge" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		respondError(http.StatusBadRequest, err.Error(), c)
		return
	}

	usr, ok := a.verifyCredentials(input.Email, input.Password)
	if !ok {
		respondError(http.StatusForbidden, "Invalid email or password", c)
		return
	}

	accRes, err := a.hc.Admin.AcceptLoginRequest(
		admin.
			NewAcceptLoginRequestParams().
			WithLoginChallenge(input.Challenge).
			WithBody(&hydraModels.AcceptLoginRequest{
				Subject:     &usr.Email,
				Remember:    true,
				RememberFor: 0,
			}),
	)
	if err != nil {
		respInternalServerErr(fmt.Errorf("login: hydra: accept login req: %w", err), c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "redirectTo": accRes.Payload.RedirectTo})
}

func (a *API) handleConsentMetadata(c *gin.Context) {
	cc := c.Param("challenge")

	res, err := a.hc.Admin.GetConsentRequest(admin.NewGetConsentRequestParams().WithConsentChallenge(cc))
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

func (a *API) acceptConsent(challenge string, payload *hydraModels.ConsentRequest) (*admin.AcceptConsentRequestOK, error) {
	email := strings.ToLower(payload.Subject)
	usr, err := a.store.User.Query().Where(user.EmailEQ(email)).WithCourse().WithDepartment().Only(context.TODO())
	if err != nil {
		return nil, err
	}
	idTkn := map[string]interface{}{
		"name":  usr.Name,
		"email": usr.Email,
		"role":  usr.Role.String(),
	}
	if usr.Role.String() == "student" {
		idTkn["rollNo"] = usr.RollNo
		idTkn["admissionTime"] = usr.AdmissionTime
		idTkn["courseEndTime"] = usr.CourseEndTime
		idTkn["course"] = usr.Edges.Course
	}
	return a.hc.Admin.AcceptConsentRequest(
		admin.
			NewAcceptConsentRequestParams().
			WithConsentChallenge(challenge).WithBody(&hydraModels.AcceptConsentRequest{
			GrantScope:               payload.RequestedScope,
			GrantAccessTokenAudience: payload.RequestedAccessTokenAudience,
			Remember:                 false,
			Session: &hydraModels.ConsentRequestSession{
				IDToken: idTkn,
			},
		}),
	)
}

func (a *API) handleConsent(c *gin.Context) {
	cc := c.Query("consent_challenge")
	if cc == "" {
		ui.RenderIndex(c)
		return
	}
	res, err := a.hc.Admin.GetConsentRequest(admin.NewGetConsentRequestParams().WithConsentChallenge(cc))
	if err != nil {
		respInternalServerErr(fmt.Errorf("handle consent: %w", err), c)
		return
	}

	if res.Payload.Skip {
		accRes, err := a.acceptConsent(cc, res.Payload)
		if err != nil {
			respInternalServerErr(fmt.Errorf("accept consent: %w", err), c)
			return
		}
		c.Redirect(http.StatusTemporaryRedirect, *accRes.Payload.RedirectTo)
		return
	}

	ui.RenderIndex(c)
}

func (a *API) handleConsentPost(c *gin.Context) {
	var input struct {
		Allow     bool   `json:"allow"`
		Challenge string `json:"challenge" binding:"required"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error", "error": err.Error()})
		return
	}

	if !input.Allow {
		res, err := a.hc.Admin.RejectConsentRequest(
			admin.
				NewRejectConsentRequestParams().
				WithConsentChallenge(input.Challenge).WithBody(&hydraModels.RejectRequest{
				Error:            "access_denied",
				ErrorDescription: "The resource owner denied the request",
			}),
		)
		if err != nil {
			respInternalServerErr(fmt.Errorf("reject consent: %w", err), c)
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "redirectTo": res.Payload.RedirectTo})
		return
	}

	res, err := a.hc.Admin.GetConsentRequest(admin.NewGetConsentRequestParams().WithConsentChallenge(input.Challenge))
	if err != nil {
		respInternalServerErr(fmt.Errorf("get consent: %w", err), c)
		return
	}

	accRes, err := a.acceptConsent(input.Challenge, res.Payload)
	if err != nil {
		respInternalServerErr(fmt.Errorf("accept consent: %w", err), c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "redirectTo": accRes.Payload.RedirectTo})
}
