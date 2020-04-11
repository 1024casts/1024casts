package plan

import (
	"strconv"

	"github.com/1024casts/1024casts/pkg/app"
	"github.com/1024casts/1024casts/pkg/errno"

	"github.com/1024casts/1024casts/service"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// @Summary Update a course info by the course identifier
// @Description Update a course by ID
// @Tags course
// @Accept  json
// @Produce  json
// @Param id path uint64 true "The course's database id index num"
// @Param user body model.CourseModel true "The course info"
// @Success 200 {object} handler.Response "{"code":0,"message":"OK","data":null}"
// @Router /plans/{id} [put]
func Update(c *gin.Context) {
	log.Info("Update plan function called.")
	// Get the course id from the url parameter.
	planID, _ := strconv.Atoi(c.Param("id"))

	// Binding the course data.
	var plan CreateRequest
	if err := c.Bind(&plan); err != nil {
		app.Response(c, errno.ErrBind, nil)
		return
	}

	srv := service.NewPlanService()
	_, err := srv.GetPlanById(planID)
	if err != nil {
		log.Warn("[plan] info", lager.Data{"id": planID})
		app.Response(c, errno.ErrCourseNotFound, nil)
		return
	}

	// Save changed fields.
	planMap := make(map[string]interface{}, 0)
	planMap["name"] = plan.Name
	planMap["description"] = plan.Description
	planMap["alias"] = plan.Alias
	planMap["price"] = plan.Price
	planMap["promo_price"] = plan.PromoPrice
	planMap["promo_start"] = plan.PromoStart
	planMap["promo_end"] = plan.PromoEnd
	planMap["status"] = plan.Status

	if err := srv.UpdatePlan(planMap, planID); err != nil {
		app.Response(c, errno.InternalServerError, nil)
		return
	}

	app.Response(c, nil, nil)
}
