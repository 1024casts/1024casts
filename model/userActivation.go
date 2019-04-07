package model

import "time"

type UserActivationModel struct {
	ID        int       `gorm:"column:id" json:"id"`
	Token     string    `gorm:"column:token" json:"token"`
	UserID    uint64    `gorm:"column:user_id" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

// TableName sets the insert table name for this struct type
func (u *UserActivationModel) TableName() string {
	return "user_activations"
}
