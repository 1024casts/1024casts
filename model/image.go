package model

import "time"

type ImageModel struct {
	ID        uint64    `gorm:"column:id" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	ImageName string    `gorm:"column:image_name" json:"image_name"`
	ImagePath string    `gorm:"column:image_path" json:"image_path"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
	UserID    uint64    `gorm:"column:user_id" json:"user_id"`
}

// TableName sets the insert table name for this struct type
func (i *ImageModel) TableName() string {
	return "images"
}
