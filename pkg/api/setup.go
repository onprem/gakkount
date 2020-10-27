package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prmsrswt/edu-accounts/ent/user"
)

func (a *API) handleSetup(c *gin.Context) {
	users, err := a.store.User.Query().All(context.TODO())
	if err != nil {
		respInternalServerErr(fmt.Errorf("api: setup: query: %w", err), c)
		return
	}

	if len(users) > 0 {
		respondError(http.StatusBadRequest, "Setup only works on new instances", c)
		return
	}

	var i struct {
		Name     string `json:"name" binding:"required"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,min=4"`
	}

	if err := c.ShouldBindJSON(&i); err != nil {
		respondError(http.StatusBadRequest, err.Error(), c)
		return
	}
	i.Email = strings.ToLower(i.Email)
	if !strings.HasSuffix(i.Email, "@"+a.domain) {
		respondError(http.StatusBadRequest, "invalid primary email", c)
		return
	}

	hash, err := generateHash(i.Password)
	if err != nil {
		respInternalServerErr(fmt.Errorf("api: setup: %w", err), c)
		return
	}

	role := user.RoleAdmin

	x := a.store.User.Create().
		SetName(i.Name).
		SetEmail(i.Email).
		SetRole(role).SetHash(hash)

	usr, err := x.Save(context.TODO())
	if err != nil {
		respInternalServerErr(fmt.Errorf("api: setup: save: %w", err), c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Admin user created", "user": usr})
}
