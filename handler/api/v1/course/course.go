package course

import (
	"github.com/1024casts/1024casts/model"
	"github.com/gin-gonic/gin"
)

func Endpoint(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/courses")

	router.POST("", Create)
	router.PUT("/:id", Update)
	router.PUT("/:id/updatePublish", UpdatePublish)
	router.GET("", List)
	router.GET("/:id", Get)
	router.GET("/:id/videos", Video)
}

type CreateRequest struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Keywords     string `json:"keywords"`
	Description  string `json:"description"`
	Slug         string `json:"slug"`
	CoverKey     string `json:"cover_key"`
	Content      string `json:"content"`
	UserId       int    `json:"user_id"`
	UpdateStatus int    `json:"update_status"`
	IsPublish    int    `json:"is_publish"`
}

type CreateResponse struct {
	Id uint64 `json:"id"`
}

type ListRequest struct {
	Name         string `json:"name"`
	UpdateStatus int    `json:"update_status"`
	Page         int    `json:"page"`
	Limit        int    `json:"limit"`
}

type SectionListRequest struct {
	CourseId uint64 `json:"course_id"`
	Name     string `json:"name"`
	Order    int    `json:"order"`
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount int                 `json:"totalCount"`
	List       []*model.CourseInfo `json:"list"`
}

type SectionListResponse struct {
	TotalCount uint64                `json:"totalCount"`
	List       []*model.SectionModel `json:"list"`
}

type VideoListResponse struct {
	TotalCount uint64              `json:"totalCount"`
	List       []*model.VideoModel `json:"list"`
}

type SwaggerListResponse struct {
	TotalCount uint64              `json:"totalCount"`
	List       []model.CourseModel `json:"list"`
}
