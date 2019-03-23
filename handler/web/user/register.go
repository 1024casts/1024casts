package user

import (
	. "github.com/1024casts/1024casts/handler"
	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"

	"net/http"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func GetRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "user/register", gin.H{
		"title":   "VIP订阅",
		"user_id": 1,
	})
}

func Register(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		app.Response(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}

	// Validate the data.
	if err := u.Validate(); err != nil {
		app.Response(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		app.Response(c, errno.ErrEncrypt, nil)
		return
	}
	// Insert the user to the database.
	srv := service.NewUserService()
	userId, err := srv.CreateUser(u)
	if err != nil {
		app.Response(c, errno.ErrDatabase, nil)
		return
	}

	resp := CreateResponse{
		Id: userId,
	}

	// Show the user information.
	app.Response(c, nil, resp)
}
