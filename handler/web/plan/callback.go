package plan

import (
	"fmt"
	"strconv"

	"github.com/1024casts/1024casts/pkg/constvar"
	"github.com/1024casts/1024casts/service"
	"github.com/lexkong/log"
	"github.com/qingwg/payjs"
	"github.com/spf13/viper"

	"github.com/qingwg/payjs/notify"

	"github.com/gin-gonic/gin"
)

func Callback(c *gin.Context) {
	log.Infof("entry pay callback...")
	payJsConfig := &payjs.Config{
		Key:       viper.GetString("pay_js.secret_key"),
		MchID:     viper.GetString("pay_js.mch_id"),
		NotifyUrl: "",
	}
	Pay := payjs.New(payJsConfig)

	// 传入request和responseWriter
	PayNotify := Pay.GetNotify(c.Request, c.Writer)

	//设置接收消息的处理方法
	PayNotify.SetMessageHandler(func(msg notify.Message) {
		//这里处理支付成功回调，一般是修改数据库订单信息等等
		//msg即为支付成功异步通知过来的内容

		log.Infof("[callback] msg info: %#v", msg)

		// 支付成功
		// step1: get order id
		orderId, err := strconv.Atoi(msg.OutTradeNo)
		if err != nil {
			log.Warnf("[plan] get order info err, %v", err)
			fmt.Printf("error")
			return
		}

		// step2: check order is exist
		orderSrv := service.NewOrderService()
		order, err := orderSrv.GetOrderById(orderId)
		if err != nil {
			log.Warnf("[plan] get order info err, %v", err)
			fmt.Printf("error")
			return
		}

		// step3: check order status is paid
		if order.Status == constvar.OrderStatusPaid {
			fmt.Printf("success")
			return
		}

		// step4: get order item
		orderItem, err := orderSrv.GetOrderItemById(orderId)
		if err != nil {
			log.Warnf("[plan] get order item err, %v", err)
			fmt.Printf("error")
			return
		}

		// step5: update order status to paid
		err = orderSrv.ConfirmOrderPaid(order, orderItem, msg.TimeEnd)
		if err != nil {
			log.Warnf("[plan] update order status to paid err, %v", err)
			fmt.Printf("error")
			return
		}
	})

	//处理消息接收以及回复
	err := PayNotify.Serve()
	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
