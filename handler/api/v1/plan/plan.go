package plan

import (
	"time"

	"github.com/1024casts/1024casts/model"
)

type CreateRequest struct {
	Name        string    `json:"name"`
	Alias       string    `json:"alias"`
	Description string    `json:"description"`
	Price       float32   `json:"price"`
	PromoPrice  float32   `json:"promo_price"`
	PromoStart  time.Time `json:"promo_start"`
	PromoEnd    time.Time `json:"promo_end"`
	Status      int       `json:"status"`
}

type CreateResponse struct {
	Id uint64 `json:"id"`
}

type ListRequest struct {
	OrderId int `json:"order_id"`
	Page    int `json:"page"`
	Limit   int `json:"limit"`
}

type ListResponse struct {
	TotalCount uint64             `json:"totalCount"`
	List       []*model.PlanModel `json:"list"`
}

type SwaggerListResponse struct {
	TotalCount uint64            `json:"totalCount"`
	List       []model.PlanModel `json:"list"`
}
