package user

import (
	. "1024casts/backend/handler"
	"1024casts/backend/model"
	"1024casts/backend/pkg/errno"
	"1024casts/backend/util"

	"1024casts/backend/service"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /users [post]
func Create(c *gin.Context) {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
		Email:    r.Email,
	}

	// Validate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user password.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	// Insert the user to the database.
	srv := service.NewUserService()
	userId, err := srv.CreateUser(u)
	if err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	resp := CreateResponse{
		Id: userId,
	}

	// Show the user information.
	SendResponse(c, nil, resp)
}
