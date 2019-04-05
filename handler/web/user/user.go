package user

import (
	"github.com/1024casts/1024casts/model"
)

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginCredentials struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type RegisterRequest struct {
	Username        string `json:"username" form:"username"`
	Email           string `json:"email" form:"email"`
	Password        string `json:"password" form:"password"`
	ConfirmPassword string `json:"confirm_password" form:"confirm_password"`
}

type CreateResponse struct {
	Id uint64 `json:"id"`
}

type UpdateReq struct {
	Status int `json:"status"`
}

type ListRequest struct {
	Username string `json:"username"`
	Offset   int    `json:"offset"`
	Limit    int    `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64             `json:"totalCount"`
	UserList   []*model.UserModel `json:"userList"`
}

type SwaggerListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []model.UserModel `json:"userList"`
}
