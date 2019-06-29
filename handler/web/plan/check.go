package plan

import (
	"strconv"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/pkg/constvar"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/service"
	"github.com/gin-gonic/gin"
)

func Check(c *gin.Context) {

	orderId := c.Param("order_id")

	oId, err := strconv.Atoi(orderId)
	if err != nil {
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
