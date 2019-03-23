package user

import (
	"strconv"

	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/service"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/gin-gonic/gin"
)

// @Summary Delete an user by the user identifier
// @Description Delete user by ID
// @Tags user
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The user's database id index num"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":null}"
// @Router /users/{id} [delete]
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))

	srv := service.NewUserService()
	err := srv.DeleteUser(uint64(userId))

	if err != nil {
		app.Response(c, errno.ErrDatabase, nil)
		return
	}

	app.Response(c, nil, nil)
}
