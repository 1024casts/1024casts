package course

import (
	"1024casts/backend/model"
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
	Username string `json:"username"`
	Page     int    `json:"page"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64               `json:"totalCount"`
	List       []*model.CourseModel `json:"list"`
}

type SwaggerListResponse struct {
	TotalCount uint64              `json:"totalCount"`
	List       []model.CourseModel `json:"list"`
}
