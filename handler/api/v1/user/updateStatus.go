package user

import (
	"strconv"

	. "github.com/1024casts/1024casts/handler"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"

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
// @Router /users/{id}/status []
func UpdateStatus(c *gin.Context) {
	log.Info("Update function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	// Get the user id from the url parameter.
	userId, _ := strconv.Atoi(c.Param("id"))

	// Binding the user data.
	var req UpdateReq
	if err := c.Bind(&req); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	srv := service.NewUserService()
	_, err := srv.GetUserById(uint64(userId))
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		log.Warn("user info", lager.Data{"id": userId})
		return
	}

	// Save changed fields.
	userMap := make(map[string]interface{}, 0)
	userMap["status"] = req.Status

	if err := srv.UpdateUser(userMap, uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
