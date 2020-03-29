package user

import (
	"github.com/1024casts/1024casts/model"
	"github.com/gin-gonic/gin"
)

func Endpoint(parentRoute *gin.RouterGroup) {
	router := parentRoute.Group("/users")
	router.POST("", Create)
	router.DELETE("/:id", Delete)
	router.PUT("/:id", Update)
	router.GET("", List)
	router.GET("/token", Get)
	//u.GET("/:id", user.Get)
	router.PUT("/:id/status", UpdateStatus)
}

type CreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LoginCredentials struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
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
	List       []*model.UserModel `json:"list"`
}

type SwaggerListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	UserList   []model.UserModel `json:"userList"`
}
