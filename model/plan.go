package model

import (
	"sync"
	"time"
)

type PlanModel struct {
	ID               uint64    `gorm:"column:id" json:"id"`
	Alias            string    `gorm:"column:alias" json:"alias"`
	Description      string    `gorm:"column:description" json:"description"`
	Name             string    `gorm:"column:name" json:"name"`
	Price            float64   `gorm:"column:price" json:"price"`
	PromoEnd         time.Time `gorm:"column:promo_end" json:"promo_end"`
	PromoPrice       float64   `gorm:"column:promo_price" json:"promo_price"`
	PromoStart       time.Time `gorm:"column:promo_start" json:"promo_start"`
	Status           int       `gorm:"column:status" json:"status"`
	UserID           int       `gorm:"column:user_id" json:"user_id"`
	ValidDays        int       `gorm:"column:valid_days" json:"valid_days"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt        time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at" json:"updated_at"`
	FormatPromoStart string    `json:"format_promo_start"`
	FormatPromoEnd   string    `json:"format_promo_end"`
}

// TableName sets the insert table name for this struct type
func (p *PlanModel) TableName() string {
	return "plans"
}

type PlanList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*PlanModel
}
