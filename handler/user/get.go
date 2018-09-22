package user

import (
	"strconv"

	. "1024casts/backend/handler"
	"1024casts/backend/model"
	"1024casts/backend/pkg/errno"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// @Summary Get an user by the user identifier
// @Description Get an user by username
// @Tags user
// @Accept  json
// @Produce  json
// @Param username path string true "Username"
// @Success 200 {object} model.UserModel "{"code":0,"message":"OK","data":{"username":"kong","password":"$2a$10$E0kwtmtLZbwW/bDQ8qI8e.eHPqhQOW9tvjwpyo/p05f/f4Qvr3OmS"}}"
// @Router /user/{username} [get]
func Get(c *gin.Context) {
	log.Info("Get function called.")

	// Get the user by the `username` from the database.
	//username := c.Param("username")
	//user, err := model.GetUser(username)

	// Get the user by the `id` from the database.
	// Get the user id from the url parameter.
	userId, _ := strconv.Atoi(c.Param("id"))

	user, err := model.GetUserById(userId)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
