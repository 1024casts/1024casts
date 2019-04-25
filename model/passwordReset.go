package model

import "time"

type PasswordResetModel struct {
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	Email     string    `gorm:"column:email" json:"email"`
	Token     string    `gorm:"column:token" json:"token"`
}

// TableName sets the insert table name for this struct type
func (p *PasswordResetModel) TableName() string {
	return "password_resets"
}
