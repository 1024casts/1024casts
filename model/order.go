package model

import (
	"sync"
	"time"
)

// order
type OrderModel struct {
	Id          uint64           `gorm:"primary_key;column:id" json:"id"`
	OrderId     string           `json:"order_id" gorm:"-"`
	OrderAmount float64          `json:"order_amount" gorm:"column:order_amount;" binding:"required"`
	PayAmount   float64          `json:"pay_amount" gorm:"column:pay_amount;" binding:"required"`
	PayMethod   string           `json:"pay_method" gorm:"column:pay_method;" binding:"required"`
	PaidAt      time.Time        `gorm:"column:paid_at" json:"paid_at"`
	CompletedAt time.Time        `gorm:"column:completed_at" json:"completed_at"`
	CanceledAt  time.Time        `gorm:"column:canceled_at" json:"canceled_at"`
	QrcodeId    int              `json:"qrcode_id" gorm:"column:qrcode_id" binding:"omitempty"`
	TradeId     string           `json:"trade_id" gorm:"column:trade_id" binding:"omitempty"`
	Status      string           `json:"status" gorm:"column:status" binding:"omitempty"`
	UserId      uint64           `json:"user_id" gorm:"column:user_id" binding:"omitempty"`
	CreatedAt   time.Time        `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time        `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   *time.Time       `gorm:"column:deleted_at" sql:"index" json:"-"`
	Items       []OrderItemModel `gorm:"FOREIGNKEY:OrderID;ASSOCIATION_FOREIGNKEY:ID"`
}

func (c *OrderModel) TableName() string {
	return "orders"
}

type OrderList struct {
	Lock  *sync.Mutex
	IdMap map[uint64]*OrderInfo
}

type OrderInfo struct {
	Id          uint64           `json:"id"`
	OrderId     string           `json:"order_id"`
	OrderAmount float64          `json:"order_amount"`
	PayAmount   float64          `json:"pay_amount"`
	PayMethod   string           `json:"pay_method"`
	PaidAt      string           `json:"paid_at"`
	CompletedAt time.Time        `json:"completed_at"`
	CanceledAt  time.Time        `json:"canceled_at"`
	QrcodeId    int              `json:"qrcode_id"`
	TradeId     string           `json:"trade_id"`
	Status      string           `json:"status"`
	UserId      uint64           `json:"user_id"`
	CreatedAt   string           `json:"created_at"`
	UpdatedAt   string           `json:"updated_at"`
	OrderItems  []OrderItemModel `json:"items"`
}
