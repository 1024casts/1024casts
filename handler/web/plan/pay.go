package plan

import (
	"net/http"

	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Pay(c *gin.Context) {
	userId := util.GetUserId(c)
	srv := service.NewUserService()
	user, _ := srv.GetUserById(userId)

	planSrv := service.NewPlanService()
	plan, err := planSrv.GetPlanByAlias(c.Param("alias"))
	if err != nil {
		log.Warnf("[plan] get plan list err: %v", err)
	}

	orderSrv := service.NewOrderService()
	orderId, qrCodeUrl, err := orderSrv.CreateOrder(userId, plan)
	if err != nil {
		log.Warnf("[plans] create order err: %v", err)
	}

	c.HTML(http.StatusOK, "plan/pay", gin.H{
		"title":   "VIP订阅",
		"user_id": userId,
		"user":    user,
		"ctx":     c,
		"plan":    plan,
		"orderId": orderId,
		"qrCode":  qrCodeUrl,
	})
}
