package video

import (
	"1024casts/backend/model"
)

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
