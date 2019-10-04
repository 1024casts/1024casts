package user

import (
	"fmt"

	"github.com/1024casts/1024casts/model"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/pkg/notification"
	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"

	"net/http"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/pkg/flash"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func GetRegister(c *gin.Context) {
	c.HTML(http.StatusOK, "user/register", gin.H{
		"title": "VIP订阅",
		"ctx":   c,
	})
}

func DoRegister(c *gin.Context) {
	log.Info("User Register function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	var r RegisterRequest
	if err := c.Bind(&r); err != nil {
		app.Response(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
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
	userId, err := srv.RegisterUser(u)
	if err != nil {
		app.Response(c, errno.ErrDatabase, nil)
		return
	}

	// send to slack
	go func() {
		msg := fmt.Sprintf("welcome new user: %s[%s]", r.Username, r.Email)
		err := notification.SendNewRegisterUserNotification(msg)
		if err != nil {
			log.Warnf("[register] send msg to slack err, %v", err)
		}
	}()

	flash.SetMessage(c.Writer, "已发送激活链接,请检查您的邮箱。")

	resp := CreateResponse{
		Id: userId,
	}

	// Show the user information.
	app.Response(c, nil, resp)
}

// 通过链接激活用户
// 格式为：https://1024casts.com/users/test105/activation/rRDHuqemg
func ActiveUser(c *gin.Context) {
	token := c.Param("token")
	userActivation := model.UserActivationModel{}
	model.DB.Self.Table(userActivation.TableName()).Where("token = ?", token).First(&userActivation)

	// 存在说明需要激活
	if userActivation.UserID > 0 {
		srv := service.NewUserService()
		userInfo, err := srv.GetUserById(userActivation.UserID)
		if err != nil {
			log.Warnf("[register] active user err: %v", err)
		}

		if userInfo.IsActivated == 1 {
			// 提示： 您的帐号已经激活
			flash.SetMessage(c.Writer, "您的帐号已经激活")
			c.Redirect(http.StatusMovedPermanently, "/login")
			c.Abort()
			return
		}

		// 1. 更新用户is_activation == 1
		userMap := map[string]interface{}{
			"is_activated": 1,
		}
		err = srv.UpdateUser(userMap, userActivation.UserID)
		if err != nil {
			log.Warnf("[register] update user is_activation err: %v", err)
		}
		// 2. 删除 user_activation 表的记录
		model.DB.Self.Table(userActivation.TableName()).Where("token = ?", token).Delete(&userActivation)

		// 3. 提示：帐号已经激活成功，可以登录啦
		flash.SetMessage(c.Writer, "帐号已经激活成功，可以登录啦")
		c.Redirect(http.StatusMovedPermanently, "/login")
		c.Abort()
		return
	}

	// 提示：无效的激活链接
	flash.SetMessage(c.Writer, "无效的激活链接")
	c.Redirect(http.StatusMovedPermanently, "/login")
	c.Abort()
	return
}
