package api

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prmsrswt/edu-accounts/ent"
	"github.com/prmsrswt/edu-accounts/ent/user"
)

func (a *API) queryAllUsers(c *gin.Context) {
	role := c.GetString("role")
	if role != "admin" {
		respondError(http.StatusForbidden, "forbidden", c)
		return
	}

	users, err := a.store.User.Query().WithCourse().WithDepartment().All(context.TODO())
	if err != nil {
		respInternalServerErr(fmt.Errorf("api: queryall: %w", err), c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "users": users})
}

func (a *API) verifyCredentials(email string, password string) (*ent.User, bool) {
	email = strings.ToLower(email)
	usr, err := a.store.User.Query().Where(user.EmailEQ(email)).Only(context.TODO())
	if err != nil {
		return nil, false
	}
	ok := checkPasswordHash(password, usr.Hash)
	if !ok {
		return nil, false
	}

	return usr, true
}

func (a *API) handleUserLogin(c *gin.Context) {
	var input struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
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

	token, err := a.createToken(usr.Email, usr.Role.String())
	if err != nil {
		respInternalServerErr(fmt.Errorf("login: create token: %w", err), c)
		return
	}

	c.SetCookie("token", token, 3600*48, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Successfully logged in"})
}
