package user

import (
	"strconv"

	. "1024casts/backend/handler"
	"1024casts/backend/pkg/errno"

	"1024casts/backend/pkg/token"
	"1024casts/backend/service"

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
	srv := service.NewUserService()
	userId, _ := strconv.Atoi(c.Param("id"))

	// 从token 中解析 userId
	if userId == 0 {
		log.Info("user debug", lager.Data{"header": c.Request.Header.Get("Authorization")})
		// get userId from token
		ctx, _ := token.ParseRequest(c)
		userId = int(ctx.ID)
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
