package video

import (
	"github.com/1024casts/1024casts/model"
	"github.com/gin-gonic/gin"
)

func Endpoint(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/videos")
	router.GET("/:course_id", List)
}

type CreateRequest struct {
	model.VideoModel
}

type CreateResponse struct {
	Id uint64 `json:"id"`
}

type ListRequest struct {
	Name     string `json:"name"`
	CourseId int    `json:"course_id"`
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64              `json:"totalCount"`
	List       []*model.VideoModel `json:"list"`
}

type SwaggerListResponse struct {
	TotalCount uint64             `json:"totalCount"`
	List       []model.VideoModel `json:"list"`
}
