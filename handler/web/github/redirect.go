package github

import (
	"net/http"

	"github.com/danilopolani/gocialite"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

var gocial = gocialite.NewDispatcher()

// Redirect to correct oAuth URL
func RedirectHandler(c *gin.Context) {
	authURL, err := gocial.New().
		Driver("github"). // Set provider
		Redirect(
			viper.GetString("github.client_id"),     // Client ID
			viper.GetString("github.client_secret"), // Client Secret
			viper.GetString("github.callback_url"),  // Redirect URL
		)

	// Check for errors (usually driver not valid)
	if err != nil {
		log.Warnf("[github] redirect err: %v", err)
		return
	}

	// Redirect with authURL
	c.Redirect(http.StatusFound, authURL) // Redirect with 302 HTTP code
}
