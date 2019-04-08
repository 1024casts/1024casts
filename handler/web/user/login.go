package user

import (
	"github.com/1024casts/1024casts/pkg/auth"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/service"

	"net/http"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

func GetLogin(c *gin.Context) {
	frm := c.Query("frm")
	c.HTML(http.StatusOK, "user/login", gin.H{
		"title": "登录",
		"ctx":   c,
		"frm":   frm,
	})
}

// @Summary Login generates the authentication token
// @Produce  json
// @Param username body string true "Username"
// @Param password body string true "Password"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ"}}"
// @Router /login [post]
func DoLogin(c *gin.Context) {
	// Binding the data with the user struct.
	var u LoginCredentials
	if err := c.Bind(&u); err != nil {
		app.Response(c, errno.ErrBind, nil)
		return
	}

	srv := service.NewUserService()
	log.Infof("login data: %+v", u)
	// Get the user information by the login username.
	d, err := srv.GetUserByEmail(u.Email)
	if err != nil {
		app.Response(c, errno.ErrUserNotFound, nil)
		return
	}

	hashed, err := auth.Encrypt(u.Password)
	log.Infof("password hashed: %s", hashed)

	// Compare the login password with the user password.
	if err := auth.Compare(d.Password, u.Password); err != nil {
		app.Response(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// set cookie 24 hour
	c.SetCookie(viper.GetString("cookie.name"), util.EncodeUid(int64(d.Id)), viper.GetInt("cookie.max_age"), "/", "http://localhost:8888", false, true)

	app.Response(c, nil, nil)
}
