package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	hydra "github.com/ory/hydra-client-go/client"
	"golang.org/x/crypto/bcrypt"

	"github.com/prmsrswt/edu-accounts/ent"
)

// API represents the HTTP API
type API struct {
	store     *ent.Client
	domain    string
	jwtSecret string
	hc        *hydra.OryHydra
}

// NewAPI creates a new instance of the API
func NewAPI(store *ent.Client, hydraAdminURL string) *API {
	return &API{
		store:     store,
		domain:    "iiitm.ac.in",
		jwtSecret: "jWt5upeRS3crE7",
		hc:        getHydraAdminClient(hydraAdminURL),
	}
}

// Register is used to register all of the API routes
func (a *API) Register(router *gin.Engine) {
	a.registerAPIRoutes(router.Group("/api"))
	a.registerOAuthRoutes(router.Group("/oauth"))
}

func generateHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func respondError(code int, err string, c *gin.Context) {
	c.JSON(code, gin.H{"status": "error", "error": err})
}

func respInternalServerErr(err error, c *gin.Context) {
	respondError(http.StatusInternalServerError, "Something went wrong", c)
	log.Println(err.Error())
}
