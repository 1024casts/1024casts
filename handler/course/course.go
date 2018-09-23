package course

import (
	"1024casts/backend/model"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type CreateResponse struct {
	Username string `json:"username"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	List   []*model.CourseModel `json:"list"`
}

type SwaggerListResponse struct {
	TotalCount uint64           `json:"totalCount"`
	List   []model.CourseModel `json:"list"`
}
