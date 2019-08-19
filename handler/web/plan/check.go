package plan

import (
	"strconv"

	"github.com/lexkong/log"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/pkg/constvar"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/service"
	"github.com/gin-gonic/gin"
)

func Check(c *gin.Context) {

	orderId, ok := c.GetQuery("order_id")
	if !ok {
		log.Warnf("[check] get order_id param err")
		app.Response(c, errno.ErrParam, nil)
		return
	}

	oId, err := strconv.Atoi(orderId)
	if err != nil {
		log.Warnf("[check] convert order_id param err")
		app.Response(c, errno.ErrParam, nil)
		return
	}

	orderSrv := service.NewOrderService()
	order, err := orderSrv.GetOrderById(oId)
	if err != nil {
		app.Response(c, errno.ErrDataIsNotExist, nil)
		return
	}

	// pay success
	if order.Status == constvar.OrderStatusPaid {
		app.Response(c, errno.OK, nil)
		return
	}

	app.Response(c, errno.InternalServerError, nil)
	return
}
