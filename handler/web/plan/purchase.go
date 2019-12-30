package plan

import (
	"net/http"

	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

// 购买会员
func Purchase(c *gin.Context) {
	userId := util.GetUserId(c)
	srv := service.NewUserService()
	user, _ := srv.GetUserById(userId)

	planSrv := service.NewPlanService()
	plan, err := planSrv.GetPlanByAlias(c.Param("alias"))
	if err != nil {
		log.Warnf("[plan] get plan list err: %v", err)
	}

	c.HTML(http.StatusOK, "plan/purchase", gin.H{
		"title":   "购买会员" + plan.Name,
		"user_id": userId,
		"user":    user,
		"ctx":     c,
		"plan":    plan,
		"qrCode":  "",
	})
}
