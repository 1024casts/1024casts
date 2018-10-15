package order

import (
	. "1024casts/backend/handler"
	"1024casts/backend/pkg/errno"
	"1024casts/backend/service"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// @Summary List the orders in the database
// @Description List orders
// @Tags order
// @Accept  json
// @Produce  json
// @Param course body orders.ListRequest true "List orders"
// @Success 200 {object} order.SwaggerListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"userList":[{"id":0,"username":"admin","random":"user 'admin' get random string 'EnqntiSig'","password":"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG","createdAt":"2018-05-28 00:25:33","updatedAt":"2018-05-28 00:25:33"}]}}"
// @Router /orders [get]
func List(c *gin.Context) {
	log.Info("List function called.")
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		log.Error("get page error", err)
	}
	limit, err := strconv.Atoi(c.Query("limit"))
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
	if payStatus != "" {
		orderMap["status"] = payStatus
	}

	infos, count, err := srv.GetOrderList(orderMap, offset, limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		List:       infos,
	})
}
