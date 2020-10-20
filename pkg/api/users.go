package api

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/prmsrswt/edu-accounts/ent"
	"github.com/prmsrswt/edu-accounts/ent/user"
)

func (a *API) queryAllUsers(c *gin.Context) {
	roleFilter := user.Role(c.Query("role"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	page, _ := strconv.Atoi(c.Query("page"))

	query := a.store.User.Query().
		WithCourse().
		WithDepartment().
		Order(ent.Desc(user.FieldAdmissionTime)).
		Offset(page * limit)

	if limit != 0 {
		query = query.Limit(limit)
	}

	if roleFilter != "" {
		query = query.Where(user.RoleEQ(roleFilter))
	}

	users, err := query.All(context.TODO())
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
	c.SetCookie("signedin", "true", 3600*48, "/", "", false, false)
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Successfully logged in"})
}

func (a *API) handleUserLogout(c *gin.Context) {
	c.SetCookie("token", "", 0, "/", "", false, true)
	c.SetCookie("signedin", "false", 0, "/", "", false, false)
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Successfully logged out"})
}

func (a *API) queryUser(c *gin.Context) {
	email := c.GetString("email")

	usr, err := a.store.User.Query().Where(user.EmailEQ(email)).WithCourse().WithDepartment().Only(context.TODO())
	if err != nil {
		respInternalServerErr(fmt.Errorf("api: query user: %w", err), c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "user": usr})
}

func (a *API) handleSelfUpdate(c *gin.Context) {
	email := c.GetString("email")

	var i struct {
		Photo    string `json:"photo" binding:"omitempty,url"`
		AltEmail string `json:"altEmail" binding:"omitempty,email"`
		Phone    string `json:"phone" binding:"-"`
		// social
		LinkedIn string `json:"linkedin" binding:"omitempty,url"`
		Twitter  string `json:"twitter" binding:"omitempty,url"`
		Facebook string `json:"facebook" binding:"omitempty,url"`
		Github   string `json:"github" binding:"omitempty,url"`
	}

	if err := c.ShouldBindJSON(&i); err != nil {
		respondError(http.StatusBadRequest, err.Error(), c)
		return
	}

	x := a.store.User.Update().Where(user.EmailEQ(email))

	if i.Photo != "" {
		x.SetPhoto(i.Photo)
	}
	if i.AltEmail != "" {
		x.SetAltEmail(i.AltEmail)
	}
	if i.Phone != "" {
		x.SetPhone(i.Phone)
	}
	if i.LinkedIn != "" {
		x.SetLinkedin(i.LinkedIn)
	}
	if i.Twitter != "" {
		x.SetTwitter(i.Twitter)
	}
	if i.Facebook != "" {
		x.SetFacebook(i.Facebook)
	}
	if i.Github != "" {
		x.SetGithub(i.Github)
	}

	err := x.Exec(context.TODO())
	if err != nil {
		respInternalServerErr(fmt.Errorf("api: update user: %w", err), c)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User info updated"})
}
