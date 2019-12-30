package model

import "time"

type UserMemberModel struct {
	Id        uint64    `gorm:"primary_key;AUTO_INCREMENT;column:id" json:"id"`
	EndTime   time.Time `gorm:"column:end_time" json:"end_time"`
	StartTime time.Time `gorm:"column:start_time" json:"start_time"`
	Status    int       `gorm:"column:status" json:"status"`
	Type      uint64    `gorm:"column:type" json:"type"`
	UserID    uint64    `gorm:"column:user_id" json:"user_id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// TableName sets the insert table name for this struct type
func (u *UserMemberModel) TableName() string {
	return "user_members"
}
