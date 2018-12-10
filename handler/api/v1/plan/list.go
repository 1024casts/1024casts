package plan

import (
	. "github.com/1024casts/1024casts/handler"
	"github.com/1024casts/1024casts/pkg/errno"
	"github.com/1024casts/1024casts/service"

	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// @Summary 获取plan列表
// @Description List plans
// @Tags order
// @Accept  json
// @Produce  json
// @Param order body order.ListRequest true "List orders"
// @Success 200 {object} order.SwaggerListResponse "{"code":0,"message":"OK","data":{"totalCount":1,"list":[{"id":0,"username":"admin","random":"user 'admin' get random string 'EnqntiSig'","password":"$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG","createdAt":"2018-05-28 00:25:33","updatedAt":"2018-05-28 00:25:33"}]}}"
// @Router /plans [get]
func List(c *gin.Context) {
	log.Info("List function called.")
	var r ListRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		log.Error("get page error", err)
	}
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		log.Error("get limit error", err)
	}

	offset := (page - 1) * limit

	srv := service.NewPlanService()
	orderMap := make(map[string]interface{})
	infos, count, err := srv.GetPlanList(orderMap, offset, limit)
	if err != nil {
		SendResponse(c, err, nil)
		return
	}

	SendResponse(c, nil, ListResponse{
		TotalCount: count,
		List:       infos,
	})
}
