package course

import (
	"github.com/1024casts/1024casts/model"
)

type CreateRequest struct {
	Name        string `json:"name"`
	Type        string `json:"type"`
	Description string `json:"description"`
	Slug        string `json:"slug"`
	CoverImage  string `json:"cover_image"`
	UserId      int    `json:"user_id"`
	IsPublish   int    `json:"is_publish"`
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
	TotalCount uint64               `json:"totalCount"`
	List       []*model.CourseModel `json:"list"`
}

type SectionListResponse struct {
	TotalCount uint64                `json:"totalCount"`
	List       []*model.SectionModel `json:"list"`
}

type SwaggerListResponse struct {
	TotalCount uint64              `json:"totalCount"`
	List       []model.CourseModel `json:"list"`
}
