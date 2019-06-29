package plan

import (
	"net/http"

	"github.com/1024casts/1024casts/pkg/constvar"

	"github.com/1024casts/1024casts/model"

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

	planSrv := service.NewPlanService()
	plan, err := planSrv.GetPlanByAlias(c.Param("alias"))
	if err != nil {
		log.Warnf("[plan] get plan list err: %v", err)
	}

	// step1: 创建订单
	orderSrv := service.NewOrderService()
	orderModel := model.OrderModel{}
	orderModel.Id = util.GenOrderNo()
	orderModel.PayMethod = constvar.PayMethodWeiXin
	orderModel.OrderAmount = plan.Price
	orderModel.PayAmount = plan.Price
	orderModel.UserId = userId
	orderModel.Status = constvar.OrderStatusPending
	orderId, err := orderSrv.CreateOrder(orderModel)
	if err != nil {
		log.Warnf("[plans] create order err: %v", err)
	}

	var qrCodeUrl string
	if orderId > 0 {
		// step2: 创建支付订单
		payJsConfig := &payjs.Config{
			Key:       viper.GetString("pay_js.secret_key"),
			MchID:     viper.GetString("pay_js.mch_id"),
			NotifyUrl: "",
		}
		pay := payjs.New(payJsConfig)
		payNative := pay.GetNative()
		payAmount := plan.Price * 100
		resp, err := payNative.Create(int64(payAmount), plan.Name, string(orderId), "", "")
		if err != nil {
			log.Warnf("[plans] create pay order err: %v", err)
		}
		log.Infof("[plans] pay create resp: %+v", resp)
		qrCodeUrl = resp.Qrcode

		// step3: write to order items

		// step4: 更新相关数据到order
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
