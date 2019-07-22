package plan

import (
	"net/http"

	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	userId := util.GetUserId(c)
	user, _ := service.UserService.GetUserById(userId)

	planMap := make(map[string]interface{})
	planMap["status"] = 1
	planSrv := service.NewPlanService()
	plans, _, err := planSrv.GetPlanList(planMap, 0, 10)
	if err != nil {
		log.Warnf("[plan] get plan list err: %v", err)
	}

	c.HTML(http.StatusOK, "plan/index", gin.H{
		"title":   "VIP订阅",
		"user_id": userId,
		"user":    user,
		"ctx":     c,
		"plans":   plans,
	})
}
