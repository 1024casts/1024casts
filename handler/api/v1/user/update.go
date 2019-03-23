package user

import (
	"strconv"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// @Summary Update a user info by the user identifier
// @Description Update a user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The user's database id index num"
// @Param user body model.UserModel true "The user info"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":null}"
// @Router /users/{id} [put]
func Update(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	// Get the user id from the url parameter.
	userId, _ := strconv.Atoi(c.Param("id"))

	// Binding the user data.
	var u model.UserModel
	if err := c.Bind(&u); err != nil {
		app.Response(c, errno.ErrBind, nil)
		return
	}

	// We update the record based on the user id.
	u.Id = uint64(userId)

	srv := service.NewUserService()
	//user, err := model.GetUserById(userId)
	_, err := srv.GetUserById(uint64(userId))
	if err != nil {
		app.Response(c, errno.ErrUserNotFound, nil)
		log.Warn("user info", lager.Data{"id": userId})
		return
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

	// Save changed fields.
	userMap := make(map[string]interface{}, 0)
	userMap["username"] = u.Username
	userMap["avatar"] = u.Avatar
	userMap["real_name"] = u.RealName

	if err := srv.UpdateUser(userMap, u.Id); err != nil {
		app.Response(c, errno.ErrDatabase, nil)
		return
	}

	app.Response(c, nil, nil)
}
