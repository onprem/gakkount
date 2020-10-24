package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/prmsrswt/edu-accounts/ent"
	"github.com/prmsrswt/edu-accounts/ent/oclient"
	"github.com/prmsrswt/edu-accounts/ent/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/ory/hydra-client-go/client/admin"
	"github.com/ory/hydra-client-go/models"
	hydraModels "github.com/ory/hydra-client-go/models"
)

func (a *API) handleNewClient(c *gin.Context) {
	var i struct {
		Name          string   `json:"name" binding:"required"`
		GrantTypes    []string `json:"grantTypes" binding:"min=1"`
		ResponseTypes []string `json:"responseTypes" binding:"min=1"`
		Scope         []string `json:"scope" binding:"min=1"`
		Callbacks     []string `json:"callbacks" binding:"min=1"`
		Origins       []string `json:"origins" binding:"min=1"`
		LogoURI       string   `json:"logoURI" binding:"omitempty,url"`
		PrivacyURI    string   `json:"privacyURI" binding:"omitempty,url"`
		TosURI        string   `json:"tosURI" binding:"omitempty,url"`
	}
	if err := c.ShouldBindJSON(&i); err != nil {
		respondError(http.StatusBadRequest, err.Error(), c)
		return
	}

	email := c.GetString("email")

	usr, err := a.store.User.Query().Where(user.EmailEQ(email)).Only(context.TODO())
	if err != nil {
		respInternalServerErr(fmt.Errorf("api: hydra: create client: fetch user: %w", err), c)
		return
	}

	cid, err := uuid.NewRandom()
	if err != nil {
		respInternalServerErr(fmt.Errorf("api: hydra: create client: %w", err), c)
		return
	}

	res, err := a.hc.Admin.CreateOAuth2Client(admin.NewCreateOAuth2ClientParams().WithBody(&hydraModels.OAuth2Client{
		ClientID:           cid.String(),
		ClientName:         i.Name,
		GrantTypes:         hydraModels.StringSlicePipeDelimiter(i.GrantTypes),
		ResponseTypes:      hydraModels.StringSlicePipeDelimiter(i.ResponseTypes),
		Scope:              strings.Join(i.Scope, " "),
		RedirectUris:       hydraModels.StringSlicePipeDelimiter(i.Callbacks),
		AllowedCorsOrigins: hydraModels.StringSlicePipeDelimiter(i.Origins),
		LogoURI:            i.LogoURI,
		PolicyURI:          i.PrivacyURI,
		TosURI:             i.TosURI,
		Owner:              usr.Email,
	}))
	if err != nil {
		respInternalServerErr(fmt.Errorf("api: hydra: create client: %w", err), c)
		return
	}

	client, err := a.store.OClient.Create().
		SetClientID(cid.String()).
		SetSecret(res.Payload.ClientSecret).
		SetName(i.Name).
		SetUser(usr).
		Save(context.TODO())
	if err != nil {
		respInternalServerErr(fmt.Errorf("api: hydra: create client: db: %w", err), c)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "client": client, "payload": res.Payload})
}

type extClient struct {
	Client  *ent.OClient         `json:"client"`
	Payload *models.OAuth2Client `json:"payload"`
}

func (a *API) queryClients(c *gin.Context) {
	isAdmin := c.GetString("role") == "admin"
	email := c.GetString("email")

	fetchAll := c.Query("all") == "true"
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	query := a.store.OClient.Query().
		Order(ent.Desc(oclient.FieldID)).
		Offset(page * limit)

	if isAdmin && fetchAll {
		query = query.WithUser()
	} else {
		query = query.Where(oclient.HasUserWith(user.Email(email)))
	}

	if limit != 0 {
		query = query.Limit(limit)
	}

	clients, err := query.All(context.TODO())
	if err != nil {
		respInternalServerErr(fmt.Errorf("api: queryall clients: %w", err), c)
		return
	}

	res, err := a.hc.Admin.ListOAuth2Clients(admin.NewListOAuth2ClientsParams())
	if err != nil {
		respInternalServerErr(fmt.Errorf("api: queryall clients: list from hydra: %w", err), c)
		return
	}
	clientMap := make(map[string]*models.OAuth2Client)
	for _, v := range res.Payload {
		clientMap[v.ClientID] = v
	}

	list := make([]extClient, len(clients))
	for i, v := range clients {
		list[i] = extClient{Client: v, Payload: clientMap[v.ClientID]}
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "clients": list})
}
