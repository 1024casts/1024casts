package user

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/pkg/pagination"
	"github.com/1024casts/1024casts/service"
	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
)

func OrderList(c *gin.Context) {
	userId := util.GetUserId(c)
	srv := service.NewUserService()
	user, _ := srv.GetUserById(userId)

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		log.Error("get page error", err)
	}
	limit := 10
	offset := (page - 1) * limit

	orderSrv := service.NewOrderService()
	orders, count, err := orderSrv.GetOrderListByUserId(userId, offset, limit)
	log.Infof("[order] get list: %v", orders)
	if err != nil {
		log.Warnf("[order] get order list err: %v", err)
	}
	pagination := pagination.NewPagination(c.Request, count, limit)

	c.HTML(http.StatusOK, "user/orderList", gin.H{
		"title":   "我的订单列表",
		"user_id": userId,
		"user":    user,
		"ctx":     c,
		"orders":  orders,
		"pages":   template.HTML(pagination.Pages()),
	})
}
