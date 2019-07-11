package user

import (
	. "github.com/1024casts/1024casts/handler"
	"github.com/1024casts/1024casts/pkg/auth"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/service"
	"github.com/lexkong/log"

	"net/http"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/gin-gonic/gin"
)

func GetLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "user/login", gin.H{
		"title": "登录",
		"ctx":   c,
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
	// Get the user information by the login username.
	d, err := srv.GetUserByEmail(u.Email)
	if err != nil {
		log.Warnf("[login] get user by email err: %v", err)
		app.Response(c, errno.ErrUserNotFound, nil)
		return
	}

	// Compare the login password with the user password.
	if err := auth.Compare(d.Password, u.Password); err != nil {
		log.Warnf("[login] compare user password err: %v", err)
		app.Response(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// set cookie 30 day
	SetLoginCookie(c, d.Id)

	app.Response(c, nil, nil)
}
