package plan

import (
	"strconv"

	. "1024casts/backend/handler"
	"1024casts/backend/pkg/errno"

	"1024casts/backend/service"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
)

// @Summary Get an plan by the plan identifier
// @Description Get an plan by alias or id
// @Tags plan
// @Accept  json
// @Produce  json
// @Param username path string true "Username"
// @Success 200 {object} model.PlanModel "{"code":0,"message":"OK","data":{"username":"kong","password":"$2a$10$E0kwtmtLZbwW/bDQ8qI8e.eHPqhQOW9tvjwpyo/p05f/f4Qvr3OmS"}}"
// @Router /plans/{alias} [get]
func Get(c *gin.Context) {
	log.Info("Get function called.")

	// Get the user by the `id` from the database.
	// Get the user id from the url parameter.
	srv := service.NewPlanService()
	planId, _ := strconv.Atoi(c.Param("id"))

	plan, err := srv.GetPlanById(planId)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	if planId == 0 {
		alias := c.Param("alias")
		plan, err = srv.GetPlanByAlias(alias)
		if err != nil {
			SendResponse(c, errno.ErrUserNotFound, nil)
			return
		}
	}

	SendResponse(c, nil, plan)
}
