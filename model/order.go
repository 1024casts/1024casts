package model

import (
	"sync"
	"time"
)

// order
type OrderModel struct {
	Id          uint64     `gorm:"primary_key;column:id" json:"id"`
	OrderAmount float64    `json:"order_amount" gorm:"column:order_amount;" binding:"required"`
	PayAmount   float64    `json:"pay_amount" gorm:"column:pay_amount;" binding:"required"`
	PayMethod   string     `json:"pay_method" gorm:"column:pay_method;" binding:"required"`
	PaidAt      time.Time  `gorm:"column:paid_at" json:"paid_at"`
	CompletedAt time.Time  `gorm:"column:completed_at" json:"completed_at"`
	CanceledAt  time.Time  `gorm:"column:canceled_at" json:"canceled_at"`
	QrcodeId    int        `json:"qrcode_id" gorm:"column:qrcode_id" binding:"omitempty"`
	TradeId     string     `json:"trade_id" gorm:"column:trade_id" binding:"omitempty"`
	Status      string     `json:"status" gorm:"column:status" binding:"omitempty"`
	UserId      int        `json:"user_id" gorm:"column:user_id" binding:"omitempty"`
	CreatedAt   time.Time  `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   *time.Time `gorm:"column:deleted_at" sql:"index" json:"-"`
}

func (c *OrderModel) TableName() string {
	return "orders"
}

type OrderList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*OrderModel
}
