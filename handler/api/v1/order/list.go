package order

import (
	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/service"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// @Summary 获取订单列表
// @Description List orders
// @Tags order
// @Accept  json
// @Produce  json
// @Param order body order.ListRequest true "List orders"
// @Success 200 {object} order.SwaggerListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"list":[{"id":0,"username":"admin","random":"user 'admin' get random string 'EnqntiSig'","password":"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG","createdAt":"2018-05-28 00:25:33","updatedAt":"2018-05-28 00:25:33"}]}}"
// @Router /orders [get]
func List(c *gin.Context) {
	log.Info("List function called.")
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		app.Response(c, errno.ErrBind, nil)
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		log.Error("get page error", err)
	}
	limit, err := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	if err != nil {
		log.Error("get limit error", err)
	}

	offset := (page - 1) * limit

	srv := service.NewOrderService()

	idType := c.Query("idType")
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	orderMap := make(map[string]interface{})
	if id != 0 {
		switch idType {
		case "orderId":
			orderMap["id"] = id
		case "tradeId":
			orderMap["trade_id"] = id
		}
	}

	// 支付状态
	payStatus := c.Query("status")
	if payStatus != "all" {
		orderMap["status"] = payStatus
	}

	infos, count, err := srv.GetOrderList(orderMap, offset, limit)
	if err != nil {
		app.Response(c, err, nil)
		return
	}

	app.Response(c, nil, ListResponse{
		TotalCount: count,
		List:       infos,
	})
}
