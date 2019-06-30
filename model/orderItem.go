package model

import "time"

type OrderItemModel struct {
	Amount    float64   `gorm:"column:amount" json:"amount"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	DeletedAt time.Time `gorm:"column:deleted_at" json:"deleted_at"`
	ID        int       `gorm:"column:id" json:"id"`
	ItemID    int       `gorm:"column:item_id" json:"item_id"`
	Name      string    `gorm:"column:name" json:"name"`
	OrderID   int64     `gorm:"column:order_id" json:"order_id"`
	Price     float64   `gorm:"column:price" json:"price"`
	Quantity  int       `gorm:"column:quantity" json:"quantity"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID    int       `gorm:"column:user_id" json:"user_id"`
}

// TableName sets the insert table name for this struct type
func (o *OrderItemModel) TableName() string {
	return "order_items"
}
