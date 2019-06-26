package plan

import (
	"net/http"

	"github.com/qingwg/payjs"
	"github.com/spf13/viper"

	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func Pay(c *gin.Context) {
	userId := util.GetUserId(c)
	srv := service.NewUserService()
	user, _ := srv.GetUserById(userId)

	planMap := make(map[string]interface{})
	planMap["status"] = 1
	planSrv := service.NewPlanService()
	plans, _, err := planSrv.GetPlanList(planMap, 0, 10)
	if err != nil {
		log.Warnf("[plan] get plan list err: %v", err)
	}

	// 创建支付订单
	payJsConfig := &payjs.Config{
		Key:       viper.GetString("pay_js.secret_key"),
		MchID:     viper.GetString("pay_js.mch_id"),
		NotifyUrl: "",
	}
	pay := payjs.New(payJsConfig)
	payNative := pay.GetNative()
	resp, err := payNative.Create(10, "测试订单001", "test1000000", "", "")
	if err != nil {
		log.Warnf("[plans] create pay order err: %v", err)
	}
	log.Infof("[plans] pay create resp: %+v", resp)

	c.HTML(http.StatusOK, "plan/index", gin.H{
		"title":   "VIP订阅",
		"user_id": userId,
		"user":    user,
		"ctx":     c,
		"plans":   plans,
		"qr_code": resp.Qrcode,
	})
}
