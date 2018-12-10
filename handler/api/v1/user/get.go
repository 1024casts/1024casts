package user

import (
	"strconv"

	. "github.com/1024casts/1024casts/handler"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/pkg/token"
	"github.com/1024casts/1024casts/service"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// @Summary Get an user by the user identifier
// @Description Get an user by username
// @Tags user
// @Accept  json
// @Produce  json
// @Param username path string true "Username"
// @Success 200 {object} model.UserModel "{"code":0,"message":"OK","data":{"username":"kong","password":"$2a$10$E0kwtmtLZbwW/bDQ8qI8e.eHPqhQOW9tvjwpyo/p05f/f4Qvr3OmS"}}"
// @Router /users/{username} [get]
func Get(c *gin.Context) {
	log.Info("Get function called.")

	// Get the user by the `id` from the database.
	// Get the user id from the url parameter.

	var userId uint64
	srv := service.NewUserService()
	uId, _ := strconv.Atoi(c.Param("id"))
	userId = uint64(uId)

	// 从token 中解析 userId
	if userId == 0 {
		log.Info("user debug", lager.Data{"header": c.Request.Header.Get("Authorization")})
		// get userId from token
		ctx, _ := token.ParseRequest(c)
		userId = ctx.ID
	}

	user, err := srv.GetUserById(userId)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	roles := []string{"admin"}
	user.Roles = roles

	SendResponse(c, nil, user)
}
