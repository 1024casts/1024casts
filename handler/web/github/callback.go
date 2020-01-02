package github

import (
	"fmt"
	"net/http"
	"time"

	"github.com/1024casts/1024casts/handler"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/service"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/lexkong/log"
)

// Redirect to correct oAuth URL
// Handle callback of provider
func CallbackHandler(c *gin.Context) {
	// Retrieve query params for code and state
	code := c.Query("code")
	state := c.Query("state")

	// Handle callback and check for errors
	githubUser, _, err := gocial.Handle(state, code)
	if err != nil {
		log.Warnf("[github] callback err: %v", err)
		c.Redirect(http.StatusFound, "/login")
		return
	}

	// Print in terminal user information
	//fmt.Printf("%#v", token)
	//fmt.Printf("%#v", user)

	fmt.Printf("github user_id: %s", githubUser.ID)
	fmt.Printf("github username: %s", githubUser.Username)

	// 如果用户存在，则更新最后登录时间和ip
	userSrv := service.NewUserService()
	user, err := userSrv.GetUserByGithubId(githubUser.ID)

	var userId uint64
	// if not exist, 创建新用户
	if err == gorm.ErrRecordNotFound {
		u := model.UserModel{
			Username:      githubUser.Username,
			Avatar:        githubUser.Avatar,
			Password:      "",
			Email:         githubUser.Email,
			GithubId:      githubUser.ID,
			GithubName:    githubUser.Username,
			LastLoginTime: time.Now(),
			LastLoginIp:   c.ClientIP(),
			IsActivated:   1,
		}
		userId, err = userSrv.CreateUser(u)
		if err != nil {
			log.Warnf("[github] create user err: %v", err)
			c.Redirect(http.StatusFound, "/login")
			return
		}
	} else {
		// get github info, err is nil
		userId = user.Id
		if err == nil {
			err := userSrv.UpdateLastLoginInfo(userId, c.ClientIP())
			if err != nil {
				log.Warnf("[github] update user login info err: %v", err)
				c.Redirect(http.StatusFound, "/login")
				return
			}
		}
	}

	// 执行登录操作
	handler.SetLoginCookie(c, userId)

	c.Redirect(http.StatusFound, "/")
	return
}
